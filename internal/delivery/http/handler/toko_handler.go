// ============================================================================
// Project Name : GoShop API
// File         : toko_handler.go
// Description  : Handler untuk manajemen toko
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi endpoint untuk CRUD toko
// - Setiap user hanya dapat memiliki satu toko
// - Mendukung upload foto toko
//
// ============================================================================

package handler

import (
	"evermos-api/internal/delivery/middleware"
	"evermos-api/internal/model"
	"evermos-api/internal/usecase"
	"evermos-api/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TokoHandler handles toko endpoints
type TokoHandler struct {
	tokoUsecase usecase.TokoUsecase
	uploadPath  string
}

// NewTokoHandler creates new toko handler
func NewTokoHandler(tokoUsecase usecase.TokoUsecase, uploadPath string) *TokoHandler {
	return &TokoHandler{
		tokoUsecase: tokoUsecase,
		uploadPath:  uploadPath,
	}
}

// GetMyToko gets current user's toko
func (h *TokoHandler) GetMyToko(c *gin.Context) {
	userID := middleware.GetUserID(c)

	toko, err := h.tokoUsecase.GetMyToko(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		toko,
	))
}

// GetTokoByID gets toko by ID
func (h *TokoHandler) GetTokoByID(c *gin.Context) {
	idParam := c.Param("id_toko")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{"Invalid toko ID"},
		))
		return
	}

	toko, err := h.tokoUsecase.GetTokoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		toko,
	))
}

// GetAllToko gets all toko with pagination
func (h *TokoHandler) GetAllToko(c *gin.Context) {
	params := utils.GetPaginationParams(c)
	nama := c.Query("nama")

	result, err := h.tokoUsecase.GetAllToko(params.Limit, params.Offset, nama)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		result,
	))
}

// UpdateToko updates toko
func (h *TokoHandler) UpdateToko(c *gin.Context) {
	userID := middleware.GetUserID(c)

	idParam := c.Param("id_toko")
	tokoID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{"Invalid toko ID"},
		))
		return
	}

	var req model.UpdateTokoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{err.Error()},
		))
		return
	}

	// Handle file upload
	file, _ := c.FormFile("photo")

	if err := h.tokoUsecase.UpdateToko(tokoID, userID, req.NamaToko, file, h.uploadPath); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to UPDATE data",
		"Update toko succeed",
	))
}
