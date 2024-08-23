package main

import (
	"backend/controllers/customer"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", customer.Register)

	router.Run(":8080")
}
