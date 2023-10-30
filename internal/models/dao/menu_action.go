package dao

import (
	"time"

	"github.com/google/uuid"
)

type MenuAction struct {
	ID          uuid.UUID `json:"id"`
	MenuId      uuid.UUID `json:"menu_id"`
	ActionId    int       `json:"action_id"`
	Menu        *Menu     `json:"menu" gorm:"foreignKey:MenuId;references:MenuId"`
	Action      *Action   `json:"action" gorm:"foreignKey:ActionId;references:ActionId"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	UpdatedDate time.Time `json:"updated_date"`
}

func (m *MenuAction) TableName() string {
	return "menu_action"
}

func (m *MenuAction) SetId(id uuid.UUID) {
	m.ID = id
}
func (m MenuAction) GetId() uuid.UUID {
	return m.ID
}
