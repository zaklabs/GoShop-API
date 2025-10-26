// ============================================================================
// Project Name : GoShop API
// File         : auth.go
// Description  : Middleware untuk autentikasi JWT
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi middleware untuk validasi JWT token
// - Memeriksa Authorization header dengan format Bearer token
// - Menyimpan user info ke dalam context untuk digunakan handler
//
// ============================================================================

package middleware

import (
	"evermos-api/internal/model"
	"evermos-api/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse(
				"Failed to GET data",
				[]string{"Unauthorized"},
			))
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse(
				"Failed to GET data",
				[]string{"Invalid authorization header format"},
			))
			c.Abort()
			return
		}

		token := parts[1]

		// Validate token
		claims, err := utils.ValidateToken(token, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse(
				"Failed to GET data",
				[]string{"Invalid or expired token"},
			))
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("isAdmin", claims.IsAdmin)

		c.Next()
	}
}

// AdminMiddleware checks if user is admin
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("isAdmin")

		if !exists {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse(
				"Failed to POST data",
				[]string{"Unauthorized - isAdmin not found in context"},
			))
			c.Abort()
			return
		}

		isAdminBool := isAdmin.(bool)

		// Check if user is admin
		if !isAdminBool {
			// Return error dengan info lebih detail
			c.JSON(http.StatusForbidden, model.ErrorResponse(
				"Failed to POST data",
				[]string{"Admin access required. Please login with admin account."},
			))
			c.Abort()
			return
		}

		c.Next()
	}
}

// GetUserID extracts userID from context
func GetUserID(c *gin.Context) int {
	userID, exists := c.Get("userID")
	if !exists {
		return 0
	}
	return userID.(int)
}

// GetIsAdmin extracts isAdmin from context
func GetIsAdmin(c *gin.Context) bool {
	isAdmin, exists := c.Get("isAdmin")
	if !exists {
		return false
	}
	return isAdmin.(bool)
}
