package team_member

import (
	"backend/domain/customer"
	"backend/domain/team_member"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TeamMemberUsecase struct {
	TeamMemberRepository team_member.TeamMemberRepository
	CustomerRepository   customer.CustomerRepository
}

func NewTeamMemberUsecase(repo team_member.TeamMemberRepository, customerRepo customer.CustomerRepository) *TeamMemberUsecase {
	return &TeamMemberUsecase{repo, customerRepo}
}

func (u *TeamMemberUsecase) AddTeamMember(teamMember *team_member.TeamMember) error {
	return u.TeamMemberRepository.AddTeamMember(teamMember)
}

func (u *TeamMemberUsecase) DeleteTeamMember(teamMemberDelInput *team_member.TeamMemberDelInput) error {
	return u.TeamMemberRepository.DeleteTeamMember(teamMemberDelInput)
}

func (u *TeamMemberUsecase) InviteToken(inviteTokenInput *team_member.InviteTokenInput) (string, error) {
	customer, err := u.CustomerRepository.FindByEmail(inviteTokenInput.Mail)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": customer.CustomerID,
		"team_id":     inviteTokenInput.TeamID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (u *TeamMemberUsecase) JoinTeam(customerID int, joinTeamInput *team_member.JoinTeamInput) error {
	token, err := jwt.Parse(joinTeamInput.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}
	jwtTeamID := int(claims["team_id"].(float64))
	jwtCustomerID := int(claims["customer_id"].(float64))

	if customerID != jwtCustomerID {
		return errors.New("invalid token")
	}

	if err := u.TeamMemberRepository.AddTeamMember(&team_member.TeamMember{
		TeamID:     jwtTeamID,
		CustomerID: customerID,
		Role:       "member",
	}); err != nil {
		return err
	}

	return nil
}
