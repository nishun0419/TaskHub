package todo

type TodoRepository interface {
	Create(todo *Todo) error
	GetTeamTodos(teamID int) ([]*Todo, error)
	GetByID(id int) (*Todo, error)
	Update(todo *Todo) error
	Delete(id int) error
}
