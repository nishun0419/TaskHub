package team

import (
	"backend/domain/customer"
	"backend/domain/team"
	"backend/domain/team_member"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TeamUsecase struct {
	TeamRepository       team.TeamRepository
	TeamMemberRepository team_member.TeamMemberRepository
	CustomerRepository   customer.CustomerRepository
}

func NewTeamUsecase(repo team.TeamRepository, teamMemberRepo team_member.TeamMemberRepository, customerRepo customer.CustomerRepository) *TeamUsecase {
	return &TeamUsecase{repo, teamMemberRepo, customerRepo}
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
func (u *TeamUsecase) GetTeamsByCustomerID(customerID int) ([]*team.TeamWithRole, error) {
	return u.TeamRepository.GetTeamsByCustomerID(customerID)
}

func (u *TeamUsecase) GenerateInviteToken(input team.InviteTokenInput) (string, error) {
	customer, err := u.CustomerRepository.FindByEmail(input.Mail)
	if err != nil {
		return "", fmt.Errorf("failed to find customer by email: %w", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": customer.CustomerID,
		"team_id":     input.TeamID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (u *TeamUsecase) JoinTeam(customerID int, input team.JoinTeamInput) error {
	token, err := jwt.Parse(input.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid token")
	}
	addCustomerID := int(claims["customer_id"].(float64))
	if addCustomerID != customerID {
		return fmt.Errorf("invalid customer ID")
	}
	teamID := int(claims["team_id"].(float64))
	team, err := u.TeamRepository.GetTeam(teamID)
	if err != nil {
		return fmt.Errorf("failed to get team: %w", err)
	}
	if team == nil {
		return fmt.Errorf("team not found")
	}
	teamMember := &team_member.TeamMember{
		TeamID:     teamID,
		CustomerID: addCustomerID,
		Role:       "member",
	}
	if err := u.TeamMemberRepository.AddTeamMember(teamMember); err != nil {
		return fmt.Errorf("failed to add team member: %w", err)
	}
	return nil
}
