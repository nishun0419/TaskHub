package teams

import (
	"backend/domain/teams"
	"fmt"
)

type TeamUsecase struct {
	TeamRepository teams.TeamRepository
}

func NewTeamUsecase(repo teams.TeamRepository) *TeamUsecase {
	return &TeamUsecase{repo}
}

func (u *TeamUsecase) CreateTeam(input teams.CreateInput) error {
	team := &teams.Team{
		Name:        input.Name,
		Description: input.Description,
	}
	if err := u.TeamRepository.CreateTeam(team); err != nil {
		return fmt.Errorf("failed to create team: %w", err)
	}
	return nil
}

func (u *TeamUsecase) GetTeam(id teams.TeamID) (*teams.Team, error) {
	return u.TeamRepository.GetTeam(id.ID)
}

func (u *TeamUsecase) UpdateTeam(input teams.UpdateInput) error {
	team := &teams.Team{
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

func (u *TeamUsecase) AddMember(teamID int, customerID int) error {
	if err := u.TeamRepository.AddMember(teamID, customerID); err != nil {
		return fmt.Errorf("failed to add member to team: %w", err)
	}
	return nil
}

func (u *TeamUsecase) RemoveMember(teamID int, customerID int) error {
	if err := u.TeamRepository.RemoveMember(teamID, customerID); err != nil {
		return fmt.Errorf("failed to remove member from team: %w", err)
	}
	return nil
}

func (u *TeamUsecase) GetTeamsByCustomerID(customerID int) ([]*teams.Team, error) {
	return u.TeamRepository.GetTeamsByCustomerID(customerID)
}
