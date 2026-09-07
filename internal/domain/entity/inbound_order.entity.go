package entity

import (
	"time"

	"github.com/google/uuid"
)

type InboundOrderStatus string

const (
	InboundStatusDraft             InboundOrderStatus = "DRAFT"
	InboundStatusOpen              InboundOrderStatus = "OPEN"
	InboundStatusPartiallyReceived InboundOrderStatus = "PARTIALLY_RECEIVED"
	InboundStatusReceived          InboundOrderStatus = "RECEIVED"
	InboundStatusCancelled         InboundOrderStatus = "CANCELLED"
)

func (s InboundOrderStatus) String() string { return string(s) }

func (s InboundOrderStatus) IsValid() bool {
	switch s {
	case InboundStatusDraft, InboundStatusOpen, InboundStatusPartiallyReceived,
		InboundStatusReceived, InboundStatusCancelled:
		return true
	}
	return false
}

// IsEditable menandai status yang masih boleh diubah item-nya.
func (s InboundOrderStatus) IsEditable() bool {
	return s == InboundStatusDraft || s == InboundStatusOpen
}

// IsFinal menandai status akhir yang tidak bisa ditransisi lagi.
func (s InboundOrderStatus) IsFinal() bool {
	return s == InboundStatusReceived || s == InboundStatusCancelled
}

type InboundOrder struct {
	ID                uuid.UUID          `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	WarehouseID       uuid.UUID          `gorm:"type:uuid;not null;index"`
	CustomerID        uuid.UUID          `gorm:"type:uuid;not null;index"`
	OrderNumber       string             `gorm:"type:varchar(50);not null;uniqueIndex"`
	Status            InboundOrderStatus `gorm:"type:varchar(30);not null;index;default:'DRAFT'"`
	ExpectedArrivalAt *time.Time         `gorm:"type:timestamptz;index"`
	Notes             *string            `gorm:"type:text"`
	CreatedAt         time.Time          `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt         time.Time          `gorm:"type:timestamptz;not null;default:now()"`

	Customer   *Customer          `gorm:"foreignKey:CustomerID"`
	Items      []InboundOrderItem `gorm:"foreignKey:InboundOrderID"`
	Receivings []Receiving        `gorm:"foreignKey:InboundOrderID"`
	Warehouse  *Warehouse         `gorm:"foreignKey:WarehouseID"`
}

func (InboundOrder) TableName() string { return "inbound_orders" }
