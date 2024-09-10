package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationModel struct {
	NotificationId uuid.UUID `json:"notification_id"`
	UserId         uuid.UUID `json:"user_id"`
	Message        string    `json:"message"`
	Timestamp      time.Time `json:"tomestamp"`
	IsReaded       bool      `json:"is_readed"`
}
