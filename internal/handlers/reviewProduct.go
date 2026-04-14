package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewProductHandler struct {
	service *service.ReviewProductService
}

func NewReviewProductHandler(s *service.ReviewProductService) *ReviewProductHandler {
	return &ReviewProductHandler{
		service: s,
	}
}

// =================================================================================== Create Review Product

// CreateReviewProduct godoc
// @Summary Create review product
// @Description Create a new review product
// @Tags ReviewProducts
// @Accept json
// @Produce json
// @Param request body dto.CreateReviewProductRequest true "Review Product Data"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /review-products [post]
func (h *ReviewProductHandler) CreateReviewProduct(ctx *gin.Context) {
	var req dto.CreateReviewProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid Request Body",
		})
		return
	}
	err := h.service.CreateReviewProduct(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Create ReviewProduct successfully",
	})
}

//============================================================================================ Get Review Product

// GetReviewProducts godoc
// @Summary Get all review products
// @Description Retrieve all review products from database
// @Tags ReviewProducts
// @Accept json
// @Produce json
// @Success 200 {array} models.ReviewProduct
// @Failure 500 {object} map[string]string
// @Router /review-product [get]
func (h *ReviewProductHandler) GetAllReviewProducts(ctx *gin.Context) {
	ReviewProducts, err := h.service.GetAllReviewProducts(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: "List of review products",
		Results: ReviewProducts,
	})
}

//================================================================================================== Get review product by id

// GetReviewProductByID godoc
// @Summary Get review product by ID
// @Description Retrieve a single review product by its ID
// @Tags ReviewProducts
// @Accept json
// @Produce json
// @Param id path int true "Review Product ID"
// @Success 200 {object} models.ReviewProduct
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /review-product/{id} [get]
func (h *ReviewProductHandler) GetReviewProductbyID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	ReviewProduct, err := h.service.GetAllReviewProductByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "ReviewProduct not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.ResponseWrap{
		Success: true,
		Message: "ReviewProduct found",
		Results: ReviewProduct,
	})
}

//================================================================================================== Update review Product

// CreateReviewProduct godoc
// @Summary Create review product
// @Description Create a new review product
// @Tags ReviewProducts
// @Accept json
// @Produce json
// @Param request body models.ReviewProduct true "Review Product Data"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /review-product [post]
func (h *ReviewProductHandler) UpdateReviewProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid ReviewProduct id",
		})
		return
	}
	var req dto.UpdateReviewProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.UpdateReviewProduct(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: "ReviewProduct not found or update failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "ReviewProduct updated successfully",
	})
}

//================================================================================ Delete Review Product

// DeleteReviewProduct godoc
// @Summary Delete review product
// @Description Delete review product by ID
// @Tags ReviewProducts
// @Accept json
// @Produce json
// @Param id path int true "Review Product ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /review-product/{id} [delete]
func (h *ReviewProductHandler) DeleteReviewProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid ReviewProduct id",
		})
		return
	}

	if err := h.service.DeleteReviewProduct(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "failed to delete ReviewProduct",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "ReviewProduct delete successfully",
	})
}
