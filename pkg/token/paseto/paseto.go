package paseto

import (
	"time"

	"github.com/o1egl/paseto"

	"webtemplate/internal/models/dao"
	"webtemplate/internal/models/dto"
)

type PasetoToken interface {
	CreateToken(user dao.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*dto.Token, error)
}

type pasetoToken struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) *pasetoToken {
	return &pasetoToken{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
}

func (p *pasetoToken) CreateToken(user dao.User, duration time.Duration) (string, error) {
	payload, err := dto.NewToken(user, duration)
	if err != nil {
		return "", err
	}

	token, err := p.paseto.Encrypt(p.symmetricKey, payload, nil)
	return token, err
}

func (p *pasetoToken) VerifyToken(token string) (*dto.Token, error) {
	payload := &dto.Token{}

	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, dto.ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
