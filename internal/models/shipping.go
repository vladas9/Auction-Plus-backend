package models

import (
	"time"

	"github.com/google/uuid"
)

type ShippingModel struct {
	ShippingId        uuid.UUID `json:"shipping_id"`
	TransactionId     uuid.UUID `json:"transaction_id"`
	ShippingAddress   string    `json:"shipping_address"`
	TrackingNumber    string    `json:"tracking_number"`
	Status            Status    `json:"status"`
	EstimatedDelivary time.Time `json:"estimated_delivary"`
}
