package domain

import (
	"log"

	"pool-pay/internal/auth"
	redis_client "pool-pay/redis"
)

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"unique;not null" json:"email"`
}

type UserRepo interface {
	AddUser(username, email, password string) error
	GetByEmail(email string) (*User, error)
	CheckIfExists(user User) (bool, error)
	Login(email, password string) (isLogin bool, err error)
}

type UserService struct {
	UserRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) AddUser(username, email, password string) error {
	return s.UserRepo.AddUser(username, email, password)
}

func (s *UserService) GetByEmail(email string) (*User, error) {
	return s.UserRepo.GetByEmail(email)
}

func (s *UserService) Login(email, password string) (token string, err error) {
	// Login
	isLogin, err := s.UserRepo.Login(email, password)
	if err != nil {
		log.Println("login failed")
		return "", err
	}
	if !isLogin {
		return "wrong email or password", nil
	}

	// Generate token
	validToken, err := auth.GenerateJWT(email)
	if err != nil {
		log.Println("generate token failed")
		log.Println(err)
	}

	// Write token into Redis
	err = redis_client.SetIntoRedis(validToken, email)
	if err != nil {
		log.Println("set token failed")
		return "", err
	}

	return validToken, nil
}
