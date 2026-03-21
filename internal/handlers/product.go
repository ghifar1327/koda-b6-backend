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
			Success: true,
			Message: "Invalid ID",
		})
	}
	product, err := h.service.GetAllProductByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: true,
			Message: "user not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
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
