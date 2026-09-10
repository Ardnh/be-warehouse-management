package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) repositories.PermissionRepository {
	return &PermissionRepositoryImpl{
		db: db,
	}
}

func (r *PermissionRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.Permission, int64, error) {

	var x []entity.Permission
	var t int64
	q := r.db.WithContext(ctx).Model(&entity.Permission{})
	if filter.Search != "" {
		q = q.Where("resource ILIKE ? OR action ILIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}
	if e := q.Count(&t).Error; e != nil {
		return nil, 0, e
	}
	if e := applyPermissionListFilter(q, filter, "resource", "action").Find(&x).Error; e != nil {
		return nil, 0, e
	}
	return x, t, nil

}

func (r *PermissionRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Permission, error) {
	var permission entity.Permission
	if e := r.db.WithContext(ctx).First(&permission, id).Error; e != nil {
		return nil, e
	}
	return &permission, nil
}

func (r *PermissionRepositoryImpl) Create(ctx context.Context, permission entity.Permission) error {
	if e := r.db.WithContext(ctx).Create(&permission).Error; e != nil {
		return e
	}
	return nil
}

func (r *PermissionRepositoryImpl) Update(ctx context.Context, permission *entity.Permission) error {
	if e := r.db.WithContext(ctx).Save(permission).Error; e != nil {
		return e
	}
	return nil
}

func (r *PermissionRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	if e := r.db.WithContext(ctx).Delete(&entity.Permission{}, id).Error; e != nil {
		return e
	}
	return nil
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
