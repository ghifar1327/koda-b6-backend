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


// GetAllReviewProductsLanding godoc
// @Summary Get all review products
// @Description Retrieve all review products for landing page
// @Tags Landing
// @Accept json
// @Produce json
// @Success 200 {array} dto.ResponseWrap
// @Failure 500 {object} dto.Response
// @Router /reviews [get]
func (h *LandingHandler) GetAllReviewProductsLanding(ctx *gin.Context) {
	products, err := h.service.GetAllReviewProductsLanding(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, dto.Response{
			Success: false,
			Message: err.Error()})
		return
	}
	ctx.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: "List of review products",
		Results: products,
	})
}


// GetReviewProductLandingByID godoc
// @Summary Get review product by ID
// @Description Retrieve review product by ID
// @Tags Landing
// @Accept json
// @Produce json
// @Param id path int true "Review Product ID"
// @Success 200 {object} dto.ResponseWrap
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /reviews/{id} [get]
func (h *LandingHandler) GetReviewProductLandingByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	product, err := h.service.GetReviewProductLandingByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: "review not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.ResponseWrap{
		Success: true,
		Message: "ReviewProduct found",
		Results: product,
	})
}


// GetRecommendedProducts godoc
// @Summary Get recommended products
// @Description Retrieve recommended products for landing page
// @Tags Landing
// @Accept json
// @Produce json
// @Success 200 {array} dto.ResponseWrap
// @Failure 500 {object} dto.Response
// @Router /recommended-product [get]
func (h *LandingHandler) GetRecommendedProducts(ctx *gin.Context) {
	products, err := h.service.GetRecommendedProducts(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, dto.Response{
			Success: false,
			Message: err.Error()})
		return
	}
	ctx.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: "List of recommended products",
		Results: products,
	})
}


// GetRecommendedProductByID godoc
// @Summary Get recommended product by ID
// @Description Retrieve recommended product detail by ID
// @Tags Landing
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ResponseWrap
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /recommended-product/{id} [get]
func (h *LandingHandler) GetRecommendedProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	product, err := h.service.GetRecommendedProductByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: "review not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.ResponseWrap{
		Success: true,
		Message: "Recommended product found",
		Results: product,
	})
}
