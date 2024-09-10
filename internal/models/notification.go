package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationModel struct {
	BaseModel
	UserId    uuid.UUID `json:"user_id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"tomestamp"`
	IsRead    bool      `json:"is_read"`
}
