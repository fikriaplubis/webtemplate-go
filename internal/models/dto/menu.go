package dto

import (
	"net/http"

	"github.com/google/uuid"

	"webtemplate/internal/models/dao"
)

type MenuRequest struct {
	MenuName     string    `json:"menu_name"`
	ParentMenuId uuid.UUID `json:"parent_menu_id"`
	Actions      []int     `json:"actions"`
}

type MenuResponseRead struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       *[]dao.Menu `json:"data"`
	TotalData  int         `json:"total_data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
}

type DataAffectedMenu struct {
	Menus      *dao.Menu         `json:"menus"`
	MenuAction *[]dao.MenuAction `json:"menu_action"`
}

type MenuResponseWrite struct {
	StatusCode   int               `json:"status_code"`
	Error        string            `json:"error"`
	Message      string            `json:"message"`
	DataRequest  MenuRequest       `json:"data_request"`
	DataAffected *DataAffectedMenu `json:"data_affected"`
}

func (req MenuRequest) SuccessWrite(message string, data_affected *DataAffectedMenu) MenuResponseWrite {
	return MenuResponseWrite{
		StatusCode:   http.StatusCreated,
		Error:        "nil",
		Message:      message,
		DataRequest:  req,
		DataAffected: data_affected,
	}
}

func (req MenuRequest) ErrorWrite(status_code int, err string, message string) MenuResponseWrite {
	return MenuResponseWrite{
		StatusCode:   status_code,
		Error:        err,
		Message:      message,
		DataRequest:  req,
		DataAffected: nil,
	}
}

// type MenuResponseTransform struct {
// 	MenuId      uuid.UUID                      `json:"menu_id"`
// 	MenuName    string                         `json:"menu_name"`
// 	ParentMenu  *dao.Menu                      `json:"parent_menu"`
// 	Actions     []*MenuActionResponseTransform `json:"actions"`
// 	CreatedBy   uuid.UUID                      `json:"created_by"`
// 	CreatedDate time.Time                      `json:"created_date"`
// 	UpdatedBy   uuid.UUID                      `json:"updated_by"`
// 	UpdatedDate time.Time                      `json:"updated_date"`
// }

// func (ma *MenuResponses) TransformResponse() []*MenuResponseTransform {
// 	var transformed []*MenuResponseTransform

// 	for _, menu := range ma.Data {
// 		transform := &MenuResponseTransform{
// 			MenuId:      menu.MenuId,
// 			MenuName:    menu.MenuName,
// 			ParentMenu:  menu.ParentMenu,
// 			Actions:     actions,
// 			CreatedBy:   menu.CreatedBy,
// 			CreatedDate: menu.CreatedDate,
// 			UpdatedBy:   menu.UpdatedBy,
// 			UpdatedDate: menu.UpdatedDate,
// 		}

// 		transformed = append(transformed, transform)
// 	}

// 	return transformed
// }
