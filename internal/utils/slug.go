// ============================================================================
// Project Name : GoShop API
// File         : slug.go
// Description  : Utility untuk generate URL-friendly slug
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi fungsi untuk generate slug dari string
// - Mengonversi ke lowercase dan replace spasi dengan hyphen
// - Menghapus karakter special dan multiple hyphen
//
// ============================================================================

package utils

import (
	"regexp"
	"strings"
)

// GenerateSlug generates URL-friendly slug from string
func GenerateSlug(s string) string {
	// Convert to lowercase
	slug := strings.ToLower(s)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters
	reg := regexp.MustCompile("[^a-z0-9-]+")
	slug = reg.ReplaceAllString(slug, "")

	// Remove multiple consecutive hyphens
	reg = regexp.MustCompile("-+")
	slug = reg.ReplaceAllString(slug, "-")

	// Trim hyphens from start and end
	slug = strings.Trim(slug, "-")

	return slug
}
