// ============================================================================
// Project Name : GoShop API
// File         : log_produk_repository.go
// Description  : Repository layer untuk operasi database LogProduk
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi interface dan implementasi untuk CRUD LogProduk
// - LogProduk menyimpan snapshot produk saat transaksi dibuat
// - Menggunakan GORM sebagai ORM
//
// ============================================================================

package repository

import (
	"evermos-api/internal/model"

	"gorm.io/gorm"
)

// LogProdukRepository interface
type LogProdukRepository interface {
	Create(log *model.LogProduk) error
	FindByID(id int) (*model.LogProduk, error)
	ExistsByProdukID(produkID int) (bool, error)
}

type logProdukRepository struct {
	db *gorm.DB
}

// NewLogProdukRepository creates new log produk repository
func NewLogProdukRepository(db *gorm.DB) LogProdukRepository {
	return &logProdukRepository{db: db}
}

func (r *logProdukRepository) Create(log *model.LogProduk) error {
	return r.db.Create(log).Error
}

func (r *logProdukRepository) FindByID(id int) (*model.LogProduk, error) {
	var log model.LogProduk
	err := r.db.First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *logProdukRepository) ExistsByProdukID(produkID int) (bool, error) {
	var count int64
	err := r.db.Model(&model.LogProduk{}).Where("id_produk = ?", produkID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
