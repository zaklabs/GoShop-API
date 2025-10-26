// ============================================================================
// Project Name : GoShop API
// File         : user_repository.go
// Description  : Repository layer untuk operasi database User
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi interface dan implementasi untuk CRUD User
// - Menggunakan GORM sebagai ORM
// - Menyediakan fungsi pencarian by ID, email, dan nomor telepon
//
// ============================================================================

package repository

import (
	"evermos-api/internal/model"

	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
	Create(user *model.User) error
	FindByID(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByNoTelp(noTelp string) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id int) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByNoTelp(noTelp string) (*model.User, error) {
	var user model.User
	err := r.db.Where("notelp = ?", noTelp).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&model.User{}, id).Error
}
