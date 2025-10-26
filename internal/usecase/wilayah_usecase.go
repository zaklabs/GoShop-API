// ============================================================================
// Project Name : GoShop API
// File         : wilayah_usecase.go
// Description  : Business logic untuk mendapatkan data wilayah
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi logic untuk mengambil data provinsi dan kota
// - Menggunakan HTTP client untuk call API eksternal
// - API source: http://www.emsifa.com/api-wilayah-indonesia
//
// ============================================================================

package usecase

import (
	"encoding/json"
	"errors"
	"evermos-api/internal/model"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseAPIURL = "https://www.emsifa.com/api-wilayah-indonesia/api"
)

// WilayahUsecase interface
type WilayahUsecase interface {
	GetListProvince() ([]model.Province, error)
	GetDetailProvince(id string) (*model.Province, error)
	GetListCity(provinceID string) ([]model.City, error)
	GetDetailCity(id string) (*model.City, error)
}

type wilayahUsecase struct {
	httpClient *http.Client
}

// NewWilayahUsecase creates new wilayah usecase
func NewWilayahUsecase() WilayahUsecase {
	return &wilayahUsecase{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetListProvince gets all provinces
func (u *wilayahUsecase) GetListProvince() ([]model.Province, error) {
	url := fmt.Sprintf("%s/provinces.json", baseAPIURL)

	resp, err := u.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch provinces: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var provinces []model.Province
	if err := json.Unmarshal(body, &provinces); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return provinces, nil
}

// GetDetailProvince gets province by ID
func (u *wilayahUsecase) GetDetailProvince(id string) (*model.Province, error) {
	url := fmt.Sprintf("%s/province/%s.json", baseAPIURL, id)

	resp, err := u.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch province: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("province not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var province model.Province
	if err := json.Unmarshal(body, &province); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &province, nil
}

// GetListCity gets all cities/regencies by province ID
func (u *wilayahUsecase) GetListCity(provinceID string) ([]model.City, error) {
	url := fmt.Sprintf("%s/regencies/%s.json", baseAPIURL, provinceID)

	resp, err := u.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cities: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("province not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var cities []model.City
	if err := json.Unmarshal(body, &cities); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return cities, nil
}

// GetDetailCity gets city/regency by ID
func (u *wilayahUsecase) GetDetailCity(id string) (*model.City, error) {
	url := fmt.Sprintf("%s/regency/%s.json", baseAPIURL, id)

	resp, err := u.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch city: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("city not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var city model.City
	if err := json.Unmarshal(body, &city); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &city, nil
}
