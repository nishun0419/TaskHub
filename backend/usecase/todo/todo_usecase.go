package todo

import (
	"backend/domain/todo"
	"fmt"
)

type TodoUsecase struct {
	todoRepository todo.TodoRepository
}

func NewTodoUsecase(todoRepository todo.TodoRepository) *TodoUsecase {
	return &TodoUsecase{todoRepository: todoRepository}
}

func (u *TodoUsecase) CreateTodo(input *todo.Todo) error {
	return u.todoRepository.Create(input)
}

func (u *TodoUsecase) GetByID(id int) (*todo.Todo, error) {
	return u.todoRepository.GetByID(id)
}

func (u *TodoUsecase) GetTodosByTeamID(teamID int) ([]*todo.Todo, error) {
	return u.todoRepository.GetTeamTodos(teamID)
}

func (u *TodoUsecase) Update(todoID int, input *todo.Todo) error {
	existingTodo, err := u.todoRepository.GetByID(todoID)
	if err != nil {
		return err
	}
	if existingTodo.CustomerID != input.CustomerID {
		return fmt.Errorf("unauthorized")
	}
	return u.todoRepository.Update(todoID, input)
}

func (u *TodoUsecase) Delete(id int) error {
	return u.todoRepository.Delete(id)
}

func (u *TodoUsecase) ChangeStatus(todoID int, customerID int, completed bool) error {
	existingTodo, err := u.todoRepository.GetByID(todoID)
	if err != nil {
		return err
	}
	if existingTodo.CustomerID != customerID {
		return fmt.Errorf("unauthorized")
	}
	return u.todoRepository.ChangeStatus(todoID, completed)
}
