// ============================================================================
// Project Name : GoShop API
// File         : invoice.go
// Description  : Utility untuk generate invoice code
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi fungsi untuk generate unique invoice code
// - Format: INV-YYYYMMDD-XXXX
// - Menggunakan timestamp dan random number
//
// ============================================================================

package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateInvoiceCode generates unique invoice code
func GenerateInvoiceCode() string {
	rand.Seed(time.Now().UnixNano())
	timestamp := time.Now().Format("20060102")
	randomNum := rand.Intn(9999)
	return fmt.Sprintf("INV-%s-%04d", timestamp, randomNum)
}
