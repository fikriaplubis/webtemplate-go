package jwt

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"

	"webtemplate/internal/models/dao"
	"webtemplate/internal/models/dto"
)

type JwtToken interface {
	CreateToken(user dao.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*dto.Token, error)
}

type jwtToken struct {
	secretKey []byte
}

func NewJWTMaker(secretKey string) *jwtToken {
	return &jwtToken{
		secretKey: []byte(secretKey),
	}
}

func (j *jwtToken) CreateToken(user dao.User, duration time.Duration) (string, error) {
	payload, err := dto.NewToken(user, duration)
	if err != nil {
		return "", err
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := at.SignedString(j.secretKey)
	if err != nil {
		log.Println("Login: Error while creating token")
		return "", err
	}

	return token, nil
}

func (j *jwtToken) VerifyToken(token string) (*dto.Token, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, dto.ErrInvalidToken
		}
		return []byte(j.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &dto.Token{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, dto.ErrExpiredToken) {
			return nil, dto.ErrExpiredToken
		}
		return nil, dto.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*dto.Token)
	if !ok {
		return nil, dto.ErrInvalidToken
	}

	return payload, nil
}
