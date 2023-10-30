package dao

import (
	"time"

	"github.com/google/uuid"
)

type RoleMenuAction struct {
	ID           uuid.UUID   `json:"id"`
	RoleId       uuid.UUID   `json:"role_id"`
	MenuActionId uuid.UUID   `json:"menu_action_id"`
	Role         *Role       `json:"role" gorm:"foreignKey:RoleId;references:RoleId"`
	MenuAction   *MenuAction `json:"menu_action" gorm:"foreignKey:MenuActionId;references:ID"`
	CreatedBy    uuid.UUID   `json:"created_by"`
	CreatedDate  time.Time   `json:"created_date"`
	UpdatedBy    uuid.UUID   `json:"updated_by"`
	UpdatedDate  time.Time   `json:"updated_date"`
}

func (m *RoleMenuAction) TableName() string {
	return "role_menu_action"
}

func (m *RoleMenuAction) SetId(id uuid.UUID) {
	m.ID = id
}
func (m RoleMenuAction) GetId() uuid.UUID {
	return m.ID
}
