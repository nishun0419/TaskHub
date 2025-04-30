package team

type TeamRepository interface {
	CreateTeam(team *Team) error
	GetTeam(id int) (*Team, error)
	UpdateTeam(id int, team *Team) error
	DeleteTeam(id int) error
	GetTeamsByCustomerID(customerID int) ([]*Team, error)
}
