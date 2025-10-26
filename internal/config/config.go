// ============================================================================
// Project Name : GoShop API
// File         : config.go
// Description  : Konfigurasi aplikasi dan environment variables
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini membaca konfigurasi dari file .env
// - Mengatur konfigurasi database, JWT, server, dan upload
// - Menyediakan default values untuk setiap konfigurasi
//
// ============================================================================

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration
type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
	Upload   UploadConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret      string
	ExpireHours int
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
}

// UploadConfig holds upload configuration
type UploadConfig struct {
	Path          string
	MaxUploadSize int64
}

var AppConfig *Config

// LoadConfig loads configuration from .env file
func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	expireHours, _ := strconv.Atoi(getEnv("JWT_EXPIRE_HOURS", "24"))
	maxUploadSize, _ := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "5242880"), 10, 64)

	config := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "evermos"),
		},
		JWT: JWTConfig{
			Secret:      getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
			ExpireHours: expireHours,
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8000"),
		},
		Upload: UploadConfig{
			Path:          getEnv("UPLOAD_PATH", "./uploads"),
			MaxUploadSize: maxUploadSize,
		},
	}

	AppConfig = config
	return config, nil
}

// GetDSN returns database connection string
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
