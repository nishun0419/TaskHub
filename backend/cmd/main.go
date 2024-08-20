package main

import (
    "backend/controllers/auth"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    public := router.Group("/api")

    public.POST("/register", auth.Register)

    router.Run(":8080")
}