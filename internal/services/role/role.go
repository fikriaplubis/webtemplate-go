package role

import (
	"net/http"

	"gorm.io/gorm"

	"webtemplate/internal/models/dao"
	"webtemplate/internal/models/dto"
	"webtemplate/internal/services/role_menu_action"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
)

type Service interface {
	Create(request dto.RoleRequest) dto.RoleResponseWrite
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{db}
}

func (s *service) Create(request dto.RoleRequest) dto.RoleResponseWrite {
	var newRole dao.Role

	if err := mapper.AutoMapper(&request, &newRole); err != nil {
		return request.ErrorWrite(http.StatusInternalServerError, "Error mapper", err.Error())
	}

	// Set others value
	roleId := uuid.New()
	newRole.SetRoleId(roleId)

	// Begin Transaction
	tx := s.db.Begin()
	rolemenuactionServices := role_menu_action.NewService(tx)

	if err := tx.Create(request).Error; err != nil {
		tx.Rollback()
		return request.ErrorWrite(http.StatusInternalServerError, "Error when trying to save data in the roles table", err.Error())
	}

	createRoleMenuAction := rolemenuactionServices.BulkCreate(request, roleId)
	if createRoleMenuAction.Error != "nil" {
		tx.Rollback()
		return request.ErrorWrite(createRoleMenuAction.StatusCode, createRoleMenuAction.Error, createRoleMenuAction.Message)
	}

	dataAffected := &dto.DataAffectedRole{
		Roles:          &newRole,
		RoleMenuAction: createRoleMenuAction.DataAffected,
	}

	tx.Commit()
	return request.SuccessWrite("Success create new role", dataAffected)
}
