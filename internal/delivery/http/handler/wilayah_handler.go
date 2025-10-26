// ============================================================================
// Project Name : GoShop API
// File         : wilayah_handler.go
// Description  : Handler untuk endpoint wilayah (provinsi dan kota)
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi endpoint untuk mendapatkan data provinsi dan kota
// - Data diambil dari API eksternal emsifa
// - Endpoint public tanpa autentikasi
//
// ============================================================================

package handler

import (
	"evermos-api/internal/model"
	"evermos-api/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WilayahHandler handles wilayah endpoints
type WilayahHandler struct {
	wilayahUsecase usecase.WilayahUsecase
}

// NewWilayahHandler creates new wilayah handler
func NewWilayahHandler(wilayahUsecase usecase.WilayahUsecase) *WilayahHandler {
	return &WilayahHandler{wilayahUsecase: wilayahUsecase}
}

// GetListProvince gets all provinces
func (h *WilayahHandler) GetListProvince(c *gin.Context) {
	provinces, err := h.wilayahUsecase.GetListProvince()
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		provinces,
	))
}

// GetDetailProvince gets province by ID
func (h *WilayahHandler) GetDetailProvince(c *gin.Context) {
	id := c.Param("id")

	province, err := h.wilayahUsecase.GetDetailProvince(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		province,
	))
}

// GetListCity gets all cities by province ID
func (h *WilayahHandler) GetListCity(c *gin.Context) {
	provinceID := c.Param("province_id")

	cities, err := h.wilayahUsecase.GetListCity(provinceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		cities,
	))
}

// GetDetailCity gets city by ID
func (h *WilayahHandler) GetDetailCity(c *gin.Context) {
	id := c.Param("id")

	city, err := h.wilayahUsecase.GetDetailCity(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		city,
	))
}
