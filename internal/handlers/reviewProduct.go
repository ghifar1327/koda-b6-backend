package handlers

import (
	"backend/internal/dto"
	"backend/internal/models"
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

func (h *ReviewProductHandler) GetReviewProducts(ctx *gin.Context) {
	ReviewProducts, err := h.service.GetAllReviewProducts(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, ReviewProducts)
}

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
	ctx.JSON(http.StatusOK, ReviewProduct)
}

func (h *ReviewProductHandler) UpdateReviewProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ReviewProduct id",
		})
		return
	}
	var req dto.UpdateReviewProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": err.Error(),
		})
		return
	}

	if err := h.service.UpdateReviewProduct(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "ReviewProduct not found or update failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "ReviewProduct updated successfully",
	})
}

func (h *ReviewProductHandler) CreateReviewProduct(ctx *gin.Context) {
	var req models.ReviewProduct

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Ok:      false,
			Message: "Invalid Request Body",
		})
		return
	}
	err := h.service.CreateReviewProduct(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Ok:      true,
		Message: "Create ReviewProduct successfully",
	})
}

func (h *ReviewProductHandler) DeleteReviewProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ReviewProduct id",
		})
		return
	}

	if err := h.service.DeleteReviewProduct(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "failed to delete ReviewProduct",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ReviewProduct delete successfully",
	})
}
