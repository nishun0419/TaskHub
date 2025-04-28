package main

import (
	"backend/controllers/customer"
	"backend/pkg/db"
	service "backend/service/customer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.ConnectDataBase()

	customerServiceInterface := service.NewCustomerService(database)
	router := gin.Default()

	// CORSミドルウェアの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	public := router.Group("/api")

	public.POST("/register", customer.RegisterHandler(customerServiceInterface))
	public.POST("/login", customer.LoginHandler(customerServiceInterface))

	router.Run(":8080")
}
