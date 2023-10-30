package dto

import (
	// "time"

	"net/http"

	"github.com/google/uuid"

	"webtemplate/internal/models/dao"
	// "webtemplate/pkg/encrypt"
)

type UserRequest struct {
	RoleId   uuid.UUID `json:"role_id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

type UserResponseWrite struct {
	StatusCode   int         `json:"status_code"`
	Error        string      `json:"error"`
	Message      string      `json:"message"`
	DataRequest  UserRequest `json:"data_request"`
	DataAffected *dao.User   `json:"data_affected"`
}

func (req UserRequest) SuccessWrite(message string, data_affected *dao.User) UserResponseWrite {
	return UserResponseWrite{
		StatusCode:   http.StatusCreated,
		Error:        "nil",
		Message:      message,
		DataRequest:  req,
		DataAffected: data_affected,
	}
}

func (req UserRequest) ErrorWrite(status_code int, err string, message string) UserResponseWrite {
	return UserResponseWrite{
		StatusCode:   status_code,
		Error:        err,
		Message:      message,
		DataRequest:  req,
		DataAffected: nil,
	}
}

// func (u *UserResponse) EncryptUserId(userId string) {
// 	u.UserId = encrypt.Encrypt(userId)
// }
