// ============================================================================
// Project Name : GoShop API
// File         : response.go
// Description  : Standard API response format
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi struct untuk format response API
// - Menyediakan helper function untuk success dan error response
// - Format response konsisten untuk semua endpoint
//
// ============================================================================

package model

// APIResponse unified response format
type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// PaginatedResponse for list endpoints
type PaginatedResponse struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Data  interface{} `json:"data"`
}

// SuccessResponse helper
func SuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

// ErrorResponse helper
func ErrorResponse(message string, errors interface{}) APIResponse {
	return APIResponse{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	}
}
