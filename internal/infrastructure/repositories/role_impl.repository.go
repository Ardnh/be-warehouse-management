package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct{ db *gorm.DB }

func NewRoleRepository(db *gorm.DB) domainrepositories.RoleRepository {
	return &RoleRepositoryImpl{db: db}
}

func (r *RoleRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.Role, int64, error) {
	var items []entity.Role
	var total int64
	q := Conn(ctx, r.db).Model(&entity.Role{})
	if filter.Search != "" {
		q = q.Where("code ILIKE ? OR name ILIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := applyMasterListFilter(q, filter).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *RoleRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	var item entity.Role
	if err := Conn(ctx, r.db).
		Preload("RolePermissions.Permission").
		First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *RoleRepositoryImpl) Create(ctx context.Context, item entity.Role) error {
	return Conn(ctx, r.db).Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("RolePermissions", "Assignments").Create(&item).Error; err != nil {
			return err
		}
		return createRolePermissions(tx, item.ID, item.RolePermissions)
	})
}

func (r *RoleRepositoryImpl) Update(ctx context.Context, item *entity.Role) error {
	return Conn(ctx, r.db).Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("RolePermissions", "Assignments").Save(item).Error; err != nil {
			return err
		}
		if err := tx.Where("role_id = ?", item.ID).Delete(&entity.RolePermission{}).Error; err != nil {
			return err
		}
		return createRolePermissions(tx, item.ID, item.RolePermissions)
	})
}

func createRolePermissions(tx *gorm.DB, roleID uuid.UUID, rolePermissions []entity.RolePermission) error {
	if len(rolePermissions) == 0 {
		return nil
	}
	for i := range rolePermissions {
		rolePermissions[i].RoleID = roleID
	}
	return tx.Create(&rolePermissions).Error
}

func (r *RoleRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.Role{}, id).Error
}
