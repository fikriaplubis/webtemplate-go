package dto

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"webtemplate/internal/models/dao"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Token struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	RoleId    uuid.UUID `json:"role_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewToken(user dao.User, duration time.Duration) (*Token, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Token{
		ID:        tokenID,
		Username:  user.Username,
		RoleId:    user.RoleId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Token) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
