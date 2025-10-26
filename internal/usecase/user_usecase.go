// ============================================================================
// Project Name : GoShop API
// File         : user_usecase.go
// Description  : Business logic untuk manajemen user
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi logic untuk mendapatkan dan update profile user
// - Menangani validasi data user
// - Format response sesuai dengan API specification
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
)

// UserUsecase interface
type UserUsecase interface {
	GetProfile(userID int) (*model.UserResponse, error)
	UpdateProfile(userID int, req model.UpdateProfileRequest) error
}

type userUsecase struct {
	userRepo       repository.UserRepository
	wilayahUsecase WilayahUsecase
}

// NewUserUsecase creates new user usecase
func NewUserUsecase(userRepo repository.UserRepository, wilayahUsecase WilayahUsecase) UserUsecase {
	return &userUsecase{
		userRepo:       userRepo,
		wilayahUsecase: wilayahUsecase,
	}
}

func (u *userUsecase) GetProfile(userID int) (*model.UserResponse, error) {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	tanggalLahir := ""
	if user.TanggalLahir != nil {
		tanggalLahir = user.TanggalLahir.Format("02/01/2006")
	}

	provinsi := make(map[string]interface{})
	kota := make(map[string]interface{})

	if user.IDProvinsi != nil {
		provinsiID := strconv.Itoa(*user.IDProvinsi)
		provinsi["id"] = provinsiID

		// Get province name from API
		if provinsiData, err := u.wilayahUsecase.GetDetailProvince(provinsiID); err == nil {
			provinsi["name"] = provinsiData.Name
		} else {
			provinsi["name"] = ""
		}
	}

	if user.IDKota != nil {
		kotaID := strconv.Itoa(*user.IDKota)
		kota["id"] = kotaID
		if user.IDProvinsi != nil {
			kota["province_id"] = strconv.Itoa(*user.IDProvinsi)
		}

		// Get city name from API
		if kotaData, err := u.wilayahUsecase.GetDetailCity(kotaID); err == nil {
			kota["name"] = kotaData.Name
		} else {
			kota["name"] = ""
		}
	}

	response := &model.UserResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: tanggalLahir,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IDProvinsi:   provinsi,
		IDKota:       kota,
	}

	return response, nil
}

func (u *userUsecase) UpdateProfile(userID int, req model.UpdateProfileRequest) error {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Update fields if provided
	if req.Nama != "" {
		user.Nama = req.Nama
	}
	if req.KataSandi != "" {
		hashedPassword, err := utils.HashPassword(req.KataSandi)
		if err != nil {
			return err
		}
		user.KataSandi = hashedPassword
	}
	if req.NoTelp != "" {
		user.NoTelp = req.NoTelp
	}
	if req.TanggalLahir != "" {
		t, err := time.Parse("02/01/2006", req.TanggalLahir)
		if err == nil {
			user.TanggalLahir = &t
		}
	}
	if req.Pekerjaan != "" {
		user.Pekerjaan = req.Pekerjaan
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.IDProvinsi != "" {
		if id, err := strconv.Atoi(req.IDProvinsi); err == nil {
			user.IDProvinsi = &id
		}
	}
	if req.IDKota != "" {
		if id, err := strconv.Atoi(req.IDKota); err == nil {
			user.IDKota = &id
		}
	}

	now := time.Now()
	user.UpdatedAt = &now

	return u.userRepo.Update(user)
}
