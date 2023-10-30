package auth

import (
	"net/http"

	"webtemplate/internal/models/dto"
	"webtemplate/internal/services/auth"
	restErrors "webtemplate/pkg/errors/rest"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service auth.Service
}

func NewHandler(service auth.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Login(c *gin.Context) {
	var loginRequestDTO dto.LoginRequest
	loginResponseDTO := &dto.LoginResponse{}

	if err := c.ShouldBindJSON(&loginRequestDTO); err != nil {
		errRest := restErrors.NewBadRequestError(err.Error())
		c.JSON(errRest.Status, gin.H{"message": "Bad Request"})
		return
	}

	loginResponseDTO, err := h.Service.Login(loginRequestDTO)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    loginResponseDTO,
	})
}
