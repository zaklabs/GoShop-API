package handler_test

import (
	"bytes"
	"encoding/json"
	"evermos-api/internal/delivery/http/handler"
	"evermos-api/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock AuthUsecase
type MockAuthUsecase struct {
	mock.Mock
}

func (m *MockAuthUsecase) Register(req model.RegisterRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockAuthUsecase) Login(req model.LoginRequest, jwtSecret string, expireHours int) (*model.LoginResponse, error) {
	args := m.Called(req, jwtSecret, expireHours)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.LoginResponse), args.Error(1)
}

// Mock CategoryUsecase
type MockCategoryUsecase struct {
	mock.Mock
}

func (m *MockCategoryUsecase) CreateCategory(req model.CategoryRequest) (int, error) {
	args := m.Called(req)
	return args.Int(0), args.Error(1)
}

func (m *MockCategoryUsecase) GetCategoryByID(id int) (*model.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryUsecase) GetAllCategory() ([]model.Category, error) {
	args := m.Called()
	return args.Get(0).([]model.Category), args.Error(1)
}

func (m *MockCategoryUsecase) UpdateCategory(id int, req model.CategoryRequest) error {
	args := m.Called(id, req)
	return args.Error(0)
}

func (m *MockCategoryUsecase) DeleteCategory(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestAuthHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Register Success", func(t *testing.T) {
		// Setup mock
		mockAuthUsecase := new(MockAuthUsecase)
		authHandler := handler.NewAuthHandler(mockAuthUsecase, "test-secret", 24)

		reqBody := model.RegisterRequest{
			Nama:         "Test User",
			KataSandi:    "password123",
			NoTelp:       "08123456789",
			Email:        "test@example.com",
			TanggalLahir: "01/01/1990",
			Pekerjaan:    "Developer",
			IDProvinsi:   "11",
			IDKota:       "1101",
		}

		mockAuthUsecase.On("Register", reqBody).Return(nil)

		// Setup router
		r := gin.New()
		r.POST("/register", authHandler.Register)

		// Make request
		jsonBody, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response model.APIResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.True(t, response.Status)
		assert.Equal(t, "Succeed to POST data", response.Message)

		mockAuthUsecase.AssertExpectations(t)
	})

	t.Run("Login Success", func(t *testing.T) {
		// Setup mock
		mockAuthUsecase := new(MockAuthUsecase)
		authHandler := handler.NewAuthHandler(mockAuthUsecase, "test-secret", 24)

		reqBody := model.LoginRequest{
			NoTelp:    "08123456789",
			KataSandi: "password123",
		}

		loginResponse := &model.LoginResponse{
			UserResponse: model.UserResponse{
				Nama:   "Test User",
				NoTelp: "08123456789",
				Email:  "test@example.com",
			},
			Token: "mock-jwt-token",
		}

		mockAuthUsecase.On("Login", reqBody, "test-secret", 24).Return(loginResponse, nil)

		// Setup router
		r := gin.New()
		r.POST("/login", authHandler.Login)

		// Make request
		jsonBody, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response model.APIResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.True(t, response.Status)
		assert.NotNil(t, response.Data)

		mockAuthUsecase.AssertExpectations(t)
	})
}

func TestCategoryHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Create Category Success", func(t *testing.T) {
		// Setup mock
		mockCategoryUsecase := new(MockCategoryUsecase)
		categoryHandler := handler.NewCategoryHandler(mockCategoryUsecase)

		reqBody := model.CategoryRequest{
			NamaCategory: "Electronics",
		}

		mockCategoryUsecase.On("CreateCategory", reqBody).Return(1, nil)

		// Setup router
		r := gin.New()
		r.POST("/category", categoryHandler.CreateCategory)

		// Make request
		jsonBody, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/category", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response model.APIResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.True(t, response.Status)
		assert.NotNil(t, response.Data)

		mockCategoryUsecase.AssertExpectations(t)
	})

	t.Run("Get All Categories Success", func(t *testing.T) {
		// Setup mock
		mockCategoryUsecase := new(MockCategoryUsecase)
		categoryHandler := handler.NewCategoryHandler(mockCategoryUsecase)

		expectedCategories := []model.Category{
			{ID: 1, NamaCategory: "Electronics"},
			{ID: 2, NamaCategory: "Fashion"},
		}

		mockCategoryUsecase.On("GetAllCategory").Return(expectedCategories, nil)

		// Setup router
		r := gin.New()
		r.GET("/category", categoryHandler.GetAllCategory)

		// Make request
		req, _ := http.NewRequest("GET", "/category", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response model.APIResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.True(t, response.Status)
		assert.NotNil(t, response.Data)

		mockCategoryUsecase.AssertExpectations(t)
	})
}
