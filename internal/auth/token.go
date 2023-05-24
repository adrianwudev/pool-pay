package auth

import (
	"log"
	"pool-pay/internal/domain"
)

func IsMemExisted(userRepo domain.UserRepo, user domain.User) bool {

	foundUser, err := userRepo.GetByUsernameAndPassword("adrain", "secret")

	if err != nil {
		log.Println(err)
	}

	return foundUser != nil
}
