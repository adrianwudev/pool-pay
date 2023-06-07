package repository

import (
	"log"
	"pool-pay/internal/constants"
	"pool-pay/internal/domain"
	"pool-pay/internal/util"
	mypg "pool-pay/pg"
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
	tx := db.Conn.Begin()
	defer tx.Rollback()

	userFriendship := &domain.Friendship{
		UserID:    userId,
		FriendID:  friendId,
		Status:    constants.FriendShipStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	friendFriendship := &domain.Friendship{
		UserID:    friendId,
		FriendID:  userId,
		Status:    constants.FriendShipStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := tx.Create(userFriendship).Error
	if err != nil {
		log.Println(err)
		pgError := mypg.GetPgError(err)
		if mypg.IsDuplicateKeyError(pgError) {
			return util.SetApiError(constants.ERRORCODE_FRIENDREQUESTALREADYEXISTS)
		}
		return util.SetApiError(constants.ERRORCODE_FAILEDTOCREATEFRIENDSHIP)
	}
	err = tx.Create(friendFriendship).Error
	if err != nil {
		log.Println(err)
		pgError := mypg.GetPgError(err)
		if mypg.IsDuplicateKeyError(pgError) {
			return util.SetApiError(constants.ERRORCODE_FRIENDREQUESTALREADYEXISTS)
		}
		return util.SetApiError(constants.ERRORCODE_FAILEDTOCREATEFRIENDSHIP)
	}

	tx.Commit()
	return nil
}

func (db *dbFriendShipRepo) GetFriendRequests(userId int64) ([]domain.Friendship, error) {
	var friendRequests []domain.Friendship
	err := db.Conn.Where("user_id = ? AND status = ?", userId, constants.FriendShipStatusPending).Find(&friendRequests).Error
	if err != nil {
		return nil, util.SetDefaultApiError(err)
	}

	return friendRequests, nil
}

func (db *dbFriendShipRepo) GetFriendshipByID(friendshipId int64) (*domain.Friendship, error) {
	var friendship domain.Friendship
	err := db.Conn.Where("friendship_id = ? AND status = ?", friendshipId, constants.FriendShipStatusPending).First(&friendship).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.SetApiError(constants.ERRORCODE_FRIENDSHIPPENDINGNOTFOUND)
		}
		return nil, util.SetDefaultApiError(err)
	}

	return &friendship, nil
}

func (db *dbFriendShipRepo) GetFriendshipIdsByUserIds(userId, friendshipId int64) ([]int64, error) {
	// get both directions of friendship by userId + friendId
	var friendshipIds []int64
	err := db.Conn.Model(&domain.Friendship{}).
		Select("friendship_id").
		Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?) AND status = ?",
			userId, friendshipId, friendshipId, userId, constants.FriendShipStatusPending).
		Find(&friendshipIds).Error
	if err != nil {
		return nil, util.SetDefaultApiError(err)
	}

	return friendshipIds, nil
}

func (db *dbFriendShipRepo) UpdateFriendShipStatus(friendshipIds []int64) error {
	// update friendship status to approved
	tx := db.Conn.Begin()
	defer tx.Rollback()

	// Update friendship status to approved
	err := tx.Model(&domain.Friendship{}).
		Where("friendship_id IN ?", friendshipIds).
		Update("status", constants.FriendShipStatusAccepted).
		Error
	if err != nil {
		return util.SetApiError(constants.ERRORCODE_UPDATEFRIENDREQUESTFAILED)
	}

	tx.Commit()
	return nil
}
