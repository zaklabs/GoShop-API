// ============================================================================
// Project Name : GoShop API
// File         : category.go
// Description  : Model dan DTO untuk entitas Category
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi struct Category dan request/response DTOs
// - Category digunakan untuk mengklasifikasikan produk
// - Hanya admin yang dapat mengelola category
//
// ============================================================================

package model

import "time"

// Category represents category table
type Category struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaCategory string     `gorm:"column:nama_category;type:varchar(255)" json:"nama_category"`
	CreatedAt    *time.Time `gorm:"column:created_at;type:date" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:date" json:"updated_at"`
}

func (Category) TableName() string {
	return "category"
}

// CategoryRequest DTO
type CategoryRequest struct {
	NamaCategory string `json:"nama_category" binding:"required"`
}
