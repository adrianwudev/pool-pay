package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pool-pay/internal/domain"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) AddUser(username, email, password string) error {
	args := m.Called(username, email, password)
	return args.Error(0)
}

func (m *MockUserRepo) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepo) CheckIfExists(user domain.User) (bool, error) {
	args := m.Called(user)
	return args.Bool(0), args.Error(1)
}

func TestGetUserByEmail(t *testing.T) {
	// Create mock user repo
	mockRepo := new(MockUserRepo)

	// Expected result
	expectedUser := &domain.User{
		Id:       123,
		Username: "john",
		Email:    "john@example.com",
	}
	mockRepo.On("GetByEmail", "john@example.com").Return(expectedUser, nil)

	// Create user service with mock repo
	userService := domain.NewUserService(mockRepo)

	// Call GetByEmail method
	user, err := userService.GetByEmail("john@example.com")

	// Check if the result matches the expected user
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	// Verify that GetByEmail method was called with the correct argument
	mockRepo.AssertCalled(t, "GetByEmail", "john@example.com")
}
