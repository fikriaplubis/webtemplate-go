package dto

import (
	"mime/multipart"
	"net/http"
	"time"

	"webtemplate/internal/models/dao"
)

type DirectTransferRequest struct {
	File            *multipart.FileHeader `form:"file" json:"file"`
	TransactionDate time.Time             `form:"transaction_date" binding:"required" json:"transaction_date"`
	Company         string                `form:"company" binding:"required" json:"company"`
	Description     string                `form:"description" binding:"required" json:"description"`
	AccountNo       string                `form:"account_no" binding:"required" json:"account_no"`
	Currency        string                `form:"currency" binding:"required" json:"currency"`
	Amount          float64               `form:"amount" binding:"required" json:"amount"`
}

type DirectTransferResponseWrite struct {
	StatusCode   int                   `json:"status_code"`
	Error        string                `json:"error"`
	Message      string                `json:"message"`
	DataRequest  DirectTransferRequest `json:"data_request"`
	DataAffected *dao.DirectTransfer   `json:"data_affected"`
}

func (req DirectTransferRequest) Get() DirectTransferRequest {
	return req
}

func (req DirectTransferRequest) SuccessWrite(message string, data_affected *dao.DirectTransfer) DirectTransferResponseWrite {
	return DirectTransferResponseWrite{
		StatusCode:   http.StatusCreated,
		Error:        "nil",
		Message:      message,
		DataRequest:  req,
		DataAffected: data_affected,
	}
}

func (req DirectTransferRequest) ErrorWrite(status_code int, err string, message string) DirectTransferResponseWrite {
	return DirectTransferResponseWrite{
		StatusCode:   status_code,
		Error:        err,
		Message:      message,
		DataRequest:  req,
		DataAffected: nil,
	}
}
