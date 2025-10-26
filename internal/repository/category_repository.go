// ============================================================================
// Project Name : GoShop API
// File         : category_repository.go
// Description  : Repository layer untuk operasi database Category
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi interface dan implementasi untuk CRUD Category
// - Menggunakan GORM sebagai ORM
// - Menyediakan fungsi untuk mendapatkan semua kategori
//
// ============================================================================

package repository

import (
	"evermos-api/internal/model"

	"gorm.io/gorm"
)

// CategoryRepository interface
type CategoryRepository interface {
	Create(category *model.Category) error
	FindByID(id int) (*model.Category, error)
	FindAll() ([]model.Category, error)
	Update(category *model.Category) error
	Delete(id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates new category repository
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FindByID(id int) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) FindAll() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id int) error {
	return r.db.Delete(&model.Category{}, id).Error
}
