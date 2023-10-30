package dto

import (
	"net/http"

	"github.com/google/uuid"

	"webtemplate/internal/models/dao"
)

type RoleMenuActionRequest struct {
	RoleId       uuid.UUID `json:"role_id"`
	MenuActionId uuid.UUID `json:"menu_action_id"`
}

type RoleMenuActionRequestTransform struct {
	RoleId       uuid.UUID   `json:"role_id"`
	MenuActionId []uuid.UUID `json:"menu_action_id"`
}

type RoleMenuActionResponseWrite struct {
	StatusCode   int                            `json:"status_code"`
	Error        string                         `json:"error"`
	Message      string                         `json:"message"`
	DataRequest  RoleMenuActionRequestTransform `json:"data_request"`
	DataAffected *[]dao.RoleMenuAction          `json:"data_affected"`
}

func (req RoleMenuActionRequestTransform) SuccessWrite(message string, data_affected *[]dao.RoleMenuAction) RoleMenuActionResponseWrite {
	return RoleMenuActionResponseWrite{
		StatusCode:   http.StatusCreated,
		Error:        "nil",
		Message:      message,
		DataRequest:  req,
		DataAffected: data_affected,
	}
}

func (req RoleMenuActionRequestTransform) ErrorWrite(status_code int, err string, message string) RoleMenuActionResponseWrite {
	return RoleMenuActionResponseWrite{
		StatusCode:   status_code,
		Error:        err,
		Message:      message,
		DataRequest:  req,
		DataAffected: nil,
	}
}
