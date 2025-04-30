package team

import (
	"backend/domain/team"
	"backend/domain/team_member"
	"fmt"
)

type TeamUsecase struct {
	TeamRepository       team.TeamRepository
	TeamMemberRepository team_member.TeamMemberRepository
}

func NewTeamUsecase(repo team.TeamRepository, teamMemberRepo team_member.TeamMemberRepository) *TeamUsecase {
	return &TeamUsecase{repo, teamMemberRepo}
}

func (u *TeamUsecase) CreateTeam(input team.CreateInput, customerID int) error {
	team := &team.Team{
		Name:        input.Name,
		Description: input.Description,
	}
	err := u.TeamRepository.CreateTeam(team)
	if err != nil {
		return fmt.Errorf("failed to create team: %w", err)
	}
	teamMember := &team_member.TeamMember{
		TeamID:     team.TeamID,
		CustomerID: customerID,
		Role:       "owner",
	}
	if err := u.TeamMemberRepository.AddTeamMember(teamMember); err != nil {
		return fmt.Errorf("failed to add team member: %w", err)
	}
	return nil
}

func (u *TeamUsecase) GetTeam(id int) (*team.Team, error) {
	return u.TeamRepository.GetTeam(id)
}

func (u *TeamUsecase) UpdateTeam(input team.UpdateInput) error {
	team := &team.Team{
		Name:        input.Name,
		Description: input.Description,
	}
	if err := u.TeamRepository.UpdateTeam(input.ID, team); err != nil {
		return fmt.Errorf("failed to update team: %w", err)
	}
	return nil
}

func (u *TeamUsecase) DeleteTeam(id int) error {
	if err := u.TeamRepository.DeleteTeam(id); err != nil {
		return fmt.Errorf("failed to delete team: %w", err)
	}
	return nil
}
func (u *TeamUsecase) GetTeamsByCustomerID(customerID int) ([]*team.Team, error) {
	return u.TeamRepository.GetTeamsByCustomerID(customerID)
}
