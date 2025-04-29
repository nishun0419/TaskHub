package teams

type TeamRepository interface {
	CreateTeam(team *Team) error
	GetTeam(id int) (*Team, error)
	UpdateTeam(id int, team *Team) error
	DeleteTeam(id int) error
	AddMember(teamID int, customerID int) error
	RemoveMember(teamID int, customerID int) error
	GetAllTeams() ([]*Team, error)
}
