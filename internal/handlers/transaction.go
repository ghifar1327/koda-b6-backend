package handlers

import (
	"backend/internal/dto"
	"backend/internal/models"
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
		ctx.JSON(500, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: "List of Transaction",
		Results: Transactions,
	})
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
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid ID",
		})
	}
	Transaction, err := h.service.GetAllTransactionByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: "transaction not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.ResponseWrap{
		Success: true,
		Message: "Data of transaction",
		Results: Transaction,
	})
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


// UpdateTransaction godoc
// @Summary Update transaction
// @Description Update transaction by ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID (UUID)"
// @Param request body models.Transaction true "Updated Transaction Data"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router admin/transactions/{id} [patch]
func (h *TransactionHandler) UpdateTransaction(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid ID",
		})
		return
	}
	
	var req models.Transaction

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.UpdateTransaction(ctx.Request.Context(), id , req); err != nil{
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Transactions updated successfully",
	})
}


// DeleteTransaction godoc
// @Summary Delete transaction
// @Description Delete transaction by ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID (UUID)"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router admin/transactions/{id} [delete]
func (h *TransactionHandler) DeleteTransaction(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid Transaction Id",
		})
		return
	}

	if err := h.service.DeleteTransaction(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "failed to delete Transaction",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Transaction delete successfully",
	})
}