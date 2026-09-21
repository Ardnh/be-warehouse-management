package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
)

type LoginSessionRepository interface {
	Create(ctx context.Context, session *entity.LoginSession) error
	Update(ctx context.Context, session *entity.LoginSession) error
}
