package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: s,
	}
}


// GetAllTransactions godoc
// @Summary Get all transactions
// @Description Retrieve all transactions
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {array} models.Transaction
// @Failure 500 {object} map[string]string
// @Router /transactions [get]
func (h *TransactionHandler) GetAllTransactions(ctx *gin.Context) {
	Transactions, err := h.service.GetAllTransactions(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, Transactions)
}


// GetTransactionByID godoc
// @Summary Get transaction by ID
// @Description Retrieve transaction detail by UUID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID (UUID)"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransactionbyID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	Transaction, err := h.service.GetAllTransactionByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, Transaction)
}


// CreateTransaction godoc
// @Summary Create transaction
// @Description Create a new transaction with transaction details
// @Tags Transactions
// @Accept json
// @Produce json
// @Param request body dto.CreateTransactionRequest true "Transaction Data"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var req dto.CreateTransactionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid Request Body",
		})
		return
	}
	err := h.service.CreateTransaction(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Create Transaction successfully",
	})
}
