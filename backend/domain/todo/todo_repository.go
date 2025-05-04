package todo

type TodoRepository interface {
	Create(input *Todo) error
	GetByID(id int) (*Todo, error)
	GetTeamTodos(teamID int) ([]*Todo, error)
	Update(todoID int, input *Todo) error
	ChangeStatus(todoID int, completed bool) error
	Delete(id int) error
}
