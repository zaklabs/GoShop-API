// ============================================================================
// Project Name : GoShop API
// File         : trx_repository.go
// Description  : Repository layer untuk operasi database Transaksi
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi interface dan implementasi untuk CRUD Transaksi
// - Menggunakan GORM sebagai ORM
// - Mendukung preload detail transaksi dengan relasi
//
// ============================================================================

package repository

import (
	"evermos-api/internal/model"

	"gorm.io/gorm"
)

// TrxRepository interface
type TrxRepository interface {
	Create(trx *model.Trx) error
	FindByID(id int) (*model.Trx, error)
	FindByIDWithDetails(id int) (*model.Trx, error)
	FindByUserID(userID int, limit, offset int) ([]model.Trx, error)
	Update(trx *model.Trx) error
}

type trxRepository struct {
	db *gorm.DB
}

// NewTrxRepository creates new trx repository
func NewTrxRepository(db *gorm.DB) TrxRepository {
	return &trxRepository{db: db}
}

func (r *trxRepository) Create(trx *model.Trx) error {
	return r.db.Create(trx).Error
}

func (r *trxRepository) FindByID(id int) (*model.Trx, error) {
	var trx model.Trx
	err := r.db.First(&trx, id).Error
	if err != nil {
		return nil, err
	}
	return &trx, nil
}

func (r *trxRepository) FindByIDWithDetails(id int) (*model.Trx, error) {
	var trx model.Trx
	err := r.db.Preload("DetailTrx.LogProduk").Preload("DetailTrx.Toko").First(&trx, id).Error
	if err != nil {
		return nil, err
	}
	return &trx, nil
}

func (r *trxRepository) FindByUserID(userID int, limit, offset int) ([]model.Trx, error) {
	var trxs []model.Trx
	err := r.db.Where("id_user = ?", userID).
		Preload("DetailTrx.LogProduk").
		Preload("DetailTrx.Toko").
		Limit(limit).Offset(offset).
		Find(&trxs).Error
	return trxs, err
}

func (r *trxRepository) Update(trx *model.Trx) error {
	return r.db.Save(trx).Error
}
