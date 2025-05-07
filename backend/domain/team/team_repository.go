package team

type TeamRepository interface {
	CreateTeam(team *Team) error
	GetTeam(id int, customerID int) (*TeamWithRole, error)
	UpdateTeam(id int, team *Team) error
	DeleteTeam(id int) error
	GetTeamsByCustomerID(customerID int) ([]*TeamWithRole, error)
}
