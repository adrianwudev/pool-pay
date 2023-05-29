package repository

import (
	"errors"
	"fmt"
	"pool-pay/internal/domain"

	"golang.org/x/crypto/bcrypt"
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

func (db *dbUserRepo) GetByEmail(email string) (*domain.User, error) {
	user := domain.User{}
	err := db.Conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *dbUserRepo) AddUser(username, email, password string) error {
	existingUser := domain.User{}
	err := db.Conn.Where("email = ?", email).First(&existingUser).Error
	if err == nil {
		return fmt.Errorf("email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Hash the password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = db.Conn.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *dbUserRepo) Login(email, password string) (isLogin bool, err error) {
	isLogin = false

	// Check email is repeated
	isEmailRepeated, err := isEmailRepeated(email, db)
	if err != nil {
		return isLogin, err
	}
	if isEmailRepeated {
		return isLogin, fmt.Errorf("email is repeated")
	}

	// Get stored hash
	var storedHash string
	err = db.Conn.Model(&domain.User{}).Select("password").Where("email = ?", email).Scan(&storedHash).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return isLogin, fmt.Errorf("invalid email")
		}
		return isLogin, err
	}

	// Check if the user exists in the DB
	isPasswordMatched := comparePasswordHash(password, storedHash)

	if isPasswordMatched {
		isLogin = true
		return isLogin, nil
	}

	return isLogin, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}

func comparePasswordHash(password, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	return err == nil
}

func isEmailRepeated(email string, db *dbUserRepo) (isRepeated bool, err error) {
	var count int64
	isRepeated = false
	err = db.Conn.Model(&domain.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return isRepeated, err
	}
	if count > 1 {
		isRepeated = true
		return isRepeated, fmt.Errorf("repeated email exists")
	}

	return isRepeated, nil
}
