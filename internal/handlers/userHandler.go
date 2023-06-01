package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"pool-pay/internal/domain"
	"pool-pay/internal/util"

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
		response := util.NewErrorResponse(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	router := c.FullPath()
	log.Printf("received request to add user for router %s. body: %s\n", router, userJson)

	userService := util.GetUserService(h.db)
	err = userService.AddUser(user.Username, user.Email, user.Password)
	if err != nil {
		response := util.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.NewSuccessResponse("user added successfully", nil)
	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	var user *domain.User

	email := c.Query("email")

	userService := util.GetUserService(h.db)
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
	if isTokenExists(c) {
		return
	}

	userService := util.GetUserService(h.db)
	token, err := userService.Login(request.Email, request.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func isTokenExists(c *gin.Context) bool {
	headerToken := c.Request.Header.Get("token")
	isEmpty := len(strings.TrimSpace(headerToken)) == 0
	if !isEmpty {
		c.JSON(http.StatusBadRequest, gin.H{"error": "already have token"})
		return true
	}
	return false
}
