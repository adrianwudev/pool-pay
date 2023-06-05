package service

import (
	"pool-pay/internal/domain"
	"pool-pay/internal/repository"

	"gorm.io/gorm"
)

func GetUserService(db *gorm.DB) *domain.UserService {
	userRepo := repository.NewDbUserRepository(db)
	userService := domain.NewUserService(userRepo)
	return userService
}

func GetFriendshipService(db *gorm.DB) *domain.FriendshipService {
	friendShipRepo := repository.NewDbFriendshipRepository(db)
	friendShipService := domain.NewFriendService(friendShipRepo)
	return friendShipService
}
