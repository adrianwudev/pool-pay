package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"pool-pay/internal/domain"
	"pool-pay/internal/service"
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
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	router := c.FullPath()
	log.Printf("received request to add user for router %s. body: %s\n", router, userJson)

	userService := service.GetUserService(h.db)
	err = userService.AddUser(user.Username, user.Email, user.Password)
	if err != nil {
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.NewSuccessResponse("user added successfully", nil)
	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	var user *domain.User

	email := c.Query("email")

	userService := service.GetUserService(h.db)
	user, err := userService.GetByEmail(email)
	if err != nil {
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.NewSuccessResponse("get user successfully", user)
	c.JSON(http.StatusOK, response)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userService := service.GetUserService(h.db)
	token, err := userService.Login(request.Email, request.Password)
	if err != nil {
		log.Println(err)
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, util.NewSuccessResponse("login successfully", map[string]interface{}{"token": token}))
}
