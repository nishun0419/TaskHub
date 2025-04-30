package main

import (
	"backend/controllers/customer"
	teamsController "backend/controllers/teams"
	repository "backend/infrastructure/customer"
	teamsRepository "backend/infrastructure/teams"
	"backend/pkg/db"
	usecase "backend/usecase/customer"
	teamsUsecase "backend/usecase/teams"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.ConnectDataBase()

	customerRepository := repository.NewCustomerRepository(database)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)
	customerController := customer.NewCustomerController(customerUsecase)

	teamRepository := teamsRepository.NewTeamRepository(database)
	teamUsecase := teamsUsecase.NewTeamUsecase(teamRepository)
	teamController := teamsController.NewTeamController(teamUsecase)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	public := router.Group("/api")

	public.POST("/register", customerController.RegisterHandler)
	public.POST("/login", customerController.LoginHandler)

	teams := public.Group("/teams")
	teams.POST("", teamController.CreateTeam)
	teams.GET("/:id", teamController.GetTeam)
	teams.PUT("/:id", teamController.UpdateTeam)
	teams.DELETE("/:id", teamController.DeleteTeam)

	router.Run(":8080")
}
