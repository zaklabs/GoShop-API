// ============================================================================
// Project Name : GoShop API
// File         : alamat.go
// Description  : Model dan DTO untuk entitas Alamat
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi struct Alamat dan request/response DTOs
// - User dapat memiliki multiple alamat pengiriman
// - Alamat digunakan untuk pengiriman produk
//
// ============================================================================

package model

import "time"

// Alamat represents alamat table
type Alamat struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser       int        `gorm:"column:id_user;index" json:"id_user"`
	JudulAlamat  string     `gorm:"column:judul_alamat;type:varchar(255)" json:"judul_alamat"`
	NamaPenerima string     `gorm:"column:nama_penerima;type:varchar(255)" json:"nama_penerima"`
	NoTelp       string     `gorm:"column:no_telp;type:varchar(255)" json:"no_telp"`
	DetailAlamat string     `gorm:"column:detail_alamat;type:varchar(255)" json:"detail_alamat"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:date" json:"updated_at"`
	CreatedAt    *time.Time `gorm:"column:created_at;type:date" json:"created_at"`
	User         *User      `gorm:"foreignKey:IDUser;references:ID" json:"-"`
}

func (Alamat) TableName() string {
	return "alamat"
}

// AlamatRequest DTO
type AlamatRequest struct {
	JudulAlamat  string `json:"judul_alamat" binding:"required"`
	NamaPenerima string `json:"nama_penerima" binding:"required"`
	NoTelp       string `json:"no_telp" binding:"required"`
	DetailAlamat string `json:"detail_alamat" binding:"required"`
}
