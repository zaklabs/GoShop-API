// ============================================================================
// Project Name : GoShop API
// File         : category_usecase.go
// Description  : Business logic untuk manajemen kategori
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi logic untuk CRUD kategori produk
// - Kategori digunakan untuk mengklasifikasikan produk
// - Hanya admin yang dapat mengelola kategori
//
// ============================================================================

package usecase

import (
	"errors"
	"evermos-api/internal/model"
	"evermos-api/internal/repository"
	"time"
)

// CategoryUsecase interface
type CategoryUsecase interface {
	GetAllCategory() ([]model.Category, error)
	GetCategoryByID(id int) (*model.Category, error)
	CreateCategory(req model.CategoryRequest) (int, error)
	UpdateCategory(id int, req model.CategoryRequest) error
	DeleteCategory(id int) error
}

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

// NewCategoryUsecase creates new category usecase
func NewCategoryUsecase(categoryRepo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepo: categoryRepo}
}

func (u *categoryUsecase) GetAllCategory() ([]model.Category, error) {
	return u.categoryRepo.FindAll()
}

func (u *categoryUsecase) GetCategoryByID(id int) (*model.Category, error) {
	category, err := u.categoryRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("No Data Category")
	}
	return category, nil
}

func (u *categoryUsecase) CreateCategory(req model.CategoryRequest) (int, error) {
	now := time.Now()
	category := &model.Category{
		NamaCategory: req.NamaCategory,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}

	if err := u.categoryRepo.Create(category); err != nil {
		return 0, err
	}

	return category.ID, nil
}

func (u *categoryUsecase) UpdateCategory(id int, req model.CategoryRequest) error {
	category, err := u.categoryRepo.FindByID(id)
	if err != nil {
		return errors.New("category not found")
	}

	category.NamaCategory = req.NamaCategory
	now := time.Now()
	category.UpdatedAt = &now

	return u.categoryRepo.Update(category)
}

func (u *categoryUsecase) DeleteCategory(id int) error {
	_, err := u.categoryRepo.FindByID(id)
	if err != nil {
		return errors.New("record not found")
	}

	return u.categoryRepo.Delete(id)
}
