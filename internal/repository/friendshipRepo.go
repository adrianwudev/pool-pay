package repository

import (
	"pool-pay/internal/constants"
	"pool-pay/internal/domain"
	"time"

	"gorm.io/gorm"
)

type dbFriendShipRepo struct {
	Conn *gorm.DB
}

func NewDbFriendshipRepository(conn *gorm.DB) domain.FriendshipRepo {
	return &dbFriendShipRepo{conn}
}

func (db *dbFriendShipRepo) AddFriendRequest(userId, friendId int64) error {
	friendship := &domain.Friendship{
		UserID:    userId,
		FriendID:  friendId,
		Status:    constants.FriendShipStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := db.Conn.Create(friendship).Error
	if err != nil {
		return err
	}

	return nil
}
