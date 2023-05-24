package main

import (
	"fmt"
	"log"
	"net/http"

	"pool-pay/config"
	"pool-pay/db"
	"pool-pay/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var cfg = config.GetConfig()
var myDb *gorm.DB

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", GreetingHandler)
	router.POST("/api/v1/user", func(c *gin.Context) {
		userHandler := handlers.NewUserHandler(myDb)
		userHandler.AddUserHandler(c)
	})

	return router
}

func main() {
	fmt.Println("Welcome to paypool")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Printf("Host: %s, Port: %s, Password: %s, User: %s, DBName: %s, SSLMode: %s\n",
		cfg.Host, cfg.Port, cfg.Password, cfg.User, cfg.DBName, cfg.SSLMode)

	myDb = db.ConnectDb(cfg)

	// routers
	router := setupRouter()
	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started on port 8080")

}

func GreetingHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}
