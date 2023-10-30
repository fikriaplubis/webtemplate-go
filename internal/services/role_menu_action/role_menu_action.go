package role_menu_action

import (
	"net/http"

	"gorm.io/gorm"

	"webtemplate/internal/models/dao"
	"webtemplate/internal/models/dto"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
)

type Service interface {
	BulkCreate(request dto.RoleRequest, menuId uuid.UUID) dto.RoleMenuActionResponseWrite
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{db}
}

func (s *service) BulkCreate(request dto.RoleRequest, roleId uuid.UUID) dto.RoleMenuActionResponseWrite {
	var newRoleMenuAction dao.RoleMenuAction
	newRoleMenuActions := []dao.RoleMenuAction{}

	var menuactions []uuid.UUID

	for _, menuaction := range request.MenuAction {
		rolemenuactionDTO := dto.RoleMenuActionRequest{
			RoleId:       roleId,
			MenuActionId: menuaction,
		}

		menuactions = append(menuactions, menuaction)
		_ = mapper.AutoMapper(&rolemenuactionDTO, &newRoleMenuAction)

		// Set others value
		newRoleMenuAction.SetId(uuid.New())

		newRoleMenuActions = append(newRoleMenuActions, newRoleMenuAction)
	}

	roleMenuActionRequest := dto.RoleMenuActionRequestTransform{
		RoleId:       roleId,
		MenuActionId: menuactions,
	}

	if err := s.db.Create(newRoleMenuActions).Error; err != nil {
		return roleMenuActionRequest.ErrorWrite(http.StatusInternalServerError, "Error when trying to save data in the role_menu_action table", err.Error())
	}

	return roleMenuActionRequest.SuccessWrite("Success create new role-menu-action", &newRoleMenuActions)
}
