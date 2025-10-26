// ============================================================================
// Project Name : GoShop API
// File         : auth_usecase.go
// Description  : Business logic untuk autentikasi user
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi logic untuk register dan login
// - Menggunakan JWT untuk autentikasi
// - Otomatis membuat toko setelah user register
//
// ============================================================================

package usecase

import (
	"errors"
	"evermos-api/internal/model"
	"evermos-api/internal/repository"
	"evermos-api/internal/utils"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// AuthUsecase interface
type AuthUsecase interface {
	Register(req model.RegisterRequest) error
	Login(req model.LoginRequest, jwtSecret string, expireHours int) (*model.LoginResponse, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
	tokoRepo repository.TokoRepository
	db       *gorm.DB
}

// NewAuthUsecase creates new auth usecase
func NewAuthUsecase(userRepo repository.UserRepository, tokoRepo repository.TokoRepository, db *gorm.DB) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		tokoRepo: tokoRepo,
		db:       db,
	}
}

func (u *authUsecase) Register(req model.RegisterRequest) error {
	// Check if email or notelp already exists
	if existingUser, _ := u.userRepo.FindByEmail(req.Email); existingUser != nil {
		return errors.New("email already registered")
	}

	if existingUser, _ := u.userRepo.FindByNoTelp(req.NoTelp); existingUser != nil {
		return errors.New("nomor telepon already registered")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.KataSandi)
	if err != nil {
		return err
	}

	// Parse dates and IDs
	var tanggalLahir *time.Time
	if req.TanggalLahir != "" {
		t, err := time.Parse("02/01/2006", req.TanggalLahir)
		if err == nil {
			tanggalLahir = &t
		}
	}

	var idProvinsi, idKota *int
	if req.IDProvinsi != "" {
		if id, err := strconv.Atoi(req.IDProvinsi); err == nil {
			idProvinsi = &id
		}
	}
	if req.IDKota != "" {
		if id, err := strconv.Atoi(req.IDKota); err == nil {
			idKota = &id
		}
	}

	now := time.Now()
	user := &model.User{
		Nama:         req.Nama,
		KataSandi:    hashedPassword,
		NoTelp:       req.NoTelp,
		TanggalLahir: tanggalLahir,
		Pekerjaan:    req.Pekerjaan,
		Email:        req.Email,
		IDProvinsi:   idProvinsi,
		IDKota:       idKota,
		IsAdmin:      false,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}

	// Use transaction to create user and toko
	return u.db.Transaction(func(tx *gorm.DB) error {
		// Create user
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// Auto-create toko for new user
		tokoName := utils.GenerateSlug(user.Nama)
		toko := &model.Toko{
			IDUser:    user.ID,
			NamaToko:  tokoName,
			CreatedAt: &now,
			UpdatedAt: &now,
		}

		if err := tx.Create(toko).Error; err != nil {
			return err
		}

		return nil
	})
}

func (u *authUsecase) Login(req model.LoginRequest, jwtSecret string, expireHours int) (*model.LoginResponse, error) {
	// Find user by notelp
	user, err := u.userRepo.FindByNoTelp(req.NoTelp)
	if err != nil {
		return nil, errors.New("No Telp atau kata sandi salah")
	}

	// Check password
	if !utils.CheckPasswordHash(req.KataSandi, user.KataSandi) {
		return nil, errors.New("No Telp atau kata sandi salah")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, user.IsAdmin, jwtSecret, expireHours)
	if err != nil {
		return nil, err
	}

	// Format response
	tanggalLahir := ""
	if user.TanggalLahir != nil {
		tanggalLahir = user.TanggalLahir.Format("02/01/2006")
	}

	provinsi := make(map[string]interface{})
	kota := make(map[string]interface{})

	if user.IDProvinsi != nil {
		provinsi["id"] = strconv.Itoa(*user.IDProvinsi)
		provinsi["name"] = "ACEH" // This should be fetched from province table in real implementation
	}

	if user.IDKota != nil {
		kota["id"] = strconv.Itoa(*user.IDKota)
		kota["province_id"] = strconv.Itoa(*user.IDProvinsi)
		kota["name"] = "KABUPATEN SIMEULUE" // This should be fetched from city table
	}

	response := &model.LoginResponse{
		UserResponse: model.UserResponse{
			Nama:         user.Nama,
			NoTelp:       user.NoTelp,
			TanggalLahir: tanggalLahir,
			Tentang:      user.Tentang,
			Pekerjaan:    user.Pekerjaan,
			Email:        user.Email,
			IDProvinsi:   provinsi,
			IDKota:       kota,
		},
		Token: token,
	}

	return response, nil
}
