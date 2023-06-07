package repository

import (
	"errors"
	"log"
	"pool-pay/internal/constants"
	"pool-pay/internal/domain"
	"pool-pay/internal/util"

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
	err := db.Conn.Select("id, username, email").Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Println(err)
		return nil, util.SetDefaultApiError(err)
	}

	return &user, nil
}

func (db *dbUserRepo) AddUser(username, email, password string) error {
	existingUser := domain.User{}
	err := db.Conn.Where("email = ?", email).First(&existingUser).Error
	if err == nil {
		log.Println(err)
		return util.SetApiError(constants.ERRORCODE_EMAILALREADYEXISTS)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		return util.SetDefaultApiError(err)
	}

	// Hash the password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Println(err)
		return util.SetDefaultApiError(err)
	}

	user := &domain.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = db.Conn.Create(user).Error
	if err != nil {
		log.Println(err)
		return util.SetDefaultApiError(err)
	}

	return nil
}

func (db *dbUserRepo) Login(email, password string) (isLogin bool, err error) {
	isLogin = false

	// Check email is repeated
	isEmailRepeated, err := isEmailRepeated(email, db)
	if err != nil {
		log.Println(err)
		return isLogin, util.SetDefaultApiError(err)
	}
	if isEmailRepeated {
		log.Println(errors.New("email is repeated"))
		return isLogin, util.SetApiError(constants.ERRORCODE_INVALIDEMAIL)
	}

	// Get stored hash
	var storedHash string
	err = db.Conn.Model(&domain.User{}).Select("password").Where("email = ?", email).Scan(&storedHash).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return isLogin, util.SetApiError(constants.ERRORCODE_INVALIDEMAIL)
		}
		return isLogin, util.SetDefaultApiError(err)
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
		return "", util.SetDefaultApiError(err)
	}

	return string(hashedPassword), util.SetDefaultApiError(err)
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
		return isRepeated, util.SetDefaultApiError(err)
	}
	if count > 1 {
		isRepeated = true
		log.Println(errors.New("repeated email exists"))
		return isRepeated, util.SetApiError(constants.ERRORCODE_INVALIDEMAIL)
	}

	return isRepeated, nil
}
