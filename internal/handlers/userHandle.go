package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

// ====================================================================== get all user

// GetUsers godoc
// @Summary Get all users
// @Description Retrieve list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} dto.UserResponse
// @Failure 500 {object} dto.Response
// @Router /admin/users [get]
func (h *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := h.service.ReadAllUser(ctx.Request.Context())

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var responses []dto.UserResponse

	for _, user := range users {
		responses = append(responses, dto.UserResponse{
			Id:       user.Id,
			FullName: user.FullName,
			Email:    user.Email,
		})
	}

	ctx.JSON(http.StatusOK, responses)
}

// ==================================================================== get user by id

// GetUserById godoc
// @Summary Get user by ID
// @Description Retrieve a single user by UUID
// @Tags Users
// @Produce json
// @Param id path string true "User ID (UUID)"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /admin/users/{id} [get]
func (h *UserHandler) GetUserById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid uuid",
		})
		return
	}

	user, err := h.service.ReadByIdUser(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: "user not found",
		})
		return
	}
	result := dto.UserResponse{
		Id:       user.Id,
		FullName: user.FullName,
		Email:    user.Email,
	}
	ctx.JSON(http.StatusOK, result)
}

// =============================================================================== Update User

// UpdateUser godoc
// @Summary Update user
// @Description Update user data by UUID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID (UUID)"
// @Param user body dto.UpdateUsersRequest true "Update User Data"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /admin/users/{id} [patch]
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid user id",
		})
		return
	}
	var req dto.UpdateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.UpdateUser(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "User updated successfully",
	})
}

//==================================================================================== Delete User

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by UUID
// @Tags Users
// @Produce json
// @Param id path string true "User ID (UUID)"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /admin/users/{id} [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid user id",
		})
		return
	}

	if err := h.service.DeleteUser(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "failed to delete user",
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "user delete successfully",
	})
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
func (h *UserHandler) Register(ctx *gin.Context) {
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
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "User registered successfully",
	})
}
