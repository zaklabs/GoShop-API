// ============================================================================
// Project Name : GoShop API
// File         : produk_usecase.go
// Description  : Business logic untuk manajemen produk
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi logic untuk CRUD produk
// - Menangani upload multiple foto produk
// - Generate slug otomatis dari nama produk
//
// ============================================================================

package usecase

import (
	"errors"
	"evermos-api/internal/model"
	"evermos-api/internal/repository"
	"evermos-api/internal/utils"
	"mime/multipart"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// ProdukUsecase interface
type ProdukUsecase interface {
	GetAllProduk(limit, offset int, filters map[string]string) (*model.PaginatedResponse, error)
	GetProdukByID(id int) (*model.Produk, error)
	CreateProduk(userID int, req model.CreateProdukRequest, files []*multipart.FileHeader, uploadPath string) (int, error)
	UpdateProduk(id, userID int, req model.UpdateProdukRequest, files []*multipart.FileHeader, uploadPath string) error
	DeleteProduk(id, userID int) error
}

type produkUsecase struct {
	produkRepo     repository.ProdukRepository
	tokoRepo       repository.TokoRepository
	fotoProdukRepo repository.FotoProdukRepository
	logProdukRepo  repository.LogProdukRepository
	db             *gorm.DB
}

// NewProdukUsecase creates new produk usecase
func NewProdukUsecase(
	produkRepo repository.ProdukRepository,
	tokoRepo repository.TokoRepository,
	fotoProdukRepo repository.FotoProdukRepository,
	logProdukRepo repository.LogProdukRepository,
	db *gorm.DB,
) ProdukUsecase {
	return &produkUsecase{
		produkRepo:     produkRepo,
		tokoRepo:       tokoRepo,
		fotoProdukRepo: fotoProdukRepo,
		logProdukRepo:  logProdukRepo,
		db:             db,
	}
}

func (u *produkUsecase) GetAllProduk(limit, offset int, filters map[string]string) (*model.PaginatedResponse, error) {
	// Convert string filters to appropriate types
	filterMap := make(map[string]interface{})

	if namaProduk, ok := filters["nama_produk"]; ok {
		filterMap["nama_produk"] = namaProduk
	}
	if categoryID, ok := filters["category_id"]; ok {
		if id, err := strconv.Atoi(categoryID); err == nil {
			filterMap["category_id"] = id
		}
	}
	if tokoID, ok := filters["toko_id"]; ok {
		if id, err := strconv.Atoi(tokoID); err == nil {
			filterMap["toko_id"] = id
		}
	}
	if minHarga, ok := filters["min_harga"]; ok {
		if price, err := strconv.Atoi(minHarga); err == nil {
			filterMap["min_harga"] = price
		}
	}
	if maxHarga, ok := filters["max_harga"]; ok {
		if price, err := strconv.Atoi(maxHarga); err == nil {
			filterMap["max_harga"] = price
		}
	}

	produks, err := u.produkRepo.FindAll(limit, offset, filterMap)
	if err != nil {
		return nil, err
	}

	return &model.PaginatedResponse{
		Page:  (offset / limit) + 1,
		Limit: limit,
		Data:  produks,
	}, nil
}

func (u *produkUsecase) GetProdukByID(id int) (*model.Produk, error) {
	produk, err := u.produkRepo.FindByIDWithRelations(id)
	if err != nil {
		return nil, errors.New("No Data Product")
	}
	return produk, nil
}

func (u *produkUsecase) CreateProduk(userID int, req model.CreateProdukRequest, files []*multipart.FileHeader, uploadPath string) (int, error) {
	// Get user's toko
	toko, err := u.tokoRepo.FindByUserID(userID)
	if err != nil {
		return 0, errors.New("you don't have a toko")
	}

	now := time.Now()
	slug := utils.GenerateSlug(req.NamaProduk)

	produk := &model.Produk{
		NamaProduk:    req.NamaProduk,
		Slug:          slug,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
		IDToko:        toko.ID,
		IDCategory:    req.CategoryID,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	return produk.ID, u.db.Transaction(func(tx *gorm.DB) error {
		// Create produk
		if err := tx.Create(produk).Error; err != nil {
			return err
		}

		// Create foto produk if files provided
		if len(files) > 0 {
			for _, file := range files {
				// Upload file using utility
				urlFoto, err := utils.UploadFile(file, uploadPath, "produk")
				if err != nil {
					return err
				}

				foto := &model.FotoProduk{
					IDProduk: produk.ID,
					//URL:       file.Filename, // In production, use upload utility
					URL:       urlFoto,
					CreatedAt: &now,
					UpdatedAt: &now,
				}
				if err := tx.Create(foto).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (u *produkUsecase) UpdateProduk(id, userID int, req model.UpdateProdukRequest, files []*multipart.FileHeader, uploadPath string) error {
	produk, err := u.produkRepo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	// Check ownership
	toko, err := u.tokoRepo.FindByID(produk.IDToko)
	if err != nil {
		return errors.New("toko not found")
	}
	if toko.IDUser != userID {
		return errors.New("unauthorized: not your product")
	}

	// Update fields if provided
	if req.NamaProduk != "" {
		produk.NamaProduk = req.NamaProduk
		produk.Slug = utils.GenerateSlug(req.NamaProduk)
	}
	if req.HargaReseller != "" {
		produk.HargaReseller = req.HargaReseller
	}
	if req.HargaKonsumen != "" {
		produk.HargaKonsumen = req.HargaKonsumen
	}
	if req.Stok > 0 {
		produk.Stok = req.Stok
	}
	if req.Deskripsi != "" {
		produk.Deskripsi = req.Deskripsi
	}
	if req.CategoryID > 0 {
		produk.IDCategory = req.CategoryID
	}

	now := time.Now()
	produk.UpdatedAt = &now

	return u.db.Transaction(func(tx *gorm.DB) error {
		// Update produk
		if err := tx.Save(produk).Error; err != nil {
			return err
		}

		// Handle file uploads if provided
		if len(files) > 0 {
			// Get old photos for deletion
			oldPhotos, _ := u.fotoProdukRepo.FindByProdukID(id)

			// Delete old photos from database
			if err := tx.Where("id_produk = ?", id).Delete(&model.FotoProduk{}).Error; err != nil {
				return err
			}

			// Delete old photo files
			for _, oldPhoto := range oldPhotos {
				_ = utils.DeleteFile(uploadPath, oldPhoto.URL)
			}

			// Upload and create new photos
			for _, file := range files {
				// Upload file using utility
				urlFoto, err := utils.UploadFile(file, uploadPath, "produk")
				if err != nil {
					return err
				}

				foto := &model.FotoProduk{
					IDProduk: produk.ID,
					//URL:       file.Filename,
					URL:       urlFoto,
					CreatedAt: &now,
					UpdatedAt: &now,
				}
				if err := tx.Create(foto).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (u *produkUsecase) DeleteProduk(id, userID int) error {
	produk, err := u.produkRepo.FindByID(id)
	if err != nil {
		return errors.New("record not found")
	}

	// Check ownership
	toko, err := u.tokoRepo.FindByID(produk.IDToko)
	if err != nil {
		return errors.New("toko not found")
	}
	if toko.IDUser != userID {
		return errors.New("unauthorized: not your product")
	}

	// Check if product has been used in transactions
	hasTransaction, err := u.logProdukRepo.ExistsByProdukID(id)
	if err != nil {
		return errors.New("failed to check product transactions")
	}

	// If product has transactions, use soft delete
	if hasTransaction {
		now := time.Now()
		produk.DeletedAt = &now
		return u.produkRepo.Update(produk)
	}

	// If no transactions, perform hard delete
	return u.db.Transaction(func(tx *gorm.DB) error {
		// Get photos for deletion
		photos, _ := u.fotoProdukRepo.FindByProdukID(id)

		// Delete photos from database first
		if err := tx.Where("id_produk = ?", id).Delete(&model.FotoProduk{}).Error; err != nil {
			return err
		}

		// Delete photo files
		// Note: uploadPath should be passed or obtained from config
		// For now, we'll skip physical deletion in delete operation
		// You can add uploadPath as parameter if needed
		_ = photos

		// Delete product
		if err := tx.Delete(&model.Produk{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}
