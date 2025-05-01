package main

import (
	"backend/controllers/customer"
	teamsController "backend/controllers/team"
	repository "backend/infrastructure/customer"
	teamsRepository "backend/infrastructure/team"
	teamMemberRepository "backend/infrastructure/team_member"
	"backend/middleware"
	"backend/pkg/db"
	usecase "backend/usecase/customer"
	teamUsecase "backend/usecase/team"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.ConnectDataBase()

	customerRepository := repository.NewCustomerRepository(database)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)
	customerController := customer.NewCustomerController(customerUsecase)

	teamRepository := teamsRepository.NewTeamRepository(database)
	teamMemberRepository := teamMemberRepository.NewTeamMemberRepository(database)
	teamUsecase := teamUsecase.NewTeamUsecase(teamRepository, teamMemberRepository)
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

	// 認証が必要なエンドポイント
	teams := public.Group("/team")
	teams.Use(middleware.AuthMiddleware())
	teams.POST("", teamController.CreateTeam)
	teams.GET("/:id", teamController.GetTeam)
	teams.PUT("/:id", teamController.UpdateTeam)
	teams.DELETE("/:id", teamController.DeleteTeam)
	teams.GET("", teamController.GetTeamsByCustomerID)
	router.Run(":8080")
}
