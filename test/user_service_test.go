package test

import (
	"errors"
	"pool-pay/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (m *MockUserRepo) Login(email, password string) (bool, error) {
	args := m.Called(email, password)
	return args.Bool(0), args.Error(1)
}

func TestUserService_Login(t *testing.T) {
	mockUserRepo := new(MockUserRepo)
	mockUserRepo.On("Login", "test@example.com", "password").Return(true, nil)

	userService := domain.NewUserService(mockUserRepo)

	token, err := userService.Login("test@example.com", "password")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	mockUserRepo.AssertCalled(t, "Login", "test@example.com", "password")
}

func TestUserService_Login_LoginError(t *testing.T) {
	mockUserRepo := new(MockUserRepo)
	mockUserRepo.On("Login", "test@example.com", "password").Return(false, errors.New("database error"))

	userService := domain.NewUserService(mockUserRepo)

	token, err := userService.Login("test@example.com", "password")

	assert.EqualError(t, err, "database error")
	assert.Empty(t, token)

	mockUserRepo.AssertCalled(t, "Login", "test@example.com", "password")
}
