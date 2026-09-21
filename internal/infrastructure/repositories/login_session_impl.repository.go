package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"gorm.io/gorm"
)

type LoginSessionRepositoryImpl struct{ db *gorm.DB }

func NewLoginSessionRepository(db *gorm.DB) repositories.LoginSessionRepository {
	return &LoginSessionRepositoryImpl{db: db}
}

func (r *LoginSessionRepositoryImpl) Create(ctx context.Context, session *entity.LoginSession) error {
	return Conn(ctx, r.db).Create(session).Error
}

func (r *LoginSessionRepositoryImpl) Update(ctx context.Context, session *entity.LoginSession) error {
	return Conn(ctx, r.db).Save(session).Error
}
