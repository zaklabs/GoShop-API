// ============================================================================
// Project Name : GoShop API
// File         : user.go
// Description  : Model dan DTO untuk entitas User
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi struct User dan request/response DTOs
// - User dapat memiliki role admin atau regular user
// - Password disimpan dalam bentuk hash
//
// ============================================================================

package model

import (
	"time"
)

// User represents users table
type User struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama         string     `gorm:"type:varchar(255)" json:"nama"`
	KataSandi    string     `gorm:"column:kata_sandi;type:varchar(255)" json:"-"`
	NoTelp       string     `gorm:"column:notelp;type:varchar(255)" json:"no_telp"`
	TanggalLahir *time.Time `gorm:"column:tanggal_lahir;type:date" json:"tanggal_lahir"`
	JenisKelamin string     `gorm:"column:jenis_kelamin;type:varchar(255)" json:"jenis_kelamin"`
	Tentang      string     `gorm:"type:text" json:"tentang"`
	Pekerjaan    string     `gorm:"type:varchar(255)" json:"pekerjaan"`
	Email        string     `gorm:"type:varchar(255)" json:"email"`
	IDProvinsi   *int       `gorm:"column:id_provinsi" json:"id_provinsi"`
	IDKota       *int       `gorm:"column:id_kota" json:"id_kota"`
	IsAdmin      bool       `gorm:"column:isAdmin;type:tinyint(1);default:0" json:"is_admin"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:date" json:"updated_at"`
	CreatedAt    *time.Time `gorm:"column:created_at;type:date" json:"created_at"`
}

func (User) TableName() string {
	return "users"
}

// RegisterRequest DTO
type RegisterRequest struct {
	Nama         string `json:"nama" binding:"required"`
	KataSandi    string `json:"kata_sandi" binding:"required,min=6"`
	NoTelp       string `json:"no_telp" binding:"required"`
	TanggalLahir string `json:"tanggal_Lahir"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email" binding:"required,email"`
	IDProvinsi   string `json:"id_provinsi"`
	IDKota       string `json:"id_kota"`
}

// LoginRequest DTO
type LoginRequest struct {
	NoTelp    string `json:"no_telp" binding:"required"`
	KataSandi string `json:"kata_sandi" binding:"required"`
}

// UpdateProfileRequest DTO
type UpdateProfileRequest struct {
	Nama         string `json:"nama"`
	KataSandi    string `json:"kata_sandi"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_Lahir"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	IDProvinsi   string `json:"id_provinsi"`
	IDKota       string `json:"id_kota"`
}

// UserResponse DTO
type UserResponse struct {
	Nama         string                 `json:"nama"`
	NoTelp       string                 `json:"no_telp"`
	TanggalLahir string                 `json:"tanggal_Lahir"`
	Tentang      string                 `json:"tentang"`
	Pekerjaan    string                 `json:"pekerjaan"`
	Email        string                 `json:"email"`
	IDProvinsi   map[string]interface{} `json:"id_provinsi"`
	IDKota       map[string]interface{} `json:"id_kota"`
}

// LoginResponse DTO
type LoginResponse struct {
	UserResponse
	Token string `json:"token"`
}
