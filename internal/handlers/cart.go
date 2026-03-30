package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type CartHandler struct{
	service *service.CartService
}

func NewCartHandler(s *service.CartService) *CartHandler{
	return &CartHandler{service: s}
}


func (h *CartHandler) AddCart(c *gin.Context) {
	var req dto.ADDCartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	res, err := h.service.AddToCart(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: "Add cart successfully",
		Results: res,
	})
}