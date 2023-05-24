package test

import (
	"testing"

	"pool-pay/config"
	"pool-pay/db"
	"pool-pay/internal/domain"
	"pool-pay/internal/repository"
)

func TestGetUserByIdAndPassword(t *testing.T) {
	config := config.GetConfig()
	dbInstance, err := db.NewConnection(&config)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	userRepo := repository.NewDbUserRepository(dbInstance)

	user, err := userRepo.GetById(123)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	expectedUser := domain.User{
		Id:       123,
		Username: "john",
		Password: "secret",
		Email:    "john@example.com",
	}

	if user.Id != expectedUser.Id {
		//||
		// user.Username != expectedUser.Username ||
		// user.Password != expectedUser.Password ||
		// user.Email != expectedUser.Email

		t.Errorf("User information mismatch. Expected: %v, Got: %v", expectedUser, user)
	}
}
