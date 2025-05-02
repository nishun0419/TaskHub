package todo

import (
	"backend/domain/todo"
)

type TodoUsecase struct {
	todoRepository todo.TodoRepository
}

func NewTodoUsecase(todoRepository todo.TodoRepository) *TodoUsecase {
	return &TodoUsecase{todoRepository: todoRepository}
}

func (u *TodoUsecase) CreateTodo(todo *todo.Todo) error {
	return u.todoRepository.Create(todo)
}

func (u *TodoUsecase) GetTeamTodos(teamID int) ([]*todo.Todo, error) {
	return u.todoRepository.GetTeamTodos(teamID)
}

func (u *TodoUsecase) GetByID(id int) (*todo.Todo, error) {
	return u.todoRepository.GetByID(id)
}

func (u *TodoUsecase) Update(todo *todo.Todo) error {
	return u.todoRepository.Update(todo)
}

func (u *TodoUsecase) Delete(id int) error {
	return u.todoRepository.Delete(id)
}
