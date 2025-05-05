package team_member

import (
	"backend/domain/customer"
	"backend/domain/team_member"
	"errors"
	"os"
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTeamMemberRepository struct {
	mock.Mock
}

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) FindByEmail(email string) (customer.Customer, error) {
	args := m.Called(email)
	return args.Get(0).(customer.Customer), args.Error(1)
}

func (m *MockCustomerRepository) RegisterCustomer(customer *customer.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockTeamMemberRepository) AddTeamMember(teamMember *team_member.TeamMember) error {
	args := m.Called(teamMember)
	return args.Error(0)
}

func (m *MockTeamMemberRepository) DeleteTeamMember(teamMemberDelInput *team_member.TeamMemberDelInput) error {
	args := m.Called(teamMemberDelInput)
	return args.Error(0)
}

func TestAddTeamMember(t *testing.T) {
	mockRepo := new(MockTeamMemberRepository)
	usecase := NewTeamMemberUsecase(mockRepo, nil)

	input := team_member.TeamMember{
		TeamID:     1,
		CustomerID: 1,
		Role:       "owner",
	}

	mockRepo.On("AddTeamMember", mock.AnythingOfType("*team_member.TeamMember")).Return(nil)

	err := usecase.AddTeamMember(&input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTeamMember(t *testing.T) {
	mockRepo := new(MockTeamMemberRepository)
	usecase := NewTeamMemberUsecase(mockRepo, nil)

	input := team_member.TeamMemberDelInput{
		TeamID:     1,
		CustomerID: 1,
	}

	mockRepo.On("DeleteTeamMember", mock.AnythingOfType("*team_member.TeamMemberDelInput")).Return(nil)

	err := usecase.DeleteTeamMember(&input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestInviteToken(t *testing.T) {
	mockRepo := new(MockTeamMemberRepository)
	customerRepo := new(MockCustomerRepository)
	usecase := NewTeamMemberUsecase(mockRepo, customerRepo)
	customerRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(customer.Customer{CustomerID: 1, Username: "test", Email: "test@example.com"}, nil)
	input := team_member.InviteTokenInput{
		TeamID: 1,
		Mail:   "test@example.com",
	}

	_, err := usecase.InviteToken(&input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestInviteToken_CustomerNotFound(t *testing.T) {
	mockRepo := new(MockTeamMemberRepository)
	customerRepo := new(MockCustomerRepository)
	usecase := NewTeamMemberUsecase(mockRepo, customerRepo)
	input := team_member.InviteTokenInput{
		TeamID: 1,
		Mail:   "test@example.com",
	}
	customerRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(customer.Customer{}, errors.New("customer not found"))

	_, err := usecase.InviteToken(&input)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestJoinTeam(t *testing.T) {
	mockRepo := new(MockTeamMemberRepository)
	usecase := NewTeamMemberUsecase(mockRepo, nil)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": 1,
		"team_id":     1,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	input := team_member.JoinTeamInput{
		Token: tokenString,
	}
	mockRepo.On("AddTeamMember", mock.AnythingOfType("*team_member.TeamMember")).Return(nil)

	err := usecase.JoinTeam(1, &input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestJoinTeam_InvalidToken(t *testing.T) {
	mockRepo := new(MockTeamMemberRepository)
	usecase := NewTeamMemberUsecase(mockRepo, nil)
	tokenString := "invalid_token"
	input := team_member.JoinTeamInput{
		Token: tokenString,
	}

	err := usecase.JoinTeam(1, &input)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestJoinTeam_InvalidCustomerID(t *testing.T) {
	mockRepo := new(MockTeamMemberRepository)
	usecase := NewTeamMemberUsecase(mockRepo, nil)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": 1,
		"team_id":     1,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	input := team_member.JoinTeamInput{
		Token: tokenString,
	}

	err := usecase.JoinTeam(2, &input)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
