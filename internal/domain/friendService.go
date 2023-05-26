package domain

import "time"

type Friend struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	UserID    int64     `gorm:"not null" json:"user_id"`
	FriendID  int64     `gorm:"not null" json:"friend_id"`
	Status    string    `gorm:"size:50" json:"status"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
