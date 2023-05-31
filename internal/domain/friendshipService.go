package domain

import (
	"time"
)

type Friendship struct {
	FriendshipID int64     `gorm:"primaryKey" json:"friendship_id"`
	UserID       int64     `gorm:"not null" json:"user_id"`
	FriendID     int64     `gorm:"not null" json:"friend_id"`
	Status       string    `gorm:"not null" json:"status"`
	CreatedAt    time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

func (Friendship) TableName() string {
	return "friendships"
}

type FriendshipRepo interface {
	AddFriendRequest(userId, FriendId int64) error
}

type FriendshipService struct {
	FriendshipRepo FriendshipRepo
}

func NewFriendService(friendShipRepo FriendshipRepo) *FriendshipService {
	return &FriendshipService{
		FriendshipRepo: friendShipRepo,
	}
}

func (s *FriendshipService) AddFriendRequest(userId, friendId int64) error {
	return s.FriendshipRepo.AddFriendRequest(userId, friendId)
}
