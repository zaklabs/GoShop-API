// ============================================================================
// Project Name : GoShop API
// File         : main.go
// Description  : Entry point aplikasi GoShop API Server
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini menginisialisasi semua komponen aplikasi
// - Mengatur koneksi database, repositories, usecases, dan handlers
// - Menjalankan HTTP server dengan Gin framework
//
// ============================================================================

package main

import (
	"evermos-api/internal/config"
	"evermos-api/internal/delivery/http"
	"evermos-api/internal/delivery/http/handler"
	"evermos-api/internal/repository"
	"evermos-api/internal/usecase"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := config.InitDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run auto migration
	if err := config.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Create uploads directory if not exists
	if err := os.MkdirAll(cfg.Upload.Path, os.ModePerm); err != nil {
		log.Fatalf("Failed to create uploads directory: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	tokoRepo := repository.NewTokoRepository(db)
	alamatRepo := repository.NewAlamatRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	produkRepo := repository.NewProdukRepository(db)
	fotoProdukRepo := repository.NewFotoProdukRepository(db)
	logProdukRepo := repository.NewLogProdukRepository(db)
	trxRepo := repository.NewTrxRepository(db)
	detailTrxRepo := repository.NewDetailTrxRepository(db)

	// Initialize usecases
	authUsecase := usecase.NewAuthUsecase(userRepo, tokoRepo, db)
	tokoUsecase := usecase.NewTokoUsecase(tokoRepo)
	alamatUsecase := usecase.NewAlamatUsecase(alamatRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	produkUsecase := usecase.NewProdukUsecase(produkRepo, tokoRepo, fotoProdukRepo, logProdukRepo, db)
	trxUsecase := usecase.NewTrxUsecase(trxRepo, detailTrxRepo, produkRepo, logProdukRepo, alamatRepo, db)
	wilayahUsecase := usecase.NewWilayahUsecase()
	userUsecase := usecase.NewUserUsecase(userRepo, wilayahUsecase)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authUsecase, cfg.JWT.Secret, cfg.JWT.ExpireHours)
	userHandler := handler.NewUserHandler(userUsecase)
	tokoHandler := handler.NewTokoHandler(tokoUsecase, cfg.Upload.Path)
	alamatHandler := handler.NewAlamatHandler(alamatUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	produkHandler := handler.NewProdukHandler(produkUsecase, cfg.Upload.Path)
	trxHandler := handler.NewTrxHandler(trxUsecase)
	wilayahHandler := handler.NewWilayahHandler(wilayahUsecase)

	// Initialize router
	router := http.NewRouter(
		authHandler,
		userHandler,
		tokoHandler,
		alamatHandler,
		categoryHandler,
		produkHandler,
		trxHandler,
		wilayahHandler,
		cfg.JWT.Secret,
	)

	// Setup Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Setup routes
	router.SetupRoutes(r)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
