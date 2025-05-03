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

func (m *MockTodoRepository) Update(todoID int, input *todo.TodoUpdate) error {
	args := m.Called(todoID, input)
	return args.Error(0)
}

func (m *MockTodoRepository) Delete(id int) error {
	args := m.Called(id)
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

	input := &todo.TodoUpdate{
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		CustomerID:  1,
	}

	repo.On("Update", 1, input).Return(nil)

	err := usecase.Update(1, input)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	repo.On("Delete", 1).Return(nil)

	err := usecase.Delete(1)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}
