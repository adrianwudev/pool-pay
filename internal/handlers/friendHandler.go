package handlers

import (
	"log"
	"net/http"
	"pool-pay/internal/auth"
	"pool-pay/internal/domain"
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
	headerToken := c.Request.Header.Get("token")
	myEmail := auth.GetEmailFromJWT(headerToken)

	log.Printf("original email in middleware: %s\n", myEmail)
	// get friend info from the request body
	type AddFriendRequest struct {
		FriendEmail string `json:"friendEmail"`
	}
	var requestBody AddFriendRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		response := util.NewErrorResponse(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	friendEmail := requestBody.FriendEmail

	// get userId by email
	userService := util.GetUserService(h.db)
	userId, err := getUserIdByEmail(userService, myEmail)
	if err != nil {
		response := util.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	friendId, err := getUserIdByEmail(userService, friendEmail)
	if err != nil {
		response := util.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	friendService := util.GetFriendshipService(h.db)
	// Add friend request
	err = friendService.AddFriendRequest(userId, friendId)
	if err != nil {
		response := util.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// return response
	response := util.NewSuccessResponse("add friend request successfully", nil)
	c.JSON(http.StatusOK, response)

}

func getUserIdByEmail(userService *domain.UserService, email string) (int64, error) {
	user, err := userService.GetByEmail(email)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
