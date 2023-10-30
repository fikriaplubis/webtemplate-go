package dao

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	MenuId       uuid.UUID `json:"menu_id"`
	MenuName     string    `json:"menu_name"`
	ParentMenuId uuid.UUID `json:"parent_menu_id"`
	ParentMenu   *Menu     `json:"parent_menu" gorm:"foreignKey:ParentMenuId;references:MenuId"`
	CreatedBy    uuid.UUID `json:"created_by"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
	UpdatedDate  time.Time `json:"updated_date"`
}

func (m *Menu) TableName() string {
	return "menus"
}

func (m *Menu) SetMenuId(menu_id uuid.UUID) {
	m.MenuId = menu_id
}
func (m Menu) GetMenuId() uuid.UUID {
	return m.MenuId
}

func (m *Menu) SetCreatedBy(created_by uuid.UUID) {
	m.CreatedBy = created_by
}
func (m Menu) GetCreatedBy() uuid.UUID {
	return m.CreatedBy
}

func (m *Menu) SetCreatedDate(created_date time.Time) {
	m.CreatedDate = created_date
}
func (m Menu) GetCreatedDate() time.Time {
	return m.CreatedDate
}

func (m *Menu) SetUpdatedBy(updated_by uuid.UUID) {
	m.UpdatedBy = updated_by
}
func (m Menu) GetUpdatedBy() uuid.UUID {
	return m.UpdatedBy
}

func (m *Menu) SetUpdatedDate(updated_date time.Time) {
	m.UpdatedDate = updated_date
}
func (m Menu) GetUpdatedDate() time.Time {
	return m.UpdatedDate
}
