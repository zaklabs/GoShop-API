// ============================================================================
// Project Name : GoShop API
// File         : upload.go
// Description  : Utility untuk upload file
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi fungsi untuk handle upload file
// - Validasi ekstensi file (jpg, jpeg, png, gif)
// - Generate unique filename menggunakan MD5 hash
//
// ============================================================================

package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}

// UploadFile handles file upload
func UploadFile(file *multipart.FileHeader, uploadPath string, subDir string) (string, error) {
	// Create directory if not exists
	dir := filepath.Join(uploadPath, subDir)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		return "", fmt.Errorf("file type not allowed: %s", ext)
	}

	// Generate unique filename
	// timestamp := time.Now().Unix()
	// filename := fmt.Sprintf("%d-%s", timestamp, file.Filename)
	// filePath := filepath.Join(dir, filename)

	// Generate unique filename with hash
	uniqueFilename := generateUniqueFilename(file.Filename)
	filePath := filepath.Join(dir, uniqueFilename)

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// Copy file
	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Return relative path
	// return filepath.Join(subDir, filename), nil
	return filepath.Join(subDir, uniqueFilename), nil
}

// generateUniqueFilename generates unique filename using timestamp and random hash
func generateUniqueFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	timestamp := time.Now().UnixNano()
	randomNum := rand.Int63()

	// Create hash from timestamp + random number
	hash := md5.Sum([]byte(fmt.Sprintf("%d%d", timestamp, randomNum)))
	hashString := fmt.Sprintf("%x", hash)

	return fmt.Sprintf("%s%s", hashString, ext)
}

// DeleteFile deletes a file
func DeleteFile(uploadPath string, filename string) error {
	if filename == "" {
		return nil
	}

	filePath := filepath.Join(uploadPath, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}

	return os.Remove(filePath)
}
