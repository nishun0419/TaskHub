package todo

type TodoRepository interface {
	Create(input *Todo) error
	GetByID(id int) (*Todo, error)
	GetTeamTodos(teamID int) ([]*Todo, error)
	Update(todoID int, input *TodoUpdate) error
	Delete(id int) error
}
