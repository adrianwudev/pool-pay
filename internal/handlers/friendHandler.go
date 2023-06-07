package handlers

import (
	"log"
	"net/http"
	"pool-pay/internal/auth"
	"pool-pay/internal/domain"
	"pool-pay/internal/service"
	"pool-pay/internal/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FriendshipHandler struct {
	db *gorm.DB
}

func NewFriendHandler(db *gorm.DB) *FriendshipHandler {
	return &FriendshipHandler{
		db: db,
	}
}

func (h *FriendshipHandler) AddFriend(c *gin.Context) {
	// get the user's info
	myEmail := getMyEmailFromToken(c)

	log.Printf("original email in middleware: %s\n", myEmail)
	// get friend info from the request body
	type AddFriendRequest struct {
		FriendEmail string `json:"friendEmail"`
	}
	var requestBody AddFriendRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	friendEmail := requestBody.FriendEmail

	// get userId by email
	userService := service.GetUserService(h.db)
	userId, err := getUserIdByEmail(userService, myEmail)
	if err != nil {
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	friendId, err := getUserIdByEmail(userService, friendEmail)
	if err != nil {
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	friendService := service.GetFriendshipService(h.db)
	// Add friend request
	err = friendService.AddFriendRequest(userId, friendId)
	if err != nil {
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// return response
	response := util.NewSuccessResponse("add friend request successfully", nil)
	c.JSON(http.StatusOK, response)

}

func getMyEmailFromToken(c *gin.Context) string {
	headerToken := c.Request.Header.Get("token")
	myEmail := auth.GetEmailFromJWT(headerToken)
	return myEmail
}

func getUserIdByEmail(userService *domain.UserService, email string) (int64, error) {
	user, err := userService.GetByEmail(email)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (h *FriendshipHandler) GetFriendRequests(c *gin.Context) {
	// get the user's info
	myEmail := getMyEmailFromToken(c)
	// get userId by email
	userService := service.GetUserService(h.db)
	userId, err := getUserIdByEmail(userService, myEmail)
	if err != nil {
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	friendService := service.GetFriendshipService(h.db)
	friendRequests, err := friendService.GetFriendRequests(userId)
	if err != nil {
		response := util.NewErrorResponse(err, util.GetApiError(err).Code)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.NewSuccessResponse("", friendRequests)
	c.JSON(http.StatusOK, response)
}

//how to return an error code rather than an error message
