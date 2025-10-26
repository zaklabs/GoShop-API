// ============================================================================
// Project Name : GoShop API
// File         : toko.go
// Description  : Model dan DTO untuk entitas Toko
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi struct Toko dan request/response DTOs
// - Setiap user hanya dapat memiliki satu toko
// - Toko dapat memiliki foto/logo
//
// ============================================================================

package model

import "time"

// Toko represents toko table
type Toko struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser    int        `gorm:"column:id_user;index" json:"user_id"`
	NamaToko  string     `gorm:"column:nama_toko;type:varchar(255)" json:"nama_toko"`
	URLFoto   string     `gorm:"column:url_toko;type:varchar(255)" json:"url_foto"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:date" json:"updated_at"`
	CreatedAt *time.Time `gorm:"column:created_at;type:date" json:"created_at"`
	User      *User      `gorm:"foreignKey:IDUser;references:ID" json:"-"`
}

func (Toko) TableName() string {
	return "toko"
}

// TokoResponse DTO
type TokoResponse struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	URLFoto  string `json:"url_foto"`
	UserID   int    `json:"user_id,omitempty"`
}

// UpdateTokoRequest DTO
type UpdateTokoRequest struct {
	NamaToko string `form:"nama_toko"`
}
