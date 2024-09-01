package main

import (
	"backend/controllers/customer"
	"backend/pkg/db"
	service "backend/service/customer"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.ConnectDataBase()

	customerService := service.NewCustomerService(database)
	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", customer.RegisterHandler(customerService))

	router.Run(":8080")
}
