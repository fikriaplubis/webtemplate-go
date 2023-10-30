package dto

import (
	"net/http"

	"github.com/google/uuid"

	"webtemplate/internal/models/dao"
)

type RoleRequest struct {
	RoleName   string      `json:"role_name"`
	MenuAction []uuid.UUID `json:"menu_action_id"`
}

type RoleResponseRead struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       *[]dao.Role `json:"data"`
	TotalData  int         `json:"total_data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
}

type DataAffectedRole struct {
	Roles          *dao.Role             `json:"roles"`
	RoleMenuAction *[]dao.RoleMenuAction `json:"role_menu_action"`
}

type RoleResponseWrite struct {
	StatusCode   int               `json:"status_code"`
	Error        string            `json:"error"`
	Message      string            `json:"message"`
	DataRequest  RoleRequest       `json:"data_request"`
	DataAffected *DataAffectedRole `json:"data_affected"`
}

func (req RoleRequest) SuccessWrite(message string, data_affected *DataAffectedRole) RoleResponseWrite {
	return RoleResponseWrite{
		StatusCode:   http.StatusCreated,
		Error:        "nil",
		Message:      message,
		DataRequest:  req,
		DataAffected: data_affected,
	}
}

func (req RoleRequest) ErrorWrite(status_code int, err string, message string) RoleResponseWrite {
	return RoleResponseWrite{
		StatusCode:   status_code,
		Error:        err,
		Message:      message,
		DataRequest:  req,
		DataAffected: nil,
	}
}
