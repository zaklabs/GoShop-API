// ============================================================================
// Project Name : GoShop API
// File         : toko_usecase.go
// Description  : Business logic untuk manajemen toko
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi logic untuk CRUD toko
// - Menangani upload foto toko
// - Validasi kepemilikan toko
//
// ============================================================================

package usecase

import (
	"errors"
	"evermos-api/internal/model"
	"evermos-api/internal/repository"
	"evermos-api/internal/utils"
	"mime/multipart"
	"time"
)

// TokoUsecase interface
type TokoUsecase interface {
	GetMyToko(userID int) (*model.TokoResponse, error)
	GetTokoByID(id int) (*model.TokoResponse, error)
	GetAllToko(limit, offset int, nama string) (*model.PaginatedResponse, error)
	UpdateToko(tokoID, userID int, namaToko string, file *multipart.FileHeader, uploadPath string) error
}

type tokoUsecase struct {
	tokoRepo repository.TokoRepository
}

// NewTokoUsecase creates new toko usecase
func NewTokoUsecase(tokoRepo repository.TokoRepository) TokoUsecase {
	return &tokoUsecase{tokoRepo: tokoRepo}
}

func (u *tokoUsecase) GetMyToko(userID int) (*model.TokoResponse, error) {
	toko, err := u.tokoRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("toko not found")
	}

	return &model.TokoResponse{
		ID:       toko.ID,
		NamaToko: toko.NamaToko,
		URLFoto:  toko.URLFoto,
		UserID:   toko.IDUser,
	}, nil
}

func (u *tokoUsecase) GetTokoByID(id int) (*model.TokoResponse, error) {
	toko, err := u.tokoRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("`Toko tidak ditemukan`")
	}

	return &model.TokoResponse{
		ID:       toko.ID,
		NamaToko: toko.NamaToko,
		URLFoto:  toko.URLFoto,
	}, nil
}

func (u *tokoUsecase) GetAllToko(limit, offset int, nama string) (*model.PaginatedResponse, error) {
	tokos, err := u.tokoRepo.FindAll(limit, offset, nama)
	if err != nil {
		return nil, err
	}

	var tokoResponses []model.TokoResponse
	for _, toko := range tokos {
		tokoResponses = append(tokoResponses, model.TokoResponse{
			ID:       toko.ID,
			NamaToko: toko.NamaToko,
			URLFoto:  toko.URLFoto,
		})
	}

	return &model.PaginatedResponse{
		Page:  (offset / limit) + 1,
		Limit: limit,
		Data:  tokoResponses,
	}, nil
}

func (u *tokoUsecase) UpdateToko(tokoID, userID int, namaToko string, file *multipart.FileHeader, uploadPath string) error {
	toko, err := u.tokoRepo.FindByID(tokoID)
	if err != nil {
		return errors.New("toko not found")
	}

	// Check ownership
	if toko.IDUser != userID {
		return errors.New("unauthorized: not your toko")
	}

	// Update nama toko if provided
	if namaToko != "" {
		toko.NamaToko = namaToko
	}

	// Handle file upload if provided
	if file != nil {
		// In production, you would use the upload utility here
		// For now, just set a placeholder
		// toko.URLFoto = file.Filename

		// Upload file using utility
		urlFoto, err := utils.UploadFile(file, uploadPath, "toko")
		if err != nil {
			return err
		}

		// Delete old photo if exists
		if toko.URLFoto != "" {
			_ = utils.DeleteFile(uploadPath, toko.URLFoto)
		}

		toko.URLFoto = urlFoto
	}

	now := time.Now()
	toko.UpdatedAt = &now

	return u.tokoRepo.Update(toko)
}
