// ============================================================================
// Project Name : GoShop API
// File         : produk_repository.go
// Description  : Repository layer untuk operasi database Produk
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi interface dan implementasi untuk CRUD Produk
// - Menggunakan GORM sebagai ORM
// - Mendukung preload relasi (Toko, Category, Photos)
//
// ============================================================================

package repository

import (
	"evermos-api/internal/model"

	"gorm.io/gorm"
)

// ProdukRepository interface
type ProdukRepository interface {
	Create(produk *model.Produk) error
	FindByID(id int) (*model.Produk, error)
	FindByIDWithRelations(id int) (*model.Produk, error)
	FindAll(limit, offset int, filters map[string]interface{}) ([]model.Produk, error)
	Update(produk *model.Produk) error
	Delete(id int) error
}

type produkRepository struct {
	db *gorm.DB
}

// NewProdukRepository creates new produk repository
func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &produkRepository{db: db}
}

func (r *produkRepository) Create(produk *model.Produk) error {
	return r.db.Create(produk).Error
}

func (r *produkRepository) FindByID(id int) (*model.Produk, error) {
	var produk model.Produk
	// err := r.db.First(&produk, id).Error
	err := r.db.Where("deleted_at IS NULL").First(&produk, id).Error
	if err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *produkRepository) FindByIDWithRelations(id int) (*model.Produk, error) {
	var produk model.Produk
	// err := r.db.Preload("Toko").Preload("Category").Preload("Photos").First(&produk, id).Error
	err := r.db.Where("deleted_at IS NULL").Preload("Toko").Preload("Category").Preload("Photos").First(&produk, id).Error
	if err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *produkRepository) FindAll(limit, offset int, filters map[string]interface{}) ([]model.Produk, error) {
	var produks []model.Produk
	// query := r.db.Preload("Toko").Preload("Category").Preload("Photos").Limit(limit).Offset(offset)
	query := r.db.Where("deleted_at IS NULL").Preload("Toko").Preload("Category").Preload("Photos").Limit(limit).Offset(offset)
	// Apply filters
	if namaProduk, ok := filters["nama_produk"].(string); ok && namaProduk != "" {
		query = query.Where("nama_produk LIKE ?", "%"+namaProduk+"%")
	}

	if categoryID, ok := filters["category_id"].(int); ok && categoryID > 0 {
		query = query.Where("id_category = ?", categoryID)
	}

	if tokoID, ok := filters["toko_id"].(int); ok && tokoID > 0 {
		query = query.Where("id_toko = ?", tokoID)
	}

	if minHarga, ok := filters["min_harga"].(int); ok && minHarga > 0 {
		query = query.Where("CAST(harga_konsumen AS UNSIGNED) >= ?", minHarga)
	}

	if maxHarga, ok := filters["max_harga"].(int); ok && maxHarga > 0 {
		query = query.Where("CAST(harga_konsumen AS UNSIGNED) <= ?", maxHarga)
	}

	err := query.Find(&produks).Error
	return produks, err
}

func (r *produkRepository) Update(produk *model.Produk) error {
	return r.db.Save(produk).Error
}

func (r *produkRepository) Delete(id int) error {
	return r.db.Delete(&model.Produk{}, id).Error
}
