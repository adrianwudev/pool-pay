package domain

import (
	"log"
	"pool-pay/auth"
	"time"

	redis_client "pool-pay/redis"

	"github.com/dgrijalva/jwt-go"
	"github.com/redis/go-redis/v9"
)

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
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

func (s *UserService) Login(email, password string, jwtToken *jwt.Token) (token string, err error) {
	client := redis_client.Client

	// Check token
	tokenSignature, err := client.Get(redis_client.Ctx, jwtToken.Signature).Result()
	if err != nil {
		if err != redis.Nil {
			return tokenSignature, nil
		} else {
			return "", err
		}
	}

	// Login
	isLogin, err := s.UserRepo.Login(email, password)
	if err != nil {
		return "", err
	}
	if isLogin {
		return "login successful", nil
	}

	// Generate token
	validToken, err := auth.GenerateJWT(email)
	if err != nil {
		log.Println(err)
	}

	// Write token into Redis
	err = client.Set(redis_client.Ctx, validToken, 1, time.Second*60).Err()
	if err != nil {
		return "", err
	}

	return validToken, nil
}

// func Authorize(UserId int64) (token string) {
// 	token = "token"
// 	return token
// }

// func GetUsers() []User{

// }
