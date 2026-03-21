package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

// GetProducts godoc
// @Summary Get all products
// @Description Retrieve all products from database
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (h *ProductHandler) GetProducts(ctx *gin.Context) {
	products, err := h.service.GetAllProducts(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: "List Of Products",
		Results: products,
	})
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Retrieve a single product by its ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductbyID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid ID",
		})
	}
	product, err := h.service.GetProductByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: "product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.ResponseWrap{
		Success: true,
		Message: "Detail of Product",
		Results: product,
	})
}

// CreateProduct godoc
// @Summary Create new product
// @Description Create a new product in database
// @Tags Products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductRequest true "Product Data"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /products [post]
func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req dto.CreateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid Request Body",
		})
		return
	}
	err := h.service.CreateProduct(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Create product successfully",
	})
}

// GetVariantByIdProduct godoc
// @Summary Get variants by product ID
// @Description Get list of variants based on product ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ResponseWrap{results=[]models.Variant}
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /products/{id}/variants [get]
func (h *ProductHandler) GetVariantByIdProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid ID",
		})
		return
	}

	variants, err := h.service.GetVariantsByIdProduct(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	message := "List of variants"
	if len(variants) == 0 {
		message = "No variants found"
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrap{
		Success: true,
		Message: message,
		Results: variants,
	})
}

// GetSizeByIdProduct godoc
// @Summary Get Sizes by product ID
// @Description Get list of Sizes based on product ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ResponseWrap{results=[]models.Size}
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /products/{id}/sizes [get]
func (h *ProductHandler) GetSizesByIdProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseWrap{
			Success: false,
			Message: "invalid ID",
			Results: nil,
		})
		return
	}

	sizes, err := h.service.GetSizesByIdProduct(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(500, dto.ResponseWrap{
			Success: false,
			Message: err.Error(),
			Results: nil,
		})
		return
	}
	message := "list of sizes"

	if len(sizes) == 0 {
		message = "No variants found"
	}

	ctx.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: message,
		Results: sizes,
	})
}
