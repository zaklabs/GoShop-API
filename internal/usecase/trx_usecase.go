// ============================================================================
// Project Name : GoShop API
// File         : trx_usecase.go
// Description  : Business logic untuk manajemen transaksi
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi logic untuk membuat dan mendapatkan transaksi
// - Generate invoice otomatis dengan format INV-YYYYMMDD-XXXX
// - Membuat snapshot produk dalam log_produk
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

// TrxUsecase interface
type TrxUsecase interface {
	GetAllTrx(userID int, limit, offset int) ([]model.Trx, error)
	GetTrxByID(id, userID int) (*model.Trx, error)
	CreateTrx(userID int, req model.CreateTrxRequest) (int, error)
}

type trxUsecase struct {
	trxRepo       repository.TrxRepository
	detailTrxRepo repository.DetailTrxRepository
	produkRepo    repository.ProdukRepository
	logProdukRepo repository.LogProdukRepository
	alamatRepo    repository.AlamatRepository
	db            *gorm.DB
}

// NewTrxUsecase creates new trx usecase
func NewTrxUsecase(
	trxRepo repository.TrxRepository,
	detailTrxRepo repository.DetailTrxRepository,
	produkRepo repository.ProdukRepository,
	logProdukRepo repository.LogProdukRepository,
	alamatRepo repository.AlamatRepository,
	db *gorm.DB,
) TrxUsecase {
	return &trxUsecase{
		trxRepo:       trxRepo,
		detailTrxRepo: detailTrxRepo,
		produkRepo:    produkRepo,
		logProdukRepo: logProdukRepo,
		alamatRepo:    alamatRepo,
		db:            db,
	}
}

func (u *trxUsecase) GetAllTrx(userID int, limit, offset int) ([]model.Trx, error) {
	return u.trxRepo.FindByUserID(userID, limit, offset)
}

func (u *trxUsecase) GetTrxByID(id, userID int) (*model.Trx, error) {
	trx, err := u.trxRepo.FindByIDWithDetails(id)
	if err != nil {
		return nil, errors.New("`No Data Trx`")
	}

	// Check ownership
	if trx.IDUser != userID {
		return nil, errors.New("unauthorized: not your transaction")
	}

	return trx, nil
}

func (u *trxUsecase) CreateTrx(userID int, req model.CreateTrxRequest) (int, error) {
	// Validate alamat ownership
	alamat, err := u.alamatRepo.FindByID(req.AlamatPengiriman)
	if err != nil {
		return 0, errors.New("alamat not found")
	}
	if alamat.IDUser != userID {
		return 0, errors.New("unauthorized: not your alamat")
	}

	// Calculate total price and validate products
	var totalHarga int
	var details []struct {
		produk    *model.Produk
		kuantitas int
	}

	for _, detail := range req.DetailTrx {
		produk, err := u.produkRepo.FindByIDWithRelations(detail.ProductID)
		if err != nil {
			return 0, errors.New("product not found: " + strconv.Itoa(detail.ProductID))
		}

		// Check if product is deleted (soft delete)
		if produk.DeletedAt != nil {
			return 0, errors.New("product is no longer available: " + produk.NamaProduk)
		}

		// Check stock
		if produk.Stok < detail.Kuantitas {
			return 0, errors.New("insufficient stock for product: " + produk.NamaProduk)
		}

		// Calculate price
		hargaKonsumen, _ := strconv.Atoi(produk.HargaKonsumen)
		detailTotal := hargaKonsumen * detail.Kuantitas
		totalHarga += detailTotal

		details = append(details, struct {
			produk    *model.Produk
			kuantitas int
		}{
			produk:    produk,
			kuantitas: detail.Kuantitas,
		})
	}

	// Generate invoice code
	kodeInvoice := utils.GenerateInvoiceCode()

	now := time.Now()
	trx := &model.Trx{
		IDUser:           userID,
		AlamatPengiriman: req.AlamatPengiriman,
		HargaTotal:       totalHarga,
		KodeInvoice:      kodeInvoice,
		MethodBayar:      req.MethodBayar,
		CreatedAt:        &now,
		UpdatedAt:        &now,
	}

	// Use transaction to create trx, detail_trx, and log_produk
	return trx.ID, u.db.Transaction(func(tx *gorm.DB) error {
		// Create trx
		if err := tx.Create(trx).Error; err != nil {
			return err
		}

		// Create detail_trx and log_produk for each product
		for _, detail := range details {
			// Create log_produk (snapshot)
			logProduk := &model.LogProduk{
				IDProduk:      detail.produk.ID,
				NamaProduk:    detail.produk.NamaProduk,
				Slug:          detail.produk.Slug,
				HargaReseller: detail.produk.HargaReseller,
				HargaKonsumen: detail.produk.HargaKonsumen,
				Deskripsi:     detail.produk.Deskripsi,
				IDToko:        detail.produk.IDToko,
				IDCategory:    detail.produk.IDCategory,
				CreatedAt:     &now,
				UpdatedAt:     &now,
			}
			if err := tx.Create(logProduk).Error; err != nil {
				return err
			}

			// Calculate detail price
			hargaKonsumen, _ := strconv.Atoi(detail.produk.HargaKonsumen)
			detailHarga := hargaKonsumen * detail.kuantitas

			// Create detail_trx
			detailTrx := &model.DetailTrx{
				IDTrx:       trx.ID,
				IDLogProduk: logProduk.ID,
				IDToko:      detail.produk.IDToko,
				Kuantitas:   detail.kuantitas,
				HargaTotal:  detailHarga,
				CreatedAt:   &now,
				UpdatedAt:   &now,
			}
			if err := tx.Create(detailTrx).Error; err != nil {
				return err
			}

			// Update product stock
			detail.produk.Stok -= detail.kuantitas
			if err := tx.Save(detail.produk).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
