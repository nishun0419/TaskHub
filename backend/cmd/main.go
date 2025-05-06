package main

import (
	"backend/controllers/customer"
	teamsController "backend/controllers/team"
	todoController "backend/controllers/todo"
	repository "backend/infrastructure/customer"
	teamsRepository "backend/infrastructure/team"
	teamMemberRepository "backend/infrastructure/team_member"
	todoRepository "backend/infrastructure/todo"
	"backend/middleware"
	"backend/pkg/db"
	usecase "backend/usecase/customer"
	teamUsecase "backend/usecase/team"
	todoUsecase "backend/usecase/todo"
	"os"

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
	teamUsecase := teamUsecase.NewTeamUsecase(teamRepository, teamMemberRepository, customerRepository)
	teamController := teamsController.NewTeamController(teamUsecase)

	todoRepository := todoRepository.NewTodoRepository(database)
	todoUsecase := todoUsecase.NewTodoUsecase(todoRepository)
	todoController := todoController.NewTodoController(todoUsecase)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	frontendOrigin := os.Getenv("FRONTEND_ORIGIN")
	if frontendOrigin == "" {
		frontendOrigin = "http://localhost:3000"
	}
	config.AllowOrigins = []string{frontendOrigin}
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
	teams.POST("/:team_id/invite", teamController.GenerateInviteToken)
	teams.POST("/join", teamController.JoinTeam)

	todo := public.Group("/todo")
	todo.Use(middleware.AuthMiddleware())
	todo.POST("/team/:team_id", todoController.CreateTodo)
	todo.GET("/:todo_id", todoController.GetTodo)
	todo.GET("/team/:team_id", todoController.GetTodosByTeamID)
	todo.PUT("/:todo_id", todoController.UpdateTodo)
	todo.PUT("/:todo_id/status", todoController.ChangeStatus)
	todo.DELETE("/:todo_id", todoController.DeleteTodo)

	router.Run(":8080")
}
