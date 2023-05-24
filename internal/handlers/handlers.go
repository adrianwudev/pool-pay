package handlers

import (
	"encoding/json"
	"log"
	"net/http"

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

	userRepo := repository.NewDbUserRepository(h.db)
	userService := domain.NewUserService(userRepo)
	err = userService.AddUser(user.Username, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User added successfully"})
}
