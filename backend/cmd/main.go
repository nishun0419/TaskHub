package main

import (
	"backend/controllers/customer"
	repository "backend/infrastructure/customer"
	"backend/pkg/db"
	usecase "backend/usecase/customer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.ConnectDataBase()

	customerRepository := repository.NewCustomerRepository(database)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)
	customerController := customer.NewCustomerController(customerUsecase)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	public := router.Group("/api")

	public.POST("/register", customerController.RegisterHandler)
	public.POST("/login", customerController.LoginHandler)

	router.Run(":8080")
}
