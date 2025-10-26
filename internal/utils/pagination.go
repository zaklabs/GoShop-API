// ============================================================================
// Project Name : GoShop API
// File         : pagination.go
// Description  : Utility untuk pagination
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi fungsi untuk extract pagination params dari query
// - Mendukung parameter page dan limit
// - Menghitung offset untuk query database
//
// ============================================================================

package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// PaginationParams holds pagination parameters
type PaginationParams struct {
	Page   int
	Limit  int
	Offset int
}

// GetPaginationParams extracts pagination params from query
func GetPaginationParams(c *gin.Context) PaginationParams {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	return PaginationParams{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}
