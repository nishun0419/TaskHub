package team_member

type TeamMemberRepository interface {
	AddTeamMember(teamMember *TeamMember) error
	DeleteTeamMember(teamMemberDelInput *TeamMemberDelInput) error
	IsTeamMember(teamID int, customerID int) (bool, error)
}
