package menu

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"webtemplate/internal/models/dto"
	"webtemplate/internal/services/menu"
)

type Handler struct {
	Service menu.Service
}

func NewHandler(service menu.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Create(c *gin.Context) {
	var request dto.MenuRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response := request.ErrorWrite(http.StatusBadRequest, "Error Binding", err.Error())
		c.JSON(response.StatusCode, &response)
		return
	}

	response := h.Service.Create(request)
	c.JSON(response.StatusCode, &response)
}

// func (h *Handler) GetByMenuId(c *gin.Context) {
// 	menuId := c.Param("menu_id")

// 	res, err := h.Service.GetByMenuId(menuId)
// 	if err != nil {
// 		c.JSON(err.Status, gin.H{
// 			"data":    res,
// 			"message": err.Message,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data":    res,
// 		"message": err.Message,
// 	})
// }

func (h *Handler) GetAll(c *gin.Context) {
	response := h.Service.GetAll()
	c.JSON(response.StatusCode, &response)
}
