package repository

import (
	"errors"
	"fmt"
	"pool-pay/internal/domain"

	"gorm.io/gorm"
)

type dbUserRepo struct {
	Conn *gorm.DB
}

func NewDbUserRepository(conn *gorm.DB) domain.UserRepo {
	return &dbUserRepo{conn}
}

func (db *dbUserRepo) CheckIfExists(user domain.User) (bool, error) {
	panic("unimplemented")
}

func (db *dbUserRepo) GetById(userId int64) (*domain.User, error) {
	return &domain.User{Id: userId, Username: "adrian", Email: "adrian@gmail.com"}, nil
}

func (db *dbUserRepo) AddUser(username, email, password string) error {
	existingUser := &domain.User{}
	err := db.Conn.Where("email = ?", email).First(existingUser).Error
	if err == nil {
		return fmt.Errorf("email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	user := &domain.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err = db.Conn.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}
