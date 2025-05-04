package todo

import (
	"backend/domain/todo"
	"time"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(input *todo.Todo) error {
	return r.db.Create(input).Error
}

func (r *TodoRepository) GetByID(id int) (*todo.Todo, error) {
	var todo todo.Todo
	if err := r.db.Where("todo_id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) GetTeamTodos(teamID int) ([]*todo.Todo, error) {
	var todos []*todo.Todo
	if err := r.db.Where("team_id = ?", teamID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) Update(id int, input *todo.Todo) error {
	updates := map[string]interface{}{
		"title":       input.Title,
		"description": input.Description,
		"completed":   input.Completed,
		"team_id":     input.TeamID,
		"updated_at":  time.Now(),
	}
	return r.db.Model(&todo.Todo{}).Where("todo_id = ?", id).Updates(updates).Error
}

func (r *TodoRepository) Delete(id int) error {
	return r.db.Delete(&todo.Todo{}, "todo_id = ?", id).Error
}

func (r *TodoRepository) ChangeStatus(todoID int, completed bool) error {
	updates := map[string]interface{}{
		"completed":  completed,
		"updated_at": time.Now(),
	}
	return r.db.Model(&todo.Todo{}).Where("todo_id = ?", todoID).Updates(updates).Error
}
