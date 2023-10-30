package dao

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	RoleId      uuid.UUID `json:"role_id"`
	RoleName    string    `json:"role_name"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	UpdatedDate time.Time `json:"updated_date"`
}

func (m *Role) TableName() string {
	return "roles"
}

func (m *Role) SetRoleId(role_id uuid.UUID) {
	m.RoleId = role_id
}
func (m Role) GetRoleId() uuid.UUID {
	return m.RoleId
}

func (m *Role) SetRoleName(role_name string) {
	m.RoleName = role_name
}
func (m Role) GetRoleName() string {
	return m.RoleName
}

func (m *Role) SetCreatedBy(created_by uuid.UUID) {
	m.CreatedBy = created_by
}
func (m Role) GetCreatedBy() uuid.UUID {
	return m.CreatedBy
}

func (m *Role) SetCreatedDate(created_date time.Time) {
	m.CreatedDate = created_date
}
func (m Role) GetCreatedDate() time.Time {
	return m.CreatedDate
}

func (m *Role) SetUpdatedBy(updated_by uuid.UUID) {
	m.UpdatedBy = updated_by
}
func (m Role) GetUpdatedBy() uuid.UUID {
	return m.UpdatedBy
}

func (m *Role) SetUpdatedDate(updated_date time.Time) {
	m.UpdatedDate = updated_date
}
func (m Role) GetUpdatedDate() time.Time {
	return m.UpdatedDate
}
