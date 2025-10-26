// ============================================================================
// Project Name : GoShop API
// File         : toko_repository.go
// Description  : Repository layer untuk operasi database Toko
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi interface dan implementasi untuk CRUD Toko
// - Menggunakan GORM sebagai ORM
// - Menyediakan fungsi pencarian dan filter toko
//
// ============================================================================

package repository

import (
	"evermos-api/internal/model"

	"gorm.io/gorm"
)

// TokoRepository interface
type TokoRepository interface {
	Create(toko *model.Toko) error
	FindByID(id int) (*model.Toko, error)
	FindByUserID(userID int) (*model.Toko, error)
	FindAll(limit, offset int, nama string) ([]model.Toko, error)
	Update(toko *model.Toko) error
	Delete(id int) error
}

type tokoRepository struct {
	db *gorm.DB
}

// NewTokoRepository creates new toko repository
func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &tokoRepository{db: db}
}

func (r *tokoRepository) Create(toko *model.Toko) error {
	return r.db.Create(toko).Error
}

func (r *tokoRepository) FindByID(id int) (*model.Toko, error) {
	var toko model.Toko
	err := r.db.First(&toko, id).Error
	if err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *tokoRepository) FindByUserID(userID int) (*model.Toko, error) {
	var toko model.Toko
	err := r.db.Where("id_user = ?", userID).First(&toko).Error
	if err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *tokoRepository) FindAll(limit, offset int, nama string) ([]model.Toko, error) {
	var tokos []model.Toko
	query := r.db.Limit(limit).Offset(offset)

	if nama != "" {
		query = query.Where("nama_toko LIKE ?", "%"+nama+"%")
	}

	err := query.Find(&tokos).Error
	return tokos, err
}

func (r *tokoRepository) Update(toko *model.Toko) error {
	return r.db.Save(toko).Error
}

func (r *tokoRepository) Delete(id int) error {
	return r.db.Delete(&model.Toko{}, id).Error
}
