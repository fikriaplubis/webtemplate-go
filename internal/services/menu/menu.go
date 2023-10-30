package menu

import (
	"net/http"

	"gorm.io/gorm"

	"webtemplate/internal/models/dao"
	"webtemplate/internal/models/dto"
	"webtemplate/internal/services/menu_action"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
)

type Service interface {
	Create(request dto.MenuRequest) dto.MenuResponseWrite
	// GetByMenuId(menuId string) (dto.MenuActionResponsesTransform, *restErrors.RestError)
	GetAll() dto.MenuResponseRead
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{db}
}

func (s *service) Create(request dto.MenuRequest) dto.MenuResponseWrite {
	var newMenu dao.Menu

	if err := mapper.AutoMapper(&request, &newMenu); err != nil {
		return request.ErrorWrite(http.StatusInternalServerError, "Error Mapper", err.Error())
	}

	// Set others value
	menuId := uuid.New()
	newMenu.SetMenuId(menuId)

	// Begin Transaction
	tx := s.db.Begin()
	menuactionServices := menu_action.NewService(tx)

	if err := tx.Create(newMenu).Error; err != nil {
		tx.Rollback()
		return request.ErrorWrite(http.StatusInternalServerError, "Error when trying to save data in the menus table", err.Error())
	}

	createMenuAction := menuactionServices.BulkCreate(request, menuId)
	if createMenuAction.Error != "nil" {
		tx.Rollback()
		return request.ErrorWrite(createMenuAction.StatusCode, createMenuAction.Message, createMenuAction.Message)
	}

	dataAffected := &dto.DataAffectedMenu{
		Menus:      &newMenu,
		MenuAction: createMenuAction.DataAffected,
	}

	tx.Commit()
	return request.SuccessWrite("Success create new menu", dataAffected)
}

// func (s *service) GetByMenuId(menuId string) (dto.MenuActionResponsesTransform, *restErrors.RestError) {
// 	menuactionServices := menu_action.NewService(s.db)

// 	result, err := menuactionServices.GetByMenuId(menuId)
// 	if err != nil {
// 		return dto.MenuActionResponsesTransform{}, restErrors.NewBadRequestError(err.Message)
// 	}

// 	return result, nil
// }

func (s *service) GetAll() dto.MenuResponseRead {
	var response dto.MenuResponseRead

	result := s.db.Find(&response.Data)
	if result.Error != nil {
		return dto.MenuResponseRead{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when trying to fetch data in the menus table",
			Data:       response.Data,
			TotalData:  0,
			Page:       0,
			Limit:      0,
		}
	}

	return dto.MenuResponseRead{
		StatusCode: http.StatusOK,
		Message:    "Success fetch data",
		Data:       response.Data,
		TotalData:  0,
		Page:       0,
		Limit:      0,
	}
}
