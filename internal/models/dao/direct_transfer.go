package dao

import (
	"time"

	"github.com/google/uuid"
)

type DirectTransfer struct {
	NoReff           string    `json:"no_reff"`
	FileName         string    `json:"file_name"`
	TransactionDate  time.Time `json:"transaction_date"`
	Company          string    `json:"company"`
	Description      string    `json:"description"`
	AccountNo        string    `json:"account_no"`
	Currency         string    `json:"currency"`
	Branch           string    `json:"branch"`
	Amount           float64   `json:"amount"`
	StatusMove       int       `json:"status_move"`
	NumbersOfAccount int       `json:"numbers_of_account"`
	UploadedBy       uuid.UUID `json:"uploaded_by"`
	UploadedDate     time.Time `json:"uploaded_date"`
	IPUpload         string    `json:"ip_upload"`
}

func (m *DirectTransfer) TableName() string {
	return "direct_transfer"
}

func (m *DirectTransfer) SetNoReff(no_reff string) {
	m.NoReff = no_reff
}
func (m DirectTransfer) GetNoReff() string {
	return m.NoReff
}
