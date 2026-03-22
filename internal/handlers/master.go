package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MasterHandler struct {
	service *service.MasterService
}

func NewMasterHandler(service *service.MasterService) *MasterHandler {
	return &MasterHandler{service: service}
}

func isValidTable(table string) bool {
	switch table {
	case "sizes", "variants", "methods":
		return true
	}
	return false
}

// Create godoc
// @Summary Create master data
// @Description Create sizes, variants, or methods
// @Tags Master
// @Accept json
// @Produce json
// @Param table path string true "Table name (sizes/variants/methods)"
// @Param request body dto.CreateMasterRequest true "Request body"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /master/{table} [post]
func (h *MasterHandler) Create(ctx *gin.Context) {
	table := ctx.Param("table")

	if !isValidTable(table) {
		ctx.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid table",
		})
		return
	}

	var req dto.CreateMasterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	err := h.service.Create(ctx, table, req)
	if err != nil {
		ctx.JSON(500, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.Response{
		Success: true,
		Message: "Created",
	})
}

// GetAll godoc
// @Summary Get all master data
// @Description Get sizes, variants, or methods
// @Tags Master
// @Produce json
// @Param table path string true "Table name"
// @Success 200 {array} dto.Master
// @Router /master/{table} [get]
func (h *MasterHandler) GetAll(ctx *gin.Context) {
	table := ctx.Param("table")

	if !isValidTable(table) {
		ctx.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid table",
		})
		return
	}

	data, err := h.service.GetAll(ctx, table)
	if err != nil {
		ctx.JSON(500, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: fmt.Sprintf("List of %s", table),
		Results: data,
	})
}

// GetById godoc
// @Summary Get master by ID
// @Description Get detail sizes, variants, or methods
// @Tags Master
// @Produce json
// @Param table path string true "Table name"
// @Param id path int true "ID"
// @Success 200 {object} dto.Master
// @Failure 404 {object} dto.Response
// @Router /master/{table}/{id} [get]
func (h *MasterHandler) GetById(ctx *gin.Context) {
	table := ctx.Param("table")
	id, err := strconv.Atoi(ctx.Param("id"))
	
	if !isValidTable(table) {
		ctx.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid table",
		})
		return
	}

	if err != nil {
		ctx.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid Id",
		})
		return
	}

	data, err := h.service.GetById(ctx, table, id)
	if err != nil {
		ctx.JSON(404, dto.Response{
			Success: false,
			Message: "data not found",
		})
		return
	}

	ctx.JSON(200, dto.ResponseWrap{
		Success: true,
		Message: fmt.Sprintf("List of %s", table),
		Results: data,
	})
}

// Update godoc
// @Summary Update master data
// @Description Update sizes, variants, or methods
// @Tags Master
// @Accept json
// @Produce json
// @Param table path string true "Table name"
// @Param id path int true "ID"
// @Param request body dto.UpdateMasterRequest true "Request body"
// @Success 200 {object} dto.Response
// @Router /master/{table}/{id} [patch]
func (h *MasterHandler) Update(ctx *gin.Context) {
	table := ctx.Param("table")

	if !isValidTable(table) {
		ctx.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid table",
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid Id",
		})
		return
	}

	var req dto.UpdateMasterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	err = h.service.Update(ctx, table, id, req)
	if err != nil {
		ctx.JSON(500, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.Response{
		Success: true,
		Message: "Updated",
	})
}

// Delete godoc
// @Summary Delete master data
// @Description Delete sizes, variants, or methods
// @Tags Master
// @Produce json
// @Param table path string true "Table name"
// @Param id path int true "ID"
// @Success 200 {object} dto.Response
// @Router /master/{table}/{id} [delete]
func (h *MasterHandler) Delete(ctx *gin.Context) {
	table := ctx.Param("table")

	if !isValidTable(table) {
		ctx.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid table",
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, dto.Response{
			Success: false,
			Message: "Invalid Id",
		})
		return
	}

	err = h.service.Delete(ctx, table, id)
	if err != nil {
		ctx.JSON(500, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.Response{
		Success: true,
		Message: "Deleted",
	})
}
