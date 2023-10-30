package dto

import (
	"time"

	"github.com/google/uuid"
)

type ActionResponse struct {
	ActionId    int       `json:"action_id"`
	ActionName  string    `json:"action_name"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	UpdatedDate time.Time `json:"updated_date"`
}
