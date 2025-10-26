// ============================================================================
// Project Name : GoShop API
// File         : produk.go
// Description  : Model dan DTO untuk entitas Produk
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi struct Produk, FotoProduk, dan LogProduk
// - Produk dapat memiliki multiple foto
// - LogProduk menyimpan snapshot produk saat transaksi
//
// ============================================================================

package model

import "time"

// Produk represents produk table
type Produk struct {
	ID            int          `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaProduk    string       `gorm:"column:nama_produk;type:varchar(255)" json:"nama_produk"`
	Slug          string       `gorm:"type:varchar(255)" json:"slug"`
	HargaReseller string       `gorm:"column:harga_reseller;type:varchar(255)" json:"harga_reseler"`
	HargaKonsumen string       `gorm:"column:harga_konsumen;type:varchar(255)" json:"harga_konsumen"`
	Stok          int          `gorm:"type:int" json:"stok"`
	Deskripsi     string       `gorm:"type:text" json:"deskripsi"`
	CreatedAt     *time.Time   `gorm:"column:created_at;type:date" json:"created_at"`
	UpdatedAt     *time.Time   `gorm:"column:updated_at;type:date" json:"updated_at"`
	DeletedAt     *time.Time   `gorm:"column:deleted_at;type:date;index" json:"-"`
	IDToko        int          `gorm:"column:id_toko;index" json:"-"`
	IDCategory    int          `gorm:"column:id_category;index" json:"-"`
	Toko          *Toko        `gorm:"foreignKey:IDToko;references:ID" json:"toko,omitempty"`
	Category      *Category    `gorm:"foreignKey:IDCategory;references:ID" json:"category,omitempty"`
	Photos        []FotoProduk `gorm:"foreignKey:IDProduk;references:ID" json:"photos,omitempty"`
}

func (Produk) TableName() string {
	return "produk"
}

// FotoProduk represents foto_produk table
type FotoProduk struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	IDProduk  int        `gorm:"column:id_produk;index" json:"product_id"`
	URL       string     `gorm:"type:varchar(255)" json:"url"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:date" json:"updated_at"`
	CreatedAt *time.Time `gorm:"column:created_at;type:date" json:"created_at"`
	Produk    *Produk    `gorm:"foreignKey:IDProduk;references:ID" json:"-"`
}

func (FotoProduk) TableName() string {
	return "foto_produk"
}

// LogProduk represents log_produk table (snapshot of product at transaction time)
type LogProduk struct {
	ID            int        `gorm:"primaryKey;autoIncrement" json:"id"`
	IDProduk      int        `gorm:"column:id_produk;index" json:"id_produk"`
	NamaProduk    string     `gorm:"column:nama_produk;type:varchar(255)" json:"nama_produk"`
	Slug          string     `gorm:"type:varchar(255)" json:"slug"`
	HargaReseller string     `gorm:"column:harga_reseller;type:varchar(255)" json:"harga_reseller"`
	HargaKonsumen string     `gorm:"column:harga_konsumen;type:varchar(255)" json:"harga_konsumen"`
	Deskripsi     string     `gorm:"type:text" json:"deskripsi"`
	CreatedAt     *time.Time `gorm:"column:created_at;type:date" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:date" json:"updated_at"`
	IDToko        int        `gorm:"column:id_toko;index" json:"id_toko"`
	IDCategory    int        `gorm:"column:id_category;index" json:"id_category"`
	Produk        *Produk    `gorm:"foreignKey:IDProduk;references:ID" json:"-"`
	Toko          *Toko      `gorm:"foreignKey:IDToko;references:ID" json:"-"`
	Category      *Category  `gorm:"foreignKey:IDCategory;references:ID" json:"-"`
}

func (LogProduk) TableName() string {
	return "log_produk"
}

// CreateProdukRequest DTO
type CreateProdukRequest struct {
	NamaProduk    string `form:"nama_produk" binding:"required"`
	HargaReseller string `form:"harga_reseller" binding:"required"`
	HargaKonsumen string `form:"harga_konsumen" binding:"required"`
	Stok          int    `form:"stok" binding:"required"`
	Deskripsi     string `form:"deskripsi" binding:"required"`
	CategoryID    int    `form:"category_id" binding:"required"`
}

// UpdateProdukRequest DTO
type UpdateProdukRequest struct {
	NamaProduk    string `form:"nama_produk"`
	HargaReseller string `form:"harga_reseller"`
	HargaKonsumen string `form:"harga_konsumen"`
	Stok          int    `form:"stok"`
	Deskripsi     string `form:"deskripsi"`
	CategoryID    int    `form:"category_id"`
}
