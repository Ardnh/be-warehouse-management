package repositories

import (
	"context"

	"github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"gorm.io/gorm"
)

type txKey struct{}

type gormTxManager struct{ db *gorm.DB }

func NewTxManager(db *gorm.DB) repositories.TxManager {
	return &gormTxManager{db: db}
}

func (m *gormTxManager) Do(ctx context.Context, fn func(context.Context) error) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(context.WithValue(ctx, txKey{}, tx))
	})
}

// helper yang dipakai semua repository, satu file yang sama
func Conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return db.WithContext(ctx)
}
