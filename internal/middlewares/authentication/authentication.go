package authentication

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	// restErrors "webtemplate/pkg/errors/rest"
	"webtemplate/pkg/token"
)

func Authentication(c *gin.Context) {
	tokenHeader := c.Request.Header.Get("Authorization")
	str := tokenHeader
	str = strings.ReplaceAll(str, "Bearer ", "")

	tokenService := token.NewService()
	payload, err := tokenService.Paseto.VerifyToken(str)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		c.Abort()
		return
	}

	username := payload.Username
	roleId := payload.RoleId

	c.Set("username", username)
	c.Set("roleId", roleId)

	c.Next()
}
