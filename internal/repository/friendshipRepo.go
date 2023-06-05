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

	tx := db.Conn.Begin()

	err := tx.Create(userFriendship).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		pgError := mypg.GetPgError(err)
		if mypg.IsDuplicateKeyError(pgError) {
			return util.SetApiError(constants.ERRORCODE_FRIENDREQUESTALREADYEXISTS)
		}
		return util.SetApiError(constants.ERRORCODE_FAILEDTOCREATEFRIENDSHIP)
	}
	err = tx.Create(friendFriendship).Error
	if err != nil {
		tx.Rollback()
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
