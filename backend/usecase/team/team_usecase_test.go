package team

import (
	"backend/domain/team"
	"backend/domain/team_member"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTeamRepository struct {
	mock.Mock
}

type MockTeamMemberRepository struct {
	mock.Mock
}

func (m *MockTeamRepository) CreateTeam(t *team.Team) error {
	args := m.Called(t)
	return args.Error(0)
}

func (m *MockTeamRepository) GetTeam(id int) (*team.Team, error) {
	args := m.Called(id)
	return args.Get(0).(*team.Team), args.Error(1)
}

func (m *MockTeamRepository) UpdateTeam(id int, team *team.Team) error {
	args := m.Called(id, team)
	return args.Error(0)
}

func (m *MockTeamRepository) DeleteTeam(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTeamRepository) GetTeamsByCustomerID(customerID int) ([]*team.Team, error) {
	args := m.Called(customerID)
	return args.Get(0).([]*team.Team), args.Error(1)
}

func (m *MockTeamMemberRepository) AddTeamMember(teamMember *team_member.TeamMember) error {
	args := m.Called(teamMember)
	return args.Error(0)
}

func (m *MockTeamMemberRepository) DeleteTeamMember(teamMemberDelInput *team_member.TeamMemberDelInput) error {
	args := m.Called(teamMemberDelInput)
	return args.Error(0)
}

func TestCreateTeam(t *testing.T) {
	teamRepo := new(MockTeamRepository)
	teamMemberRepo := new(MockTeamMemberRepository)
	usecase := NewTeamUsecase(teamRepo, teamMemberRepo)

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
	usecase := NewTeamUsecase(mockRepo, nil)

	team := &team.Team{
		TeamID:      1,
		Name:        "Test Team",
		Description: "Test Description",
	}

	mockRepo.On("GetTeam", 1).Return(team, nil)

	result, err := usecase.GetTeam(1)

	assert.NoError(t, err)
	assert.Equal(t, team, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTeam(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo, nil)

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
	usecase := NewTeamUsecase(mockRepo, nil)

	mockRepo.On("DeleteTeam", 1).Return(nil)

	err := usecase.DeleteTeam(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
func TestGetTeamsByCustomerID(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo, nil)

	teams := []*team.Team{
		{
			TeamID:      1,
			Name:        "Test Team",
			Description: "Test Description",
		},
		{
			TeamID:      2,
			Name:        "Test Team 2",
			Description: "Test Description 2",
		},
	}
	mockRepo.On("GetTeamsByCustomerID", 1).Return(teams, nil)

	result, err := usecase.GetTeamsByCustomerID(1)

	assert.NoError(t, err)
	assert.Equal(t, teams, result)
	mockRepo.AssertExpectations(t)
}
