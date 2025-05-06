package todo

import (
	"testing"
	"time"

	"backend/domain/todo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) Create(todo *todo.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockTodoRepository) GetByID(id int) (*todo.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*todo.Todo), args.Error(1)
}

func (m *MockTodoRepository) GetTeamTodos(teamID int) ([]*todo.Todo, error) {
	args := m.Called(teamID)
	return args.Get(0).([]*todo.Todo), args.Error(1)
}

func (m *MockTodoRepository) Update(todoID int, todo *todo.Todo) error {
	args := m.Called(todoID, todo)
	return args.Error(0)
}

func (m *MockTodoRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTodoRepository) ChangeStatus(todoID int, completed bool) error {
	args := m.Called(todoID, completed)
	return args.Error(0)
}

func TestCreateTodo(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	repo.On("Create", mock.Anything).Return(nil)

	input := &todo.Todo{
		Title:       "Test Todo",
		Description: "Test Description",
		TeamID:      1,
		CustomerID:  1,
		Completed:   false,
	}

	err := usecase.CreateTodo(input)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	todo := &todo.Todo{
		TodoID:      1,
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		TeamID:      1,
		CustomerID:  1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	repo.On("GetByID", 1).Return(todo, nil)

	todo, err := usecase.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, todo, todo)
}

func TestGetTodosByTeamID(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	repo.On("GetTeamTodos", 1).Return([]*todo.Todo{
		{
			TodoID:      1,
			Title:       "Test Todo 1",
			Description: "Test Description 1",
			TeamID:      1,
			CustomerID:  1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil)

	todos, err := usecase.GetTodosByTeamID(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(todos))

	repo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	input := &todo.Todo{
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		CustomerID:  1,
		TeamID:      1,
		Completed:   false,
	}

	repo.On("GetByID", 1).Return(input, nil)
	repo.On("Update", 1, input).Return(nil)

	err := usecase.Update(1, input)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestUpdateUnauthorized(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	wrongTodo := &todo.Todo{
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		CustomerID:  2,
		TeamID:      1,
		Completed:   false,
	}
	input := &todo.Todo{
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		CustomerID:  1,
		TeamID:      1,
		Completed:   false,
	}

	repo.On("GetByID", 1).Return(wrongTodo, nil)

	err := usecase.Update(1, input)
	assert.Error(t, err)

	repo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	repo.On("Delete", 1).Return(nil)
	repo.On("GetByID", 1).Return(&todo.Todo{
		TodoID:      1,
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		TeamID:      1,
		CustomerID:  1,
	}, nil)

	err := usecase.Delete(1, 1)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestDeleteUnauthorized(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	repo.On("GetByID", 1).Return(&todo.Todo{
		TodoID:      1,
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		TeamID:      1,
		CustomerID:  2,
	}, nil)
	err := usecase.Delete(1, 1)
	assert.Error(t, err)

	repo.AssertExpectations(t)
}

func TestChangeStatus(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	repo.On("ChangeStatus", 1, true).Return(nil)
	repo.On("GetByID", 1).Return(&todo.Todo{
		TodoID:      1,
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		TeamID:      1,
		CustomerID:  1,
	}, nil)
	err := usecase.ChangeStatus(1, 1, true)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestChangeStatusUnauthorized(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	repo.On("GetByID", 1).Return(&todo.Todo{
		TodoID:      1,
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		TeamID:      1,
		CustomerID:  2,
	}, nil)

	err := usecase.ChangeStatus(1, 1, true)
	assert.Error(t, err)

	repo.AssertExpectations(t)
}
