// ============================================================================
// Project Name : GoShop API
// File         : wilayah.go
// Description  : Model dan DTO untuk data wilayah (provinsi dan kota)
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi struct untuk response provinsi dan kota
// - Data diambil dari API eksternal emsifa
// - Tidak ada database storage untuk wilayah
//
// ============================================================================

package model

// Province represents province data
type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// City represents city/regency data
type City struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}
