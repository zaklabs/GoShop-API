// ============================================================================
// Project Name : GoShop API
// File         : alamat_repository.go
// Description  : Repository layer untuk operasi database Alamat
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi interface dan implementasi untuk CRUD Alamat
// - Menggunakan GORM sebagai ORM
// - Menyediakan fungsi untuk mendapatkan alamat user
//
// ============================================================================

package repository

import (
	"evermos-api/internal/model"

	"gorm.io/gorm"
)

// AlamatRepository interface
type AlamatRepository interface {
	Create(alamat *model.Alamat) error
	FindByID(id int) (*model.Alamat, error)
	FindByUserID(userID int) ([]model.Alamat, error)
	Update(alamat *model.Alamat) error
	Delete(id int) error
}

type alamatRepository struct {
	db *gorm.DB
}

// NewAlamatRepository creates new alamat repository
func NewAlamatRepository(db *gorm.DB) AlamatRepository {
	return &alamatRepository{db: db}
}

func (r *alamatRepository) Create(alamat *model.Alamat) error {
	return r.db.Create(alamat).Error
}

func (r *alamatRepository) FindByID(id int) (*model.Alamat, error) {
	var alamat model.Alamat
	err := r.db.First(&alamat, id).Error
	if err != nil {
		return nil, err
	}
	return &alamat, nil
}

func (r *alamatRepository) FindByUserID(userID int) ([]model.Alamat, error) {
	var alamats []model.Alamat
	err := r.db.Where("id_user = ?", userID).Find(&alamats).Error
	return alamats, err
}

func (r *alamatRepository) Update(alamat *model.Alamat) error {
	return r.db.Save(alamat).Error
}

func (r *alamatRepository) Delete(id int) error {
	return r.db.Delete(&model.Alamat{}, id).Error
}
