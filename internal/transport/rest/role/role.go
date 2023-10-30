package role

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"webtemplate/internal/models/dto"
	"webtemplate/internal/services/role"
)

type Handler struct {
	Service role.Service
}

func NewHandler(service role.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Create(c *gin.Context) {
	var request dto.RoleRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response := request.ErrorWrite(http.StatusBadRequest, "Error Binding", err.Error())
		c.JSON(response.StatusCode, &response)
		return
	}

	response := h.Service.Create(request)
	c.JSON(response.StatusCode, &response)
}
