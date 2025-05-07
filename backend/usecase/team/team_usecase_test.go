package team

import (
	"backend/domain/customer"
	"backend/domain/team"
	"backend/domain/team_member"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTeamRepository struct {
	mock.Mock
}

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

func (m *MockTeamRepository) CreateTeam(t *team.Team) error {
	args := m.Called(t)
	return args.Error(0)
}

func (m *MockTeamRepository) GetTeam(teamID int, customerID int) (*team.TeamWithRole, error) {
	args := m.Called(teamID, customerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*team.TeamWithRole), args.Error(1)
}

func (m *MockTeamRepository) UpdateTeam(id int, team *team.Team) error {
	args := m.Called(id, team)
	return args.Error(0)
}

func (m *MockTeamRepository) DeleteTeam(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTeamRepository) GetTeamsByCustomerID(customerID int) ([]*team.TeamWithRole, error) {
	args := m.Called(customerID)
	return args.Get(0).([]*team.TeamWithRole), args.Error(1)
}

func (m *MockTeamRepository) GenerateInviteToken(token *team.InviteTokenInput) (*team.Team, error) {
	args := m.Called(token)
	return args.Get(0).(*team.Team), args.Error(1)
}

func (m *MockTeamRepository) JoinTeam(teamID int, customerID int) error {
	args := m.Called(teamID, customerID)
	return args.Error(0)
}

func (m *MockTeamMemberRepository) AddTeamMember(teamMember *team_member.TeamMember) error {
	args := m.Called(teamMember)
	return args.Error(0)
}

func (m *MockTeamMemberRepository) GetTeamMember(customerID int, teamID int) (*team_member.TeamMember, error) {
	args := m.Called(customerID, teamID)
	return args.Get(0).(*team_member.TeamMember), args.Error(1)
}

func (m *MockTeamMemberRepository) DeleteTeamMember(teamMemberDelInput *team_member.TeamMemberDelInput) error {
	args := m.Called(teamMemberDelInput)
	return args.Error(0)
}
func TestCreateTeam(t *testing.T) {
	teamRepo := new(MockTeamRepository)
	teamMemberRepo := new(MockTeamMemberRepository)
	usecase := NewTeamUsecase(teamRepo, teamMemberRepo, nil)

	input := team.CreateInput{
		Name:        "Test Team",
		Description: "Test Description",
	}

	teamRepo.On("CreateTeam", mock.AnythingOfType("*team.Team")).Return(nil)
	teamMemberRepo.On("AddTeamMember", mock.AnythingOfType("*team_member.TeamMember")).Return(nil)

	err := usecase.CreateTeam(input, 1)

	assert.NoError(t, err)
	teamRepo.AssertExpectations(t)
	teamMemberRepo.AssertExpectations(t)
}

func TestGetTeam(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	mockTeamMemberRepo := new(MockTeamMemberRepository)
	usecase := NewTeamUsecase(mockRepo, mockTeamMemberRepo, nil)

	team := &team.TeamWithRole{
		Team: team.Team{
			TeamID:      1,
			Name:        "Test Team",
			Description: "Test Description",
		},
		Role: "owner",
	}

	mockRepo.On("GetTeam", 1, 1).Return(team, nil)
	result, err := usecase.GetTeam(1, 1)

	assert.NoError(t, err)
	assert.Equal(t, team, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTeam_TeamNotFound(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	mockTeamMemberRepo := new(MockTeamMemberRepository)
	usecase := NewTeamUsecase(mockRepo, mockTeamMemberRepo, nil)
	mockRepo.On("GetTeam", 1, 1).Return(nil, errors.New("team not found"))

	_, err := usecase.GetTeam(1, 1)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTeam_TeamMemberNotFound(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	mockTeamMemberRepo := new(MockTeamMemberRepository)
	usecase := NewTeamUsecase(mockRepo, mockTeamMemberRepo, nil)
	mockRepo.On("GetTeam", 2, 1).Return(nil, errors.New("team not found"))

	_, err := usecase.GetTeam(2, 1)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
	mockTeamMemberRepo.AssertExpectations(t)
}
func TestUpdateTeam(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo, nil, nil)

	input := team.UpdateInput{
		ID:          1,
		Name:        "updated Test Team",
		Description: "updated Test Description",
	}

	mockRepo.On("UpdateTeam", 1, mock.AnythingOfType("*team.Team")).Return(nil)

	err := usecase.UpdateTeam(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTeam(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo, nil, nil)

	mockRepo.On("DeleteTeam", 1).Return(nil)

	err := usecase.DeleteTeam(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTeamsByCustomerID(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo, nil, nil)

	teams := []*team.TeamWithRole{
		{
			Team: team.Team{
				TeamID:      1,
				Name:        "Test Team",
				Description: "Test Description",
			},
			Role: "owner",
		},
		{
			Team: team.Team{
				TeamID:      2,
				Name:        "Test Team 2",
				Description: "Test Description 2",
			},
			Role: "member",
		},
	}
	mockRepo.On("GetTeamsByCustomerID", 1).Return(teams, nil)

	result, err := usecase.GetTeamsByCustomerID(1)

	assert.NoError(t, err)
	assert.Equal(t, teams, result)
	mockRepo.AssertExpectations(t)
}

func TestInviteToken(t *testing.T) {
	customerRepo := new(MockCustomerRepository)
	usecase := NewTeamUsecase(nil, nil, customerRepo)
	customerRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(customer.Customer{CustomerID: 1, Username: "test", Email: "test@example.com"}, nil)
	input := team.InviteTokenInput{
		TeamID: 1,
		Mail:   "test@example.com",
	}

	_, err := usecase.GenerateInviteToken(input)

	assert.NoError(t, err)
	customerRepo.AssertExpectations(t)
}

func TestInviteToken_CustomerNotFound(t *testing.T) {
	customerRepo := new(MockCustomerRepository)
	usecase := NewTeamUsecase(nil, nil, customerRepo)
	input := team.InviteTokenInput{
		TeamID: 1,
		Mail:   "test@example.com",
	}
	customerRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(customer.Customer{}, errors.New("customer not found"))

	_, err := usecase.GenerateInviteToken(input)

	assert.Error(t, err)
	customerRepo.AssertExpectations(t)
}

func TestJoinTeam(t *testing.T) {
	teamRepo := new(MockTeamRepository)
	teamMemberRepo := new(MockTeamMemberRepository)
	usecase := NewTeamUsecase(teamRepo, teamMemberRepo, nil)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": 1,
		"team_id":     1,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	input := team.JoinTeamInput{
		Token: tokenString,
	}

	// モックの戻り値として具体的なTeamWithRoleオブジェクトを作成
	mockTeam := &team.TeamWithRole{
		Team: team.Team{
			TeamID: 1,
		},
		Role: "member",
	}

	teamRepo.On("GetTeam", 1, 1).Return(mockTeam, nil)
	teamMemberRepo.On("AddTeamMember", mock.AnythingOfType("*team_member.TeamMember")).Return(nil)

	teamID, err := usecase.JoinTeam(1, input)

	assert.NoError(t, err)
	assert.Equal(t, 1, teamID)
	teamRepo.AssertExpectations(t)
	teamMemberRepo.AssertExpectations(t)
}

func TestJoinTeam_InvalidToken(t *testing.T) {
	mockRepo := new(MockTeamMemberRepository)
	usecase := NewTeamUsecase(nil, mockRepo, nil)
	tokenString := "invalid_token"
	input := team.JoinTeamInput{
		Token: tokenString,
	}

	teamID, err := usecase.JoinTeam(1, input)

	assert.Error(t, err)
	assert.Equal(t, 0, teamID)
	mockRepo.AssertExpectations(t)
}

func TestJoinTeam_InvalidCustomerID(t *testing.T) {
	mockRepo := new(MockTeamMemberRepository)
	usecase := NewTeamUsecase(nil, mockRepo, nil)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": 1,
		"team_id":     1,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	input := team.JoinTeamInput{
		Token: tokenString,
	}

	teamID, err := usecase.JoinTeam(2, input)

	assert.Error(t, err)
	assert.Equal(t, 0, teamID)
	mockRepo.AssertExpectations(t)
}

func TestJoinTeam_TeamNotFound(t *testing.T) {
	teamMemberRepo := new(MockTeamMemberRepository)
	teamRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(teamRepo, teamMemberRepo, nil)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": 1,
		"team_id":     1,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	input := team.JoinTeamInput{
		Token: tokenString,
	}
	teamRepo.On("GetTeam", 1, 1).Return(nil, errors.New("team not found"))

	teamID, err := usecase.JoinTeam(1, input)

	assert.Error(t, err)
	assert.Equal(t, 0, teamID)
	teamRepo.AssertExpectations(t)
}
