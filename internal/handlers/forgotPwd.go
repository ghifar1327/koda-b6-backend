package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPwdHandler struct {
	service *service.ForgotPwdService
}

func NewForgitPwdHandler(s *service.ForgotPwdService) *ForgotPwdHandler {
	return &ForgotPwdHandler{service: s}
}

func (h *ForgotPwdHandler) RequestForgotPwd(ctx *gin.Context) {
	var email string
	if err := ctx.ShouldBindJSON(&email); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Ok:      false,
			Message: "invalid request body",
		})
		return
	}

	if err := h.service.RequestForgotPwd(ctx, email); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Ok:      true,
		Message: "OTP code has been sent successfully",
	})
}

func (h *ForgotPwdHandler) Reretpassword(ctx *gin.Context) {
	var req dto.ResetPwdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Ok:      false,
			Message: "invalid request body",
		})
		return
	}
	if err := h.service.ResetPassword(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.Response{
		Ok:      true,
		Message: "reset password successfully",
	})
}
