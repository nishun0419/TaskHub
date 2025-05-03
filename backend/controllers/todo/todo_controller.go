package todo

import (
	domain "backend/domain/todo"
	usecase "backend/usecase/todo"
	utils "backend/utils"

	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoUsecase *usecase.TodoUsecase
}

func NewTodoController(todoUsecase *usecase.TodoUsecase) *TodoController {
	return &TodoController{todoUsecase: todoUsecase}
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
	var input domain.Todo

	customerID := c.GetInt("customer_id")
	input.CustomerID = customerID

	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		utils.ErrorResponse(c, "Invalid team ID")
		return
	}
	input.TeamID = teamID

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	err = tc.todoUsecase.CreateTodo(&input)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "Todo created successfully", input)
}

func (tc *TodoController) GetTodo(c *gin.Context) {
	todoID := c.GetInt("todo_id")
	todo, err := tc.todoUsecase.GetByID(todoID)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "Todo retrieved successfully", todo)
}

func (tc *TodoController) GetTodosByTeamID(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		utils.ErrorResponse(c, "Invalid team ID")
		return
	}
	todos, err := tc.todoUsecase.GetTodosByTeamID(teamID)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "Todos retrieved successfully", todos)
}

func (tc *TodoController) UpdateTodo(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		utils.ErrorResponse(c, "Invalid team ID")
		return
	}
	var input domain.TodoUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	if err := tc.todoUsecase.Update(todoID, &input); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "Todo updated successfully", nil)
}

func (tc *TodoController) DeleteTodo(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		utils.ErrorResponse(c, "Invalid team ID")
		return
	}
	if err := tc.todoUsecase.Delete(todoID); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "Todo deleted successfully", nil)
}
