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

type InboundOrderRepositoryImpl struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewInboundOrderRepository(db *gorm.DB, log *logrus.Logger) domainrepositories.InboundOrderRepository {
	return &InboundOrderRepositoryImpl{db: db, log: log}
}

func (r *InboundOrderRepositoryImpl) FindAll(ctx context.Context, filter domainrepositories.Filter) ([]entity.InboundOrder, int64, error) {
	entry := r.log.WithFields(logrus.Fields{
		"component": "inbound_order_repository",
		"method":    "FindAll",
		"page":      filter.Page,
		"page_size": filter.PageSize,
	})
	start := time.Now()
	var list []entity.InboundOrder
	var total int64
	query := Conn(ctx, r.db).Model(&entity.InboundOrder{}).Preload("Customer").Preload("Items")
	if filter.Search != "" {
		query = query.Where("order_number ILIKE ?", "%"+filter.Search+"%")
	}
	if err := query.Count(&total).Error; err != nil {
		entry.WithError(err).Error("failed to count inbound orders")
		return nil, 0, err
	}
	query = applyMasterListFilter(query, filter)
	if err := query.Find(&list).Error; err != nil {
		entry.WithError(err).Error("failed to query inbound orders")
		return nil, 0, err
	}
	entry.WithFields(logrus.Fields{
		"result_count": len(list),
		"total":        total,
		"duration_ms":  time.Since(start).Milliseconds(),
	}).Debug("inbound orders queried")
	return list, total, nil
}

func (r *InboundOrderRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.InboundOrder, error) {
	entry := r.log.WithFields(logrus.Fields{
		"component":        "inbound_order_repository",
		"method":           "FindByID",
		"inbound_order_id": id.String(),
	})
	var item entity.InboundOrder
	if err := Conn(ctx, r.db).Preload("Customer").Preload("Items").First(&item, id).Error; err != nil {
		entry.WithError(err).Error("failed to query inbound order")
		return nil, err
	}
	entry.Debug("inbound order queried")
	return &item, nil
}

func (r *InboundOrderRepositoryImpl) Create(ctx context.Context, item entity.InboundOrder) error {
	entry := r.log.WithFields(logrus.Fields{
		"component":        "inbound_order_repository",
		"method":           "Create",
		"inbound_order_id": item.ID.String(),
		"warehouse_id":     item.WarehouseID.String(),
		"customer_id":      item.CustomerID.String(),
		"item_count":       len(item.Items),
	})
	if err := Conn(ctx, r.db).Session(&gorm.Session{FullSaveAssociations: true}).Create(&item).Error; err != nil {
		entry.WithError(err).Error("failed to insert inbound order")
		return err
	}
	entry.WithField("order_number", item.OrderNumber).Debug("inbound order inserted")
	return nil
}

func (r *InboundOrderRepositoryImpl) Update(ctx context.Context, item *entity.InboundOrder) error {
	entry := r.log.WithFields(logrus.Fields{
		"component":        "inbound_order_repository",
		"method":           "Update",
		"inbound_order_id": item.ID.String(),
	})
	if err := Conn(ctx, r.db).Save(item).Error; err != nil {
		entry.WithError(err).Error("failed to update inbound order")
		return err
	}
	entry.Debug("inbound order updated")
	return nil
}

func (r *InboundOrderRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	entry := r.log.WithFields(logrus.Fields{
		"component":        "inbound_order_repository",
		"method":           "Delete",
		"inbound_order_id": id.String(),
	})
	if err := Conn(ctx, r.db).Delete(&entity.InboundOrder{}, id).Error; err != nil {
		entry.WithError(err).Error("failed to delete inbound order")
		return err
	}
	entry.Debug("inbound order deleted")
	return nil
}
