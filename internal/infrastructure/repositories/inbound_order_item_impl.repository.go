package repositories

import (
	"context"
	"time"

	"github.com/Ardnh/be-warehouse-management/internal/domain/entity"
	domainrepositories "github.com/Ardnh/be-warehouse-management/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InboundOrderItemRepositoryImpl struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewInboundOrderItemRepository(db *gorm.DB, log *logrus.Logger) domainrepositories.InboundOrderItemRepository {
	return &InboundOrderItemRepositoryImpl{db: db, log: log}
}

func (r *InboundOrderItemRepositoryImpl) FindAllByInboundOrder(ctx context.Context, orderID uuid.UUID) ([]entity.InboundOrderItem, error) {
	entry := r.log.WithFields(logrus.Fields{
		"component":        "inbound_order_item_repository",
		"method":           "FindAllByInboundOrder",
		"inbound_order_id": orderID.String(),
	})
	start := time.Now()
	var list []entity.InboundOrderItem
	if err := Conn(ctx, r.db).
		Where("inbound_order_id = ?", orderID).
		Preload("Product").
		Find(&list).Error; err != nil {
		entry.WithError(err).Error("failed to query inbound order items")
		return nil, err
	}
	entry.WithFields(logrus.Fields{
		"result_count": len(list),
		"duration_ms":  time.Since(start).Milliseconds(),
	}).Debug("inbound order items queried")
	return list, nil
}

func (r *InboundOrderItemRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.InboundOrderItem, error) {
	entry := r.log.WithFields(logrus.Fields{
		"component":             "inbound_order_item_repository",
		"method":                "FindByID",
		"inbound_order_item_id": id.String(),
	})
	var item entity.InboundOrderItem
	if err := Conn(ctx, r.db).Preload("Product").First(&item, id).Error; err != nil {
		entry.WithError(err).Error("failed to query inbound order item")
		return nil, err
	}
	entry.Debug("inbound order item queried")
	return &item, nil
}

func (r *InboundOrderItemRepositoryImpl) Create(ctx context.Context, item entity.InboundOrderItem) error {
	entry := r.log.WithFields(logrus.Fields{
		"component":             "inbound_order_item_repository",
		"method":                "Create",
		"inbound_order_item_id": item.ID.String(),
		"inbound_order_id":      item.InboundOrderID.String(),
		"product_id":            item.ProductID.String(),
	})
	if err := Conn(ctx, r.db).Create(&item).Error; err != nil {
		entry.WithError(err).Error("failed to insert inbound order item")
		return err
	}
	entry.Debug("inbound order item inserted")
	return nil
}

func (r *InboundOrderItemRepositoryImpl) Update(ctx context.Context, item *entity.InboundOrderItem) error {
	entry := r.log.WithFields(logrus.Fields{
		"component":             "inbound_order_item_repository",
		"method":                "Update",
		"inbound_order_item_id": item.ID.String(),
	})
	if err := Conn(ctx, r.db).Save(item).Error; err != nil {
		entry.WithError(err).Error("failed to update inbound order item")
		return err
	}
	entry.Debug("inbound order item updated")
	return nil
}

func (r *InboundOrderItemRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	entry := r.log.WithFields(logrus.Fields{
		"component":             "inbound_order_item_repository",
		"method":                "Delete",
		"inbound_order_item_id": id.String(),
	})
	if err := Conn(ctx, r.db).Delete(&entity.InboundOrderItem{}, id).Error; err != nil {
		entry.WithError(err).Error("failed to delete inbound order item")
		return err
	}
	entry.Debug("inbound order item deleted")
	return nil
}
