// ============================================================================
// Project Name : GoShop API
// File         : alamat_usecase.go
// Description  : Business logic untuk manajemen alamat
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi logic untuk CRUD alamat pengiriman
// - Validasi kepemilikan alamat oleh user
// - Setiap alamat terikat dengan user tertentu
//
// ============================================================================

package usecase

import (
	"errors"
	"evermos-api/internal/model"
	"evermos-api/internal/repository"
	"time"
)

// AlamatUsecase interface
type AlamatUsecase interface {
	GetMyAlamat(userID int) ([]model.Alamat, error)
	GetAlamatByID(id, userID int) (*model.Alamat, error)
	CreateAlamat(userID int, req model.AlamatRequest) (int, error)
	UpdateAlamat(id, userID int, req model.AlamatRequest) error
	DeleteAlamat(id, userID int) error
}

type alamatUsecase struct {
	alamatRepo repository.AlamatRepository
}

// NewAlamatUsecase creates new alamat usecase
func NewAlamatUsecase(alamatRepo repository.AlamatRepository) AlamatUsecase {
	return &alamatUsecase{alamatRepo: alamatRepo}
}

func (u *alamatUsecase) GetMyAlamat(userID int) ([]model.Alamat, error) {
	return u.alamatRepo.FindByUserID(userID)
}

func (u *alamatUsecase) GetAlamatByID(id, userID int) (*model.Alamat, error) {
	alamat, err := u.alamatRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("alamat not found")
	}

	// Check ownership
	if alamat.IDUser != userID {
		return nil, errors.New("unauthorized: not your alamat")
	}

	return alamat, nil
}

func (u *alamatUsecase) CreateAlamat(userID int, req model.AlamatRequest) (int, error) {
	now := time.Now()
	alamat := &model.Alamat{
		IDUser:       userID,
		JudulAlamat:  req.JudulAlamat,
		NamaPenerima: req.NamaPenerima,
		NoTelp:       req.NoTelp,
		DetailAlamat: req.DetailAlamat,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}

	if err := u.alamatRepo.Create(alamat); err != nil {
		return 0, err
	}

	return alamat.ID, nil
}

func (u *alamatUsecase) UpdateAlamat(id, userID int, req model.AlamatRequest) error {
	alamat, err := u.alamatRepo.FindByID(id)
	if err != nil {
		return errors.New("alamat not found")
	}

	// Check ownership
	if alamat.IDUser != userID {
		return errors.New("unauthorized: not your alamat")
	}

	// Update fields
	alamat.JudulAlamat = req.JudulAlamat
	alamat.NamaPenerima = req.NamaPenerima
	alamat.NoTelp = req.NoTelp
	alamat.DetailAlamat = req.DetailAlamat

	now := time.Now()
	alamat.UpdatedAt = &now

	return u.alamatRepo.Update(alamat)
}

func (u *alamatUsecase) DeleteAlamat(id, userID int) error {
	alamat, err := u.alamatRepo.FindByID(id)
	if err != nil {
		return errors.New("alamat not found")
	}

	// Check ownership
	if alamat.IDUser != userID {
		return errors.New("unauthorized: not your alamat")
	}

	return u.alamatRepo.Delete(id)
}
