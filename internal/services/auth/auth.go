package auth

import (
	"log"
	"time"

	"gorm.io/gorm"

	"webtemplate/internal/models/dao"
	"webtemplate/internal/models/dto"
	"webtemplate/pkg/errors/rest"
	"webtemplate/pkg/hash"
	"webtemplate/pkg/token"
)

type Service interface {
	Login(loginRequest dto.LoginRequest) (*dto.LoginResponse, *rest.Error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{db}
}

func (s *service) Login(loginRequest dto.LoginRequest) (*dto.LoginResponse, *rest.Error) {
	var user dao.User

	// Find User by Username
	res := s.db.Where("username = ?", loginRequest.Username).Find(&user)

	if res.Error != nil {
		log.Println(res.Error)
		return nil, rest.NewInternalServerError("Internal server error")
	}

	// Check if User Exist
	if user.Username == "" {
		return nil, rest.NewBadRequestError("User or password are not valid")
	}

	// Check if Password is Valid
	if !hash.CheckPasswordHash(loginRequest.Password, user.Password) {
		return nil, rest.NewBadRequestError("User or password are not valid")
	}

	tokenService := token.NewService()
	pasetoToken, err := tokenService.Paseto.CreateToken(user, time.Minute)
	if err != nil {
		return nil, rest.NewInternalServerError("Internal server error")
	}

	loginResponse := &dto.LoginResponse{
		Token: pasetoToken,
	}

	return loginResponse, nil
}
