package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(s *service.CartService) *CartHandler {
	return &CartHandler{service: s}
}

func (h *CartHandler) GetCartByUserId(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid Id",
		})
		return
	}

	cart, err := h.service.GetCartByUserId(c.Request.Context(), userId)
	if err != nil {
		c.JSON(500, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: "List of cart items",
		Results: cart,
	})
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

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, dto.Response{
			Success: false,
			Message: "User not authenticated",
		})
		return
	}

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		c.JSON(500, dto.Response{
			Success: false,
			Message: "Invalid authenticated user",
		})
		return
	}

	req.UserID = userID

	cart, err := h.service.AddToCart(c.Request.Context(), req)
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
		Results: cart,
	})
}

func (h *CartHandler) DeleteCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid Id",
		})
		return
	}

	items, err := h.service.DeleteCart(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: "Delete cart successfully",
		Results: items,
	})

}
