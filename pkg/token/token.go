package token

import (
	"os"
	"webtemplate/pkg/random"

	"webtemplate/pkg/token/jwt"
	"webtemplate/pkg/token/paseto"
)

type tokenMethod struct {
	Paseto paseto.PasetoToken
	JWT    jwt.JwtToken
}

func NewService() *tokenMethod {
	return &tokenMethod{
		Paseto: paseto.NewPasetoMaker(random.RandomString(32)),
		JWT:    jwt.NewJWTMaker(os.Getenv("JWT_SECRET")),
	}
}
