// ============================================================================
// Project Name : GoShop API
// File         : router.go
// Description  : Setup routing untuk semua endpoint API
// Author       : Zaki Fuadi
// Version      : v1.0
// License      : MIT
// ============================================================================
//
// Notes:
// - File ini berisi konfigurasi semua routes API
// - Menggunakan Gin framework untuk routing
// - Menerapkan middleware untuk auth, CORS, dan logging
//
// ============================================================================

package http

import (
	"evermos-api/internal/delivery/http/handler"
	"evermos-api/internal/delivery/middleware"

	"github.com/gin-gonic/gin"
)

// Router sets up all routes
type Router struct {
	authHandler     *handler.AuthHandler
	userHandler     *handler.UserHandler
	tokoHandler     *handler.TokoHandler
	alamatHandler   *handler.AlamatHandler
	categoryHandler *handler.CategoryHandler
	produkHandler   *handler.ProdukHandler
	trxHandler      *handler.TrxHandler
	wilayahHandler  *handler.WilayahHandler
	jwtSecret       string
}

// NewRouter creates new router
func NewRouter(
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	tokoHandler *handler.TokoHandler,
	alamatHandler *handler.AlamatHandler,
	categoryHandler *handler.CategoryHandler,
	produkHandler *handler.ProdukHandler,
	trxHandler *handler.TrxHandler,
	wilayahHandler *handler.WilayahHandler,
	jwtSecret string,
) *Router {
	return &Router{
		authHandler:     authHandler,
		userHandler:     userHandler,
		tokoHandler:     tokoHandler,
		alamatHandler:   alamatHandler,
		categoryHandler: categoryHandler,
		produkHandler:   produkHandler,
		trxHandler:      trxHandler,
		wilayahHandler:  wilayahHandler,
		jwtSecret:       jwtSecret,
	}
}

// SetupRoutes configures all routes
func (r *Router) SetupRoutes(router *gin.Engine) {
	// Apply global middlewares
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.ErrorHandlerMiddleware())

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Auth routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", r.authHandler.Register)
			auth.POST("/login", r.authHandler.Login)
		}

		// Category routes
		category := v1.Group("/category")
		{
			category.GET("", r.categoryHandler.GetAllCategory)
			category.GET("/:id", r.categoryHandler.GetCategoryByID)

			// Admin only
			categoryAdmin := category.Use(middleware.AuthMiddleware(r.jwtSecret), middleware.AdminMiddleware())
			{
				categoryAdmin.POST("", r.categoryHandler.CreateCategory)
				categoryAdmin.PUT("/:id", r.categoryHandler.UpdateCategory)
				categoryAdmin.DELETE("/:id", r.categoryHandler.DeleteCategory)
			}
		}

		// Toko routes
		toko := v1.Group("/toko")
		{
			toko.GET("", r.tokoHandler.GetAllToko)
			toko.GET("/:id_toko", r.tokoHandler.GetTokoByID)

			// Authenticated routes
			tokoAuth := toko.Use(middleware.AuthMiddleware(r.jwtSecret))
			{
				tokoAuth.GET("/my", r.tokoHandler.GetMyToko)
				tokoAuth.PUT("/:id_toko", r.tokoHandler.UpdateToko)
			}
		}

		// Product routes
		product := v1.Group("/product")
		{
			product.GET("", r.produkHandler.GetAllProduk)
			product.GET("/:id", r.produkHandler.GetProdukByID)

			// Authenticated routes
			productAuth := product.Use(middleware.AuthMiddleware(r.jwtSecret))
			{
				productAuth.POST("", r.produkHandler.CreateProduk)
				productAuth.PUT("/:id", r.produkHandler.UpdateProduk)
				productAuth.DELETE("/:id", r.produkHandler.DeleteProduk)
			}
		}

		// User routes (authenticated)
		user := v1.Group("/user").Use(middleware.AuthMiddleware(r.jwtSecret))
		{
			user.GET("", r.userHandler.GetProfile)
			user.PUT("", r.userHandler.UpdateProfile)

			// Alamat routes
			user.GET("/alamat", r.alamatHandler.GetMyAlamat)
			user.GET("/alamat/:id", r.alamatHandler.GetAlamatByID)
			user.POST("/alamat", r.alamatHandler.CreateAlamat)
			user.PUT("/alamat/:id", r.alamatHandler.UpdateAlamat)
			user.DELETE("/alamat/:id", r.alamatHandler.DeleteAlamat)
		}

		// Transaction routes (authenticated)
		trx := v1.Group("/trx").Use(middleware.AuthMiddleware(r.jwtSecret))
		{
			trx.GET("", r.trxHandler.GetAllTrx)
			trx.GET("/:id", r.trxHandler.GetTrxByID)
			trx.POST("", r.trxHandler.CreateTrx)
		}

		// Wilayah routes (public)
		wilayah := v1.Group("/provcity")
		{
			wilayah.GET("/listprovincies", r.wilayahHandler.GetListProvince)
			wilayah.GET("/detailprovince/:id", r.wilayahHandler.GetDetailProvince)
			wilayah.GET("/listcities/:province_id", r.wilayahHandler.GetListCity)
			wilayah.GET("/detailcity/:id", r.wilayahHandler.GetDetailCity)
		}
	}
}
