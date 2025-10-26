// ============================================================================
// Project Name : GoShop API
// File         : trx.go
// Description  : Model dan DTO untuk entitas Transaksi
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi struct Trx dan DetailTrx
// - Trx menyimpan informasi transaksi utama
// - DetailTrx menyimpan detail produk dalam transaksi
//
// ============================================================================

package model

import "time"

// Trx represents trx table
type Trx struct {
	ID               int         `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser           int         `gorm:"column:id_user;index" json:"id_user"`
	AlamatPengiriman int         `gorm:"column:alamat_pengiriman;index" json:"alamat_pengiriman"`
	HargaTotal       int         `gorm:"column:harga_total" json:"harga_total"`
	KodeInvoice      string      `gorm:"column:kode_invoice;type:varchar(255)" json:"kode_invoice"`
	MethodBayar      string      `gorm:"column:method_bayar;type:varchar(255)" json:"method_bayar"`
	UpdatedAt        *time.Time  `gorm:"column:updated_at;type:date" json:"updated_at"`
	CreatedAt        *time.Time  `gorm:"column:created_at;type:date" json:"created_at"`
	User             *User       `gorm:"foreignKey:IDUser;references:ID" json:"-"`
	Alamat           *Alamat     `gorm:"foreignKey:AlamatPengiriman;references:ID" json:"-"`
	DetailTrx        []DetailTrx `gorm:"foreignKey:IDTrx;references:ID" json:"detail_trx,omitempty"`
}

func (Trx) TableName() string {
	return "trx"
}

// DetailTrx represents detail_trx table
type DetailTrx struct {
	ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	IDTrx       int        `gorm:"column:id_trx;index" json:"id_trx"`
	IDLogProduk int        `gorm:"column:id_log_produk;index" json:"id_log_produk"`
	IDToko      int        `gorm:"column:id_toko;index" json:"id_toko"`
	Kuantitas   int        `gorm:"type:int" json:"kuantitas"`
	HargaTotal  int        `gorm:"column:harga_total" json:"harga_total"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:date" json:"updated_at"`
	CreatedAt   *time.Time `gorm:"column:created_at;type:date" json:"created_at"`
	Trx         *Trx       `gorm:"foreignKey:IDTrx;references:ID" json:"-"`
	LogProduk   *LogProduk `gorm:"foreignKey:IDLogProduk;references:ID" json:"log_produk,omitempty"`
	Toko        *Toko      `gorm:"foreignKey:IDToko;references:ID" json:"toko,omitempty"`
}

func (DetailTrx) TableName() string {
	return "detail_trx"
}

// CreateTrxRequest DTO
type CreateTrxRequest struct {
	AlamatPengiriman int                `json:"alamat_kirim" binding:"required"`
	MethodBayar      string             `json:"method_bayar" binding:"required"`
	DetailTrx        []DetailTrxRequest `json:"detail_trx" binding:"required,min=1"`
}

// DetailTrxRequest DTO
type DetailTrxRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	Kuantitas int `json:"kuantitas" binding:"required,min=1"`
}
