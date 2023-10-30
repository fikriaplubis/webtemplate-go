package menu_action

import (
	"net/http"

	"gorm.io/gorm"

	"webtemplate/internal/models/dao"
	"webtemplate/internal/models/dto"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
)

type Service interface {
	BulkCreate(request dto.MenuRequest, menuId uuid.UUID) dto.RoleMenuActionRequest
	// GetByMenuId(menuId string) (dto.MenuActionResponsesTransform, *restErrors.RestError)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{db}
}

func (s *service) BulkCreate(request dto.MenuRequest, menuId uuid.UUID) dto.MenuActionResponseWrite {
	var newMenuAction dao.MenuAction
	newMenuActions := []dao.MenuAction{}

	var actions []int

	for _, action := range request.Actions {
		menuactionDTO := dto.MenuActionRequest{
			MenuId:   menuId,
			ActionId: action,
		}

		actions = append(actions, action)
		_ = mapper.AutoMapper(&menuactionDTO, &newMenuAction)

		// Set others value
		newMenuAction.SetId(uuid.New())

		newMenuActions = append(newMenuActions, newMenuAction)
	}

	menuActionRequest := dto.MenuActionRequestTransform{
		MenuId:   menuId,
		ActionId: actions,
	}

	if err := s.db.Create(newMenuActions).Error; err != nil {
		return menuActionRequest.ErrorWrite(http.StatusInternalServerError, "Error when trying to save data in the menu_action table", err.Error())
	}

	return menuActionRequest.SuccessWrite("Success create new menu-action", &newMenuActions)
}

// func (s *service) GetByMenuId(menuId string) (dto.MenuActionResponsesTransform, *restErrors.RestError) {
// 	var menuactions dto.MenuActionResponses

// 	getMenuAction := s.db.Where(`menu_id = ?`, menuId).Find(&menuactions.Data)
// 	if getMenuAction.Error != nil {
// 		return dto.MenuActionResponsesTransform{}, restErrors.NewBadRequestError(getMenuAction.Error.Error())
// 	}

// 	result := menuactions.TransformResponse()

// 	return result, nil
// }
