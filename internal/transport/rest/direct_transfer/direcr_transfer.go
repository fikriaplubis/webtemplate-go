package direct_transfer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"webtemplate/internal/models/dto"
	"webtemplate/internal/services/direct_transfer"
)

type Handler struct {
	Service direct_transfer.Service
}

func NewHandler(service direct_transfer.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Create(c *gin.Context) {
	var request dto.DirectTransferRequest

	if err := c.ShouldBind(&request); err != nil {
		response := request.ErrorWrite(http.StatusBadRequest, "Error Binding", err.Error())
		c.JSON(response.StatusCode, &response)
		return
	}

	response := h.Service.Create(request)
	c.JSON(response.StatusCode, &response)
	return
}
