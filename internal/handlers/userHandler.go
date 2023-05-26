package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"pool-pay/internal/domain"
	"pool-pay/internal/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (h *UserHandler) AddUserHandler(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	router := c.FullPath()
	log.Printf("Received request to add user for router %s. Body: %s\n", router, userJson)

	userService := getUserService(h)
	err = userService.AddUser(user.Username, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User added successfully"})
}

func (h *UserHandler) GetUserHandler(c *gin.Context) {
	var user *domain.User

	email := c.Query("email")

	userService := getUserService(h)
	user, err := userService.GetByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userJson, err := json.Marshal(user)
	fmt.Println(userJson)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": string(userJson)})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}
	// if token equals to token from redis, then login successfully
	//TODO
	var jwtToken *jwt.Token
	var mySigningKey string
	var err error
	if c.Request.Header["token"] != nil {
		jwtToken, err := jwt.Parse(r.Header["token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	userService := getUserService(h)
	token, err := userService.Login(request.Email, request.Password, jwtToken)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func getUserService(h *UserHandler) *domain.UserService {
	userRepo := repository.NewDbUserRepository(h.db)
	userService := domain.NewUserService(userRepo)
	return userService
}
