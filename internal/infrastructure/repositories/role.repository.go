package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	dr "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct{ db *gorm.DB }

func NewRoleRepository(db *gorm.DB) dr.RoleRepository { return &RoleRepositoryImpl{db: db} }
func (r *RoleRepositoryImpl) FindAll(c context.Context, f dr.Filter) ([]entity.Role, int64, error) {
	var x []entity.Role
	var t int64
	q := r.db.WithContext(c).Model(&entity.Role{})
	if f.Search != "" {
		q = q.Where("code ILIKE ? OR name ILIKE ?", "%"+f.Search+"%", "%"+f.Search+"%")
	}
	if e := q.Count(&t).Error; e != nil {
		return nil, 0, e
	}
	if e := applyMasterListFilter(q, f).Find(&x).Error; e != nil {
		return nil, 0, e
	}
	return x, t, nil
}
func (r *RoleRepositoryImpl) FindByID(c context.Context, id uuid.UUID) (*entity.Role, error) {
	var x entity.Role
	e := r.db.WithContext(c).First(&x, id).Error
	if e != nil {
		return nil, e
	}
	return &x, nil
}
func (r *RoleRepositoryImpl) Create(c context.Context, x entity.Role) error {
	return r.db.WithContext(c).Create(&x).Error
}
func (r *RoleRepositoryImpl) Update(c context.Context, x *entity.Role) error {
	return r.db.WithContext(c).Save(x).Error
}
func (r *RoleRepositoryImpl) Delete(c context.Context, id uuid.UUID) error {
	return r.db.WithContext(c).Delete(&entity.Role{}, id).Error
}
