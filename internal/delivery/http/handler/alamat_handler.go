// ============================================================================
// Project Name : GoShop API
// File         : alamat_handler.go
// Description  : Handler untuk manajemen alamat pengiriman
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi endpoint untuk CRUD alamat pengiriman
// - Setiap user dapat memiliki multiple alamat
// - Satu alamat dapat ditandai sebagai alamat utama
//
// ============================================================================

package handler

import (
	"evermos-api/internal/delivery/middleware"
	"evermos-api/internal/model"
	"evermos-api/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AlamatHandler handles alamat endpoints
type AlamatHandler struct {
	alamatUsecase usecase.AlamatUsecase
}

// NewAlamatHandler creates new alamat handler
func NewAlamatHandler(alamatUsecase usecase.AlamatUsecase) *AlamatHandler {
	return &AlamatHandler{alamatUsecase: alamatUsecase}
}

// GetMyAlamat gets current user's alamat
func (h *AlamatHandler) GetMyAlamat(c *gin.Context) {
	userID := middleware.GetUserID(c)

	alamats, err := h.alamatUsecase.GetMyAlamat(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		alamats,
	))
}

// GetAlamatByID gets alamat by ID
func (h *AlamatHandler) GetAlamatByID(c *gin.Context) {
	userID := middleware.GetUserID(c)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{"Invalid alamat ID"},
		))
		return
	}

	alamat, err := h.alamatUsecase.GetAlamatByID(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		alamat,
	))
}

// CreateAlamat creates new alamat
func (h *AlamatHandler) CreateAlamat(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req model.AlamatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to POST data",
			[]string{err.Error()},
		))
		return
	}

	id, err := h.alamatUsecase.CreateAlamat(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to POST data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to POST data",
		id,
	))
}

// UpdateAlamat updates alamat
func (h *AlamatHandler) UpdateAlamat(c *gin.Context) {
	userID := middleware.GetUserID(c)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{"Invalid alamat ID"},
		))
		return
	}

	var req model.AlamatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{err.Error()},
		))
		return
	}

	if err := h.alamatUsecase.UpdateAlamat(id, userID, req); err != nil {
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

// DeleteAlamat deletes alamat
func (h *AlamatHandler) DeleteAlamat(c *gin.Context) {
	userID := middleware.GetUserID(c)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to DELETE data",
			[]string{"Invalid alamat ID"},
		))
		return
	}

	if err := h.alamatUsecase.DeleteAlamat(id, userID); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to DELETE data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		"",
	))
}
