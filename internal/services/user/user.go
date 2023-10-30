package user

import (
	"net/http"

	"gorm.io/gorm"

	"webtemplate/internal/models/dao"
	"webtemplate/internal/models/dto"
	"webtemplate/pkg/hash"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
)

type Service interface {
	Create(request dto.UserRequest) dto.UserResponseWrite
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{db}
}

func (s *service) Create(request dto.UserRequest) dto.UserResponseWrite {
	var newUser dao.User

	if err := mapper.AutoMapper(&request, &newUser); err != nil {
		return request.ErrorWrite(http.StatusInternalServerError, "Error mapper", err.Error())
	}

	// Set others value
	newUser.SetUserId(uuid.New())
	newUser.SetPassword(hash.HashPassword(request.Password))

	if err := s.db.Create(newUser).Error; err != nil {
		return request.ErrorWrite(http.StatusInternalServerError, "Error when trying to save data in the users table", err.Error())
	}

	return request.SuccessWrite("Success create new user", &newUser)
}
