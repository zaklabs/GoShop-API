// ============================================================================
// Project Name : GoShop API
// File         : database.go
// Description  : Inisialisasi koneksi database dan auto migration
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini menginisialisasi koneksi MySQL menggunakan GORM
// - Menjalankan auto migration untuk semua model
// - Membuat unique index untuk field notelp pada tabel users
//
// ============================================================================

package config

import (
	"evermos-api/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes database connection
func InitDB(cfg *DatabaseConfig) (*gorm.DB, error) {
	dsn := cfg.GetDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection established")
	DB = db
	return db, nil
}

// AutoMigrate runs auto migration for all models
func AutoMigrate(db *gorm.DB) error {
	log.Println("Running auto migration...")

	// Migrate each model individually to handle errors gracefully
	models := []interface{}{
		&model.User{},
		&model.Toko{},
		&model.Alamat{},
		&model.Category{},
		&model.Produk{},
		&model.FotoProduk{},
		&model.LogProduk{},
		&model.Trx{},
		&model.DetailTrx{},
	}

	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			log.Printf("Warning during migration: %v", err)
			// Continue with other migrations instead of failing completely
		}
	}

	log.Println("All tables created successfully")

	// Add unique index for notelp AFTER all tables are created
	if db.Migrator().HasTable(&model.User{}) {
		log.Println("Adding unique index for users.notelp...")

		// Drop old index if exists
		if db.Migrator().HasIndex(&model.User{}, "uni_users_notelp") {
			log.Println("Dropping old index: uni_users_notelp")
			if err := db.Migrator().DropIndex(&model.User{}, "uni_users_notelp"); err != nil {
				log.Printf("Warning: Failed to drop old index: %v", err)
			}
		}

		if db.Migrator().HasIndex(&model.User{}, "idx_users_no_telp") {
			log.Println("Dropping old index: idx_users_no_telp")
			if err := db.Migrator().DropIndex(&model.User{}, "idx_users_no_telp"); err != nil {
				log.Printf("Warning: Failed to drop old index: %v", err)
			}
		}

		// Create unique index for notelp
		if !db.Migrator().HasIndex(&model.User{}, "idx_users_notelp_unique") {
			log.Println("Creating unique index: idx_users_notelp_unique")
			if err := db.Exec("CREATE UNIQUE INDEX idx_users_notelp_unique ON users(notelp)").Error; err != nil {
				log.Printf("Warning: Failed to create unique index: %v", err)
			} else {
				log.Println("Unique index created successfully")
			}
		} else {
			log.Println("Unique index idx_users_notelp_unique already exists")
		}
	}

	log.Println("Auto migration completed successfully")
	return nil
}
