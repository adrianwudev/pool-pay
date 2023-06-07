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
	GetFriendRequests(userId int64) ([]Friendship, error)
	GetFriendshipByID(friendshipId int64) (*Friendship, error)
	GetFriendshipIdsByUserIds(userId, friendshipId int64) ([]int64, error)
	UpdateFriendShipStatus(friendshipIds []int64) error
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

func (s *FriendshipService) GetFriendRequests(userId int64) ([]Friendship, error) {
	return s.FriendshipRepo.GetFriendRequests(userId)
}

func (s *FriendshipService) ApproveRequest(friendshipId int64) error {
	// get userId and friendId by friendshipId
	friendship, err := s.FriendshipRepo.GetFriendshipByID(friendshipId)
	if err != nil {
		return err
	}

	// get both directions of friendship by userId + friendId
	friendshipIds, err := s.FriendshipRepo.GetFriendshipIdsByUserIds(friendship.UserID, friendship.FriendID)
	if err != nil {
		return err
	}

	err = s.FriendshipRepo.UpdateFriendShipStatus(friendshipIds)
	if err != nil {
		return err
	}

	return nil
}
