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

func (h *TransactionHandler) GetTransactions(ctx *gin.Context) {
	Transactions, err := h.service.GetAllTransactions(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, Transactions)
}

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

func (h *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var req dto.CreateRransactionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Ok:      false,
			Message: "Invalid Request Body",
		})
		return
	}
	err := h.service.CreateTransaction(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Ok:      true,
		Message: "Create Transaction successfully",
	})
}
