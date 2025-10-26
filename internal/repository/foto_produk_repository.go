// ============================================================================
// Project Name : GoShop API
// File         : foto_produk_repository.go
// Description  : Repository layer untuk operasi database FotoProduk
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi interface dan implementasi untuk CRUD FotoProduk
// - Menggunakan GORM sebagai ORM
// - Menyediakan fungsi untuk mengelola foto produk
//
// ============================================================================

package repository

import (
	"evermos-api/internal/model"

	"gorm.io/gorm"
)

// FotoProdukRepository interface
type FotoProdukRepository interface {
	Create(foto *model.FotoProduk) error
	FindByProdukID(produkID int) ([]model.FotoProduk, error)
	DeleteByProdukID(produkID int) error
}

type fotoProdukRepository struct {
	db *gorm.DB
}

// NewFotoProdukRepository creates new foto produk repository
func NewFotoProdukRepository(db *gorm.DB) FotoProdukRepository {
	return &fotoProdukRepository{db: db}
}

func (r *fotoProdukRepository) Create(foto *model.FotoProduk) error {
	return r.db.Create(foto).Error
}

func (r *fotoProdukRepository) FindByProdukID(produkID int) ([]model.FotoProduk, error) {
	var fotos []model.FotoProduk
	err := r.db.Where("id_produk = ?", produkID).Find(&fotos).Error
	return fotos, err
}

func (r *fotoProdukRepository) DeleteByProdukID(produkID int) error {
	return r.db.Where("id_produk = ?", produkID).Delete(&model.FotoProduk{}).Error
}
