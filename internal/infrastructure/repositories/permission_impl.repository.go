package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PermissionRepositoryImpl struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewPermissionRepository(db *gorm.DB, log *logrus.Logger) domainrepositories.PermissionRepository {
	return &PermissionRepositoryImpl{db: db, log: log}
}

func (r *PermissionRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]*entity.Permission, int64, error) {
	var items []*entity.Permission
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
		args := make([]any, 0, len(searchFields))
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

func (r *PermissionRepositoryImpl) HasPermission(ctx context.Context, userID uuid.UUID, resource string, action string) (bool, error) {
	entry := r.log.WithFields(logrus.Fields{
		"component": "permission_repository",
		"method":    "HasPermission",
		"user_id":   userID.String(),
		"resource":  resource,
		"action":    action,
	})

	var count int64

	start := time.Now()
	err := Conn(ctx, r.db).
		Table("user_assignments ua").
		Joins("JOIN roles r ON r.id = ua.role_id").
		Joins("JOIN role_permissions rp ON rp.role_id = r.id").
		Joins("JOIN permissions p ON p.id = rp.permission_id").
		Where(`
            ua.user_id = ?
            AND p.resource = ?
            AND p.action = ?
        `,
			userID,
			resource,
			action,
		).
		Count(&count).Error

	entry = entry.WithField("duration_ms", time.Since(start).Milliseconds())

	if err != nil {
		entry.WithError(err).Error("failed to query user permission")
		return false, fmt.Errorf("check permission %s:%s for user %s: %w", resource, action, userID, err)
	}

	entry.WithFields(logrus.Fields{
		"match_count": count,
		"allowed":     count > 0,
	}).Debug("permission check completed")

	return count > 0, nil
}
