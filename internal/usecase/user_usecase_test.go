package usecase_test

import (
	"evermos-api/internal/model"
	"evermos-api/internal/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByID(id int) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) FindByNoTelp(noTelp string) (*model.User, error) {
	args := m.Called(noTelp)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Update(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockTokoRepository is a mock implementation of TokoRepository
type MockTokoRepository struct {
	mock.Mock
}

func (m *MockTokoRepository) Create(toko *model.Toko) error {
	args := m.Called(toko)
	return args.Error(0)
}

func (m *MockTokoRepository) FindByID(id int) (*model.Toko, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Toko), args.Error(1)
}

func (m *MockTokoRepository) FindByUserID(userID int) (*model.Toko, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Toko), args.Error(1)
}

func (m *MockTokoRepository) FindAll(limit, offset int, nama string) ([]model.Toko, error) {
	args := m.Called(limit, offset, nama)
	return args.Get(0).([]model.Toko), args.Error(1)
}

func (m *MockTokoRepository) Update(toko *model.Toko) error {
	args := m.Called(toko)
	return args.Error(0)
}

func (m *MockTokoRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUserUsecase_GetProfile(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	userUsecase := usecase.NewUserUsecase(mockUserRepo, nil)

	userID := 1
	now := time.Now()
	tanggalLahir := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	idProvinsi := 11
	idKota := 1171

	expectedUser := &model.User{
		ID:           userID,
		Nama:         "Test User",
		NoTelp:       "08123456789",
		TanggalLahir: &tanggalLahir,
		Email:        "test@example.com",
		Pekerjaan:    "Developer",
		Tentang:      "Test about",
		IDProvinsi:   &idProvinsi,
		IDKota:       &idKota,
		CreatedAt:    &now,
	}

	mockUserRepo.On("FindByID", userID).Return(expectedUser, nil)

	profile, err := userUsecase.GetProfile(userID)

	assert.NoError(t, err)
	assert.NotNil(t, profile)
	assert.Equal(t, "Test User", profile.Nama)
	assert.Equal(t, "test@example.com", profile.Email)
	mockUserRepo.AssertExpectations(t)
}
