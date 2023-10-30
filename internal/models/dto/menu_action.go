package dto

import (
	"net/http"

	"github.com/google/uuid"

	"webtemplate/internal/models/dao"
)

type MenuActionRequest struct {
	MenuId   uuid.UUID `json:"menu_id"`
	ActionId int       `json:"action_id"`
}

type MenuActionRequestTransform struct {
	MenuId   uuid.UUID `json:"menu_id"`
	ActionId []int     `json:"action_id"`
}

type MenuActionResponseRead struct {
	StatusCode int               `json:"status_code"`
	Message    string            `json:"message"`
	Data       *[]dao.MenuAction `json:"data"`
	TotalData  int               `json:"total_data"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
}

type MenuActionResponseWrite struct {
	StatusCode   int                        `json:"status_code"`
	Error        string                     `json:"error"`
	Message      string                     `json:"message"`
	DataRequest  MenuActionRequestTransform `json:"data_request"`
	DataAffected *[]dao.MenuAction          `json:"data_affected"`
}

func (req MenuActionRequestTransform) SuccessWrite(message string, data_affected *[]dao.MenuAction) MenuActionResponseWrite {
	return MenuActionResponseWrite{
		StatusCode:   http.StatusCreated,
		Error:        "nil",
		Message:      message,
		DataRequest:  req,
		DataAffected: data_affected,
	}
}

func (req MenuActionRequestTransform) ErrorWrite(status_code int, err string, message string) MenuActionResponseWrite {
	return MenuActionResponseWrite{
		StatusCode:   status_code,
		Error:        err,
		Message:      message,
		DataRequest:  req,
		DataAffected: nil,
	}
}

// type MenuActionResponseTransform struct {
// 	MenuActionId uuid.UUID   `json:"menu_action_id"`
// 	Action       *dao.Action `json:"action"`
// 	CreatedBy    uuid.UUID   `json:"created_by"`
// 	CreatedDate  time.Time   `json:"created_date"`
// 	UpdatedBy    uuid.UUID   `json:"updated_by"`
// 	UpdatedDate  time.Time   `json:"updated_date"`
// }

// func (mar *MenuActionResponses) TransformResponse() []*MenuActionResponseTransform {
// 	var transformed []*MenuActionResponseTransform

// 	for _, menuAction := range mar.Data {
// 		transform := &MenuActionResponseTransform{
// 			MenuActionId: menuAction.ID,
// 			Action:       menuAction.Action,
// 			CreatedBy:    menuAction.CreatedBy,
// 			CreatedDate:  menuAction.CreatedDate,
// 			UpdatedBy:    menuAction.UpdatedBy,
// 			UpdatedDate:  menuAction.UpdatedDate,
// 		}

// 		transformed = append(transformed, transform)
// 	}

// 	return transformed
// }
