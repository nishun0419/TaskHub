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
	todoID, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		utils.ErrorResponse(c, "Invalid todo ID")
		return
	}
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
		utils.ErrorResponse(c, "Invalid todo ID")
		return
	}

	// 既存のTODOを取得
	existingTodo, err := tc.todoUsecase.GetByID(todoID)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	var input domain.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	// 必要な情報を設定
	customerID := c.GetInt("customer_id")
	input.CustomerID = customerID
	input.TeamID = existingTodo.TeamID
	input.TodoID = todoID

	if err := tc.todoUsecase.Update(todoID, &input); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "Todo updated successfully", nil)
}

func (tc *TodoController) DeleteTodo(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		utils.ErrorResponse(c, "Invalid todo ID")
		return
	}
	if err := tc.todoUsecase.Delete(todoID); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "Todo deleted successfully", nil)
}

func (tc *TodoController) ChangeStatus(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		utils.ErrorResponse(c, "Invalid todo ID")
		return
	}

	var input domain.ChangeStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	customerID := c.GetInt("customer_id")
	if err := tc.todoUsecase.ChangeStatus(todoID, customerID, input.Completed); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.SuccessResponse(c, "Todo status updated successfully", nil)
}
