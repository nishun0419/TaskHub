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

func (m *MockTodoRepository) GetTeamTodos(teamID int) ([]*todo.Todo, error) {
	args := m.Called(teamID)
	return args.Get(0).([]*todo.Todo), args.Error(1)
}

func (m *MockTodoRepository) GetByID(id int) (*todo.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*todo.Todo), args.Error(1)
}

func (m *MockTodoRepository) Update(todo *todo.Todo) error {
	args := m.Called(todo)
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

	todo := &todo.Todo{
		Title:       "Test Todo",
		Description: "Test Description",
		TeamID:      1,
		CustomerID:  1,
	}

	err := usecase.CreateTodo(todo)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestGetTeamTodos(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	repo.On("GetTeamTodos", 1).Return([]*todo.Todo{
		{
			ID:          1,
			Title:       "Test Todo 1",
			Description: "Test Description 1",
			TeamID:      1,
			CustomerID:  1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil)

	todos, err := usecase.GetTeamTodos(1)
	assert.NoError(t, err)
	assert.Len(t, todos, 1)
	assert.Equal(t, "Test Todo 1", todos[0].Title)

	repo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	repo.On("GetByID", 1).Return(&todo.Todo{
		ID:          1,
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		TeamID:      1,
		CustomerID:  1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil)

	todo, err := usecase.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, todo.ID)
	assert.Equal(t, "Test Todo 1", todo.Title)
	assert.Equal(t, "Test Description 1", todo.Description)
	assert.Equal(t, 1, todo.TeamID)
	assert.Equal(t, 1, todo.CustomerID)

	repo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	repo := new(MockTodoRepository)
	usecase := NewTodoUsecase(repo)

	repo.On("Update", mock.Anything).Return(nil)

	todo := &todo.Todo{
		ID:          1,
		Title:       "Test Todo 1",
		Description: "Test Description 1",
		TeamID:      1,
		CustomerID:  1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := usecase.Update(todo)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}
