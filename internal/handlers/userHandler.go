package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"pool-pay/internal/domain"
	"pool-pay/internal/repository"

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

func (h *UserHandler) RegisterHandler(c *gin.Context) {
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
	log.Printf("received request to add user for router %s. body: %s\n", router, userJson)

	userService := getUserService(h)
	err = userService.AddUser(user.Username, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user added successfully"})
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	var user *domain.User

	email := c.Query("email")

	userService := getUserService(h)
	user, err := userService.GetByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userJson, err := json.Marshal(user)
	log.Println(userJson)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if isGotToken(c) {
		return
	}

	userService := getUserService(h)
	token, err := userService.Login(request.Email, request.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func isGotToken(c *gin.Context) bool {
	headerToken := c.Request.Header.Get("token")
	isEmpty := len(strings.TrimSpace(headerToken)) == 0
	if !isEmpty {
		c.JSON(http.StatusBadRequest, gin.H{"error": "already have token"})
		return true
	}
	return false
}

func getUserService(h *UserHandler) *domain.UserService {
	userRepo := repository.NewDbUserRepository(h.db)
	userService := domain.NewUserService(userRepo)
	return userService
}
