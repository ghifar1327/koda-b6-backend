package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

// ==================================================================== register

// Register godoc
// @Summary Register new user
// @Description Create a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "Register Request"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /auth/register [post]
func (h *AuthHandler) Register(ctx *gin.Context) {
	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	err := h.service.Register(ctx, req)
	if err != nil {

		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				ctx.JSON(400, dto.Response{
					Success: false,
					Message: "Email already registered",
				})
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Internal server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "User registered successfully",
	})
}

// Login godoc
// @Summary Login new user
// @Description Create a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body dto.LoginRequest true "Login Request"
// @Success 200 {object} dto.ResponseToken
// @Failure 400 {object} dto.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err := h.service.GetUserBYEmail(ctx.Request.Context(), req.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	token, err := h.service.Login(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.ResponseToken{
		Success: true,
		Message: "Login Success",
		Token:   token,
		User: dto.UserResponse{
			Id:       user.Id,
			Picture:  user.Picture.String,
			Email:    user.Email,
			FullName: user.FullName,
			RoleId:   user.RoleId,
		},
	})
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
func (h *AuthHandler) RequestForgotPwd(ctx *gin.Context) {
	var req dto.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid request body",
		})
		return
	}

	if err := h.service.RequestForgotPwd(ctx.Request.Context(), req.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
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
// @Router /auth/reset-password [patch]
func (h *AuthHandler) ResetPassword(ctx *gin.Context) {
	var req dto.ResetPwdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid request body",
		})
		return
	}
	if err := h.service.ResetPassword(ctx.Request.Context(), req); err != nil {
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

// UpdateProfile godoc
// @Summary Update Profile
// @Description Update Profile data by UUID
// @Tags Profile
// @Accept json
// @Produce json
// @Param id path string true "User ID (UUID)"
// @Param user body dto.UpdateProfileRequest true "Update User Data"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /auth/{id}/update [patch]
func (h *AuthHandler) UpdateProfile(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid user id",
		})
		return
	}
	var req dto.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.UpdateProfile(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Provile updated successfully",
	})
}
