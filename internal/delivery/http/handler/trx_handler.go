// ============================================================================
// Project Name : GoShop API
// File         : trx_handler.go
// Description  : Handler untuk manajemen transaksi
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi endpoint untuk CRUD transaksi
// - Mendukung pembuatan transaksi dengan multiple produk
// - Otomatis generate nomor invoice
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

// TrxHandler handles transaction endpoints
type TrxHandler struct {
	trxUsecase usecase.TrxUsecase
}

// NewTrxHandler creates new trx handler
func NewTrxHandler(trxUsecase usecase.TrxUsecase) *TrxHandler {
	return &TrxHandler{trxUsecase: trxUsecase}
}

// GetAllTrx gets all user's transactions
func (h *TrxHandler) GetAllTrx(c *gin.Context) {
	userID := middleware.GetUserID(c)
	params := utils.GetPaginationParams(c)

	trxs, err := h.trxUsecase.GetAllTrx(userID, params.Limit, params.Offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		trxs,
	))
}

// GetTrxByID gets transaction by ID
func (h *TrxHandler) GetTrxByID(c *gin.Context) {
	userID := middleware.GetUserID(c)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{"Invalid transaction ID"},
		))
		return
	}

	trx, err := h.trxUsecase.GetTrxByID(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		trx,
	))
}

// CreateTrx creates new transaction
func (h *TrxHandler) CreateTrx(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req model.CreateTrxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to POST data",
			[]string{err.Error()},
		))
		return
	}

	id, err := h.trxUsecase.CreateTrx(userID, req)
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
