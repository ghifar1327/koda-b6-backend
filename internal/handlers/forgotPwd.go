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

// RequestForgotPwd godoc
// @Summary Request forgot password
// @Description Send OTP code to user email for password reset
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.ForgotPasswordRequest true "User Email"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /auth/forgot-password [post]
func (h *ForgotPwdHandler) RequestForgotPwd(ctx *gin.Context) {
	var req dto.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid request body",
		})
		return
	}

	if err := h.service.RequestForgotPwd(ctx, req.Email); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "OTP code has been sent successfully",
	})
}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset user password using OTP code
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.ResetPwdRequest true "Reset Password Request"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /auth/reset-password [post]
func (h *ForgotPwdHandler) Resetpassword(ctx *gin.Context) {
	var req dto.ResetPwdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid request body",
		})
		return
	}
	if err := h.service.ResetPassword(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.Response{
		Success: true,
		Message: "reset password successfully",
	})
}
