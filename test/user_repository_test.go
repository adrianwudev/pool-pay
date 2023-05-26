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

func TestGetUserByIdAndPassword(t *testing.T) {
	// 建立偽造的資料庫物件
	mockRepo := new(MockUserRepo)

	// 在測試前設定期望的回傳值
	expectedUser := &domain.User{
		Id:       123,
		Username: "john",
		Email:    "john@example.com",
	}
	mockRepo.On("GetByEmail", "john@example.com").Return(expectedUser, nil)

	// 執行測試
	userService := domain.NewUserService(mockRepo)
	user, err := userService.GetByEmail("john@example.com")

	// 驗證結果是否符合期望
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	// 驗證偽造物件的方法是否被呼叫過
	mockRepo.AssertExpectations(t)
}
