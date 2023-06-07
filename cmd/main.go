package main

import (
	"fmt"
	"log"
	"net/http"

	"pool-pay/config"
	"pool-pay/db"
	"pool-pay/internal/handlers"
	"pool-pay/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var cfg = config.GetConfig()
var myDb *gorm.DB

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	// middleware group
	authenticatedGroup := router.Group("/auth", middleware.AuthMiddleware)

	// handlers
	userHandler := handlers.NewUserHandler(myDb)
	friendHandler := handlers.NewFriendHandler(myDb)

	// routes
	router.GET("/", GreetingHandler)
	router.POST("/api/v1/user/register", func(c *gin.Context) {
		userHandler.RegisterHandler(c)
	})
	router.POST("/api/v1/user/login", userHandler.Login)
	authenticatedGroup.GET("/api/v1/user", userHandler.GetUserByEmail)
	authenticatedGroup.POST("/api/v1/friend", friendHandler.AddFriend)
	authenticatedGroup.GET("/api/v1/friend/requests", friendHandler.GetFriendRequests)

	return router
}

func main() {
	fmt.Println("welcome to paypool")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Printf("Host: %s, Port: %s, Password: %s, User: %s, DBName: %s, SSLMode: %s\n",
		cfg.Host, cfg.Port, cfg.Password, cfg.User, cfg.DBName, cfg.SSLMode)

	myDb = db.ConnectDb(cfg)

	// routers
	router := setupRouter()
	err := router.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server started on port 8081")

}

func GreetingHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}
