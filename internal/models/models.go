package models

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

type Model interface {
	Id() uuid.UUID
}

type BaseModel struct {
	ID uuid.UUID `json:"uuid"`
}

func (m BaseModel) Id() uuid.UUID {
	return m.ID
}

func ModelFrom[T Model](jsonStr string) (T, error) {
	var model T
	err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(model)
	return model, err
}
