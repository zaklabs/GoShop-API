// ============================================================================
// Project Name : GoShop API
// File         : detail_trx_repository.go
// Description  : Repository layer untuk operasi database DetailTrx
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi interface dan implementasi untuk CRUD DetailTrx
// - DetailTrx menyimpan item-item dalam transaksi
// - Menggunakan GORM sebagai ORM
//
// ============================================================================

package repository

import (
	"evermos-api/internal/model"

	"gorm.io/gorm"
)

// DetailTrxRepository interface
type DetailTrxRepository interface {
	Create(detail *model.DetailTrx) error
	FindByTrxID(trxID int) ([]model.DetailTrx, error)
}

type detailTrxRepository struct {
	db *gorm.DB
}

// NewDetailTrxRepository creates new detail trx repository
func NewDetailTrxRepository(db *gorm.DB) DetailTrxRepository {
	return &detailTrxRepository{db: db}
}

func (r *detailTrxRepository) Create(detail *model.DetailTrx) error {
	return r.db.Create(detail).Error
}

func (r *detailTrxRepository) FindByTrxID(trxID int) ([]model.DetailTrx, error) {
	var details []model.DetailTrx
	err := r.db.Where("id_trx = ?", trxID).Find(&details).Error
	return details, err
}
