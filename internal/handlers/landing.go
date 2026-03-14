package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type LandingHandler struct {
	service *service.LandingService
}

func NewLandingHandler(s *service.LandingService) *LandingHandler {
	return &LandingHandler{
		service: s,
	}
}

func (h *LandingHandler) GetAllReviewProducts(ctx *gin.Context) {
	products, err := h.service.GetAllReviewProducts(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, dto.Response{
			Ok : false,
			Message: err.Error()})
		return
	}
	ctx.JSON(200, products)
}

func (h *LandingHandler) GetReviwProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	product, err := h.service.GetReviwProductByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Ok: false,
			Message: "review not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (h *LandingHandler) GetRecommendedProducts(ctx *gin.Context) {
	products, err := h.service.GetRecommendedProducts(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, dto.Response{
			Ok : false,
			Message: err.Error()})
		return
	}
	ctx.JSON(200, products)
}

func (h *LandingHandler) GetRecommendedProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	product, err := h.service.GetRecommendedProductByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Ok: false,
			Message: "review not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}