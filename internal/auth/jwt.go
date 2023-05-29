package auth

import (
	"fmt"
	"log"
	"pool-pay/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(email string) (string, error) {
	log.Println("generate JWT")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	signedKey := []byte(config.MySigningKey)
	tokenString, err := token.SignedString(signedKey)

	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println("JWT generated.")
	return tokenString, nil
}

func GetEmailFromJWT(validToken string) (email string) {
	mySigningKey := []byte(config.MySigningKey)

	// parse JWT
	token, err := jwt.Parse(validToken, func(token *jwt.Token) (interface{}, error) {
		// verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		log.Println(err)
		return
	}

	// retrieve original email
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		email = claims["email"].(string)
		log.Println("original email:", email)
	} else {
		log.Println("invalid token claims")
	}
	return email
}
