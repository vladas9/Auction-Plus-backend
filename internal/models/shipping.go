package models

import (
	"time"

	"github.com/google/uuid"
)

type ShippingModel struct {
	BaseModel
	TransactionId     uuid.UUID `json:"transaction_id"`
	ShippingAddress   string    `json:"shipping_address"`
	TrackingNumber    string    `json:"tracking_number"`
	Status            Status    `json:"status"`
	EstimatedDelivery time.Time `json:"estimated_delivery"`
}
