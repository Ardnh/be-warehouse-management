package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionRepositoryImpl struct{ db *gorm.DB }

func NewPermissionRepository(db *gorm.DB) domainrepositories.PermissionRepository {
	return &PermissionRepositoryImpl{db: db}
}

func (r *PermissionRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.Permission, int64, error) {
	var items []entity.Permission
	var total int64
	q := Conn(ctx, r.db).Model(&entity.Permission{})
	if filter.Search != "" {
		q = q.Where("resource ILIKE ? OR action ILIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := applyPermissionListFilter(q, filter, "resource", "action").Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *PermissionRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Permission, error) {
	var permission entity.Permission
	if err := Conn(ctx, r.db).First(&permission, id).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepositoryImpl) Create(ctx context.Context, permission entity.Permission) error {
	return Conn(ctx, r.db).Create(&permission).Error
}

func (r *PermissionRepositoryImpl) Update(ctx context.Context, permission *entity.Permission) error {
	return Conn(ctx, r.db).Save(permission).Error
}

func (r *PermissionRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return Conn(ctx, r.db).Delete(&entity.Permission{}, id).Error
}

func applyPermissionListFilter(query *gorm.DB, filter domainrepositories.Filter, searchFields ...string) *gorm.DB {
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PageSize <= 0 || filter.PageSize >= 1000 {
		filter.PageSize = 30
	}
	if filter.Search != "" && len(searchFields) > 0 {
		pattern := "%" + filter.Search + "%"
		condition := ""
		args := make([]interface{}, 0, len(searchFields))
		for _, field := range searchFields {
			if condition != "" {
				condition += " OR "
			}
			condition += field + " ILIKE ?"
			args = append(args, pattern)
		}
		query = query.Where(condition, args...)
	}
	// if filter.SortBy != "" {
	// 	sortable := map[string]bool{"id": true, "code": true, "name": true, "type": true, "status": true, "created_at": true, "updated_at": true}
	// 	if sortable[filter.SortBy] {
	// 		direction := strings.ToUpper(filter.SortDir)
	// 		if direction != "DESC" {
	// 			direction = "ASC"
	// 		}
	// 		query = query.Order(fmt.Sprintf("%s %s", filter.SortBy, direction))
	// 	}
	// }
	return query.Offset((filter.Page - 1) * filter.PageSize).Limit(filter.PageSize)
}
