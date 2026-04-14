package handlers

import (
	"backend/internal/dto"
	"backend/internal/service"
	"backend/internal/utils"
	"fmt"
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
		fmt.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var responses []dto.UserResponse

	for _, user := range users {
		user.Picture = utils.BuildImageURL(ctx, user.Picture)
		responses = append(responses, dto.UserResponse{
			Id:       user.Id,
			Picture:  user.Picture,
			FullName: user.FullName,
			Email:    user.Email,
			Phone:    user.Phone,
			Address:  user.Address,
			RoleId:   user.RoleId,
		})
	}

	ctx.JSON(http.StatusOK, dto.ResponseWrap{
		Success: true,
		Message: "list of users",
		Results: responses,
	})
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
	user.Picture = utils.BuildImageURL(ctx, user.Picture)
	result := dto.UserResponse{
		Id:       user.Id,
		Picture:  user.Picture,
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		Address:  user.Address,
		RoleId:   user.RoleId,
	}
	ctx.JSON(http.StatusOK, dto.ResponseWrap{
		Success: true,
		Message: "user data",
		Results: result,
	})
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

	user, err := h.service.UpdateUser(ctx.Request.Context(), id, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	user.Picture = utils.BuildImageURL(ctx, user.Picture)
	fmt.Println(user)
	ctx.JSON(http.StatusOK, dto.ResponseWrap{
		Success: true,
		Message: "User updated successfully",
		Results: user,
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
