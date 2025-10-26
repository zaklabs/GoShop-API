// ============================================================================
// Project Name : GoShop API
// File         : auth_handler.go
// Description  : Handler untuk login dan register pengguna
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi endpoint login dan register pengguna
// - Token JWT dihasilkan setelah login berhasil
//
// ============================================================================

package handler

import (
	"evermos-api/internal/delivery/middleware"
	"evermos-api/internal/model"
	"evermos-api/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles auth endpoints
type AuthHandler struct {
	authUsecase usecase.AuthUsecase
	jwtSecret   string
	expireHours int
}

// NewAuthHandler creates new auth handler
func NewAuthHandler(authUsecase usecase.AuthUsecase, jwtSecret string, expireHours int) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
		jwtSecret:   jwtSecret,
		expireHours: expireHours,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to POST data",
			[]string{err.Error()},
		))
		return
	}

	if err := h.authUsecase.Register(req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to POST data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to POST data",
		"Register Succeed",
	))
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to POST data",
			[]string{err.Error()},
		))
		return
	}

	response, err := h.authUsecase.Login(req, h.jwtSecret, h.expireHours)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(
			"Failed to POST data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to POST data",
		response,
	))
}

// UserHandler handles user endpoints
type UserHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler creates new user handler
func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

// GetProfile gets user profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)

	profile, err := h.userUsecase.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		profile,
	))
}

// UpdateProfile updates user profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req model.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{err.Error()},
		))
		return
	}

	if err := h.userUsecase.UpdateProfile(userID, req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		"",
	))
}
