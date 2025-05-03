package todo

import (
	"backend/domain/todo"

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

func (r *TodoRepository) Update(id int, input *todo.TodoUpdate) error {
	return r.db.Model(&todo.Todo{}).Where("todo_id = ?", id).Updates(input).Error
}

func (r *TodoRepository) Delete(id int) error {
	return r.db.Delete(&todo.Todo{}, "todo_id = ?", id).Error
}
