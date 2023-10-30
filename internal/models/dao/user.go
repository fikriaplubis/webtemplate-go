package dao

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId      uuid.UUID `json:"user_id"`
	RoleId      uuid.UUID `json:"role_id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Role        *Role     `json:"role" gorm:"foreignKey:RoleId;references:RoleId"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	UpdatedDate time.Time `json:"updated_date"`
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) SetUserId(user_id uuid.UUID) {
	m.UserId = user_id
}
func (m User) GetUserId() uuid.UUID {
	return m.UserId
}

func (m *User) SetRoleId(role_id uuid.UUID) {
	m.RoleId = role_id
}
func (m User) GetRoleId() uuid.UUID {
	return m.RoleId
}

func (m *User) SetUsername(username string) {
	m.Username = username
}
func (m User) GetUsername() string {
	return m.Username
}

func (m *User) SetPassword(username string) {
	m.Password = username
}
func (m User) GetPassword() string {
	return m.Password
}

func (m *User) SetCreatedBy(created_by uuid.UUID) {
	m.CreatedBy = created_by
}
func (m User) GetCreatedBy() uuid.UUID {
	return m.CreatedBy
}

func (m *User) SetCreatedDate(created_date time.Time) {
	m.CreatedDate = created_date
}
func (m User) GetCreatedDate() time.Time {
	return m.CreatedDate
}

func (m *User) SetUpdatedBy(updated_by uuid.UUID) {
	m.UpdatedBy = updated_by
}
func (m User) GetUpdatedBy() uuid.UUID {
	return m.UpdatedBy
}

func (m *User) SetUpdatedDate(updated_date time.Time) {
	m.UpdatedDate = updated_date
}
func (m User) GetUpdatedDate() time.Time {
	return m.UpdatedDate
}
