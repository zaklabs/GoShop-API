// ============================================================================
// Project Name : GoShop API
// File         : category_handler.go
// Description  : Handler untuk manajemen kategori produk
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi endpoint untuk CRUD kategori produk
// - Kategori digunakan untuk mengklasifikasikan produk
// - Hanya admin yang dapat mengelola kategori
//
// ============================================================================

package handler

import (
	"evermos-api/internal/model"
	"evermos-api/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CategoryHandler handles category endpoints
type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

// NewCategoryHandler creates new category handler
func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: categoryUsecase}
}

// GetAllCategory gets all categories
func (h *CategoryHandler) GetAllCategory(c *gin.Context) {
	categories, err := h.categoryUsecase.GetAllCategory()
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		categories,
	))
}

// GetCategoryByID gets category by ID
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to GET data",
			[]string{"Invalid category ID"},
		))
		return
	}

	category, err := h.categoryUsecase.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(
			"Failed to GET data",
			[]string{err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(
		"Succeed to GET data",
		category,
	))
}

// CreateCategory creates new category (admin only)
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req model.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to POST data",
			[]string{err.Error()},
		))
		return
	}

	id, err := h.categoryUsecase.CreateCategory(req)
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

// UpdateCategory updates category (admin only)
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{"Invalid category ID"},
		))
		return
	}

	var req model.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to UPDATE data",
			[]string{err.Error()},
		))
		return
	}

	if err := h.categoryUsecase.UpdateCategory(id, req); err != nil {
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

// DeleteCategory deletes category (admin only)
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(
			"Failed to DELETE data",
			[]string{"Invalid category ID"},
		))
		return
	}

	if err := h.categoryUsecase.DeleteCategory(id); err != nil {
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
