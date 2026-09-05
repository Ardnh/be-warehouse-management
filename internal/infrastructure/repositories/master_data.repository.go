package repositories

import (
	"fmt"
	"strings"

	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"gorm.io/gorm"
)

func applyMasterListFilter(query *gorm.DB, filter domainrepositories.Filter, searchFields ...string) *gorm.DB {
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
	if filter.SortBy != "" {
		sortable := map[string]bool{"id": true, "code": true, "name": true, "type": true, "status": true, "created_at": true, "updated_at": true}
		if sortable[filter.SortBy] {
			direction := strings.ToUpper(filter.SortDir)
			if direction != "DESC" {
				direction = "ASC"
			}
			query = query.Order(fmt.Sprintf("%s %s", filter.SortBy, direction))
		}
	}
	return query.Offset((filter.Page - 1) * filter.PageSize).Limit(filter.PageSize)
}
