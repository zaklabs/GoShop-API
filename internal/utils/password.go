// ============================================================================
// Project Name : GoShop API
// File         : password.go
// Description  : Utility untuk hashing dan verifikasi password
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi fungsi untuk hash dan compare password
// - Menggunakan bcrypt untuk keamanan password
// - Default cost digunakan untuk hashing
//
// ============================================================================

package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
