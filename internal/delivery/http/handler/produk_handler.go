// ============================================================================
// Project Name : GoShop API
// File         : produk_handler.go
// Description  : Handler untuk manajemen produk
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi endpoint untuk CRUD produk
// - Mendukung upload multiple foto produk
// - Menyediakan fitur filter dan pencarian produk
//
// ============================================================================

package handler

import (
	"evermos-api/internal/delivery/middleware"
	"evermos-api/internal/model"
	"evermos-api/internal/usecase"
	"evermos-api/internal/utils"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProdukHandler handles produk endpoints
type ProdukHandler struct {
	produkUsecase usecase.ProdukUsecase
	uploadPath    string
}

// NewProdukHandler creates new produk handler
func NewProdukHandler(produkUsecase usecase.ProdukUsecase, uploadPath string) *ProdukHandler {
	return &ProdukHandler{
		produkUsecase: produkUsecase,
		uploadPath:    uploadPath,
	}
}

// GetAllProduk gets all produk with pagination and filters
func (h *ProdukHandler) GetAllProduk(c *gin.Context) {
	params := utils.GetPaginationParams(c)

	// Get filters
	filters := make(map[string]string)
	if namaProduk := c.Query("nama_produk"); namaProduk != "" {
		filters["nama_produk"] = namaProduk
	}
	if categoryID := c.Query("category_id"); categoryID != "" {
		filters["category_id"] = categoryID
	}
	if tokoID := c.Query("toko_id"); tokoID != "" {
		filters["toko_id"] = tokoID
	}
	if minHarga := c.Query("min_harga"); minHarga != "" {
		filters["min_harga"] = minHarga
	}
	if maxHarga := c.Query("max_harga"); maxHarga != "" {
		filters["max_harga"] = maxHarga
	}

	result, err := h.produkUsecase.GetAllProduk(params.Limit, params.Offset, filters)
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

// GetProdukByID gets produk by ID
func (h *ProdukHandler) GetProdukByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{"Invalid product ID"},
		))
		return
	}

	produk, err := h.produkUsecase.GetProdukByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		produk,
	))
}

// CreateProduk creates new produk
func (h *ProdukHandler) CreateProduk(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req model.CreateProdukRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to POST data",
			[]string{err.Error()},
		))
		return
	}

	// Handle multiple file uploads
	form, _ := c.MultipartForm()
	files := form.File["photos"]

	id, err := h.produkUsecase.CreateProduk(userID, req, files, h.uploadPath)
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

// UpdateProduk updates produk
func (h *ProdukHandler) UpdateProduk(c *gin.Context) {
	userID := middleware.GetUserID(c)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{"Invalid product ID"},
		))
		return
	}

	var req model.UpdateProdukRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{err.Error()},
		))
		return
	}

	// Handle multiple file uploads
	form, _ := c.MultipartForm()
	var files []*multipart.FileHeader
	if form != nil {
		files = form.File["photos"]
	}

	if err := h.produkUsecase.UpdateProduk(id, userID, req, files, h.uploadPath); err != nil {
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

// DeleteProduk deletes produk
func (h *ProdukHandler) DeleteProduk(c *gin.Context) {
	userID := middleware.GetUserID(c)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to DELETE data",
			[]string{"Invalid product ID"},
		))
		return
	}

	if err := h.produkUsecase.DeleteProduk(id, userID); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		"",
	))
}
