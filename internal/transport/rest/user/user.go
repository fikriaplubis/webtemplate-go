package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"webtemplate/internal/models/dto"
	"webtemplate/internal/services/user"
)

type Handler struct {
	Service user.Service
}

func NewHandler(service user.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Create(c *gin.Context) {
	var request dto.UserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response := request.ErrorWrite(http.StatusBadRequest, "Error Binding", err.Error())
		c.JSON(response.StatusCode, &response)
		return
	}

	response := h.Service.Create(request)
	c.JSON(response.StatusCode, &response)
}
