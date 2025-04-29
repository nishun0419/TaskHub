package teams

import (
	"backend/domain/teams"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTeamRepository struct {
	mock.Mock
}

func (m *MockTeamRepository) CreateTeam(team *teams.Team) error {
	args := m.Called(team)
	return args.Error(0)
}

func (m *MockTeamRepository) GetTeam(id int) (*teams.Team, error) {
	args := m.Called(id)
	return args.Get(0).(*teams.Team), args.Error(1)
}

func (m *MockTeamRepository) UpdateTeam(id int, team *teams.Team) error {
	args := m.Called(id, team)
	return args.Error(0)
}

func (m *MockTeamRepository) DeleteTeam(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTeamRepository) AddMember(teamID int, customerID int) error {
	args := m.Called(teamID, customerID)
	return args.Error(0)
}

func (m *MockTeamRepository) RemoveMember(teamID int, customerID int) error {
	args := m.Called(teamID, customerID)
	return args.Error(0)
}

func (m *MockTeamRepository) GetAllTeams() ([]*teams.Team, error) {
	args := m.Called()
	return args.Get(0).([]*teams.Team), args.Error(1)
}

func TestCreateTeam(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo)

	input := teams.CreateInput{
		Name:        "Test Team",
		Description: "Test Description",
	}

	mockRepo.On("CreateTeam", mock.AnythingOfType("*teams.Team")).Return(nil)

	err := usecase.CreateTeam(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTeam(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo)

	team := &teams.Team{
		ID:          1,
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
	usecase := NewTeamUsecase(mockRepo)

	input := teams.UpdateInput{
		ID:          1,
		Name:        "updated Test Team",
		Description: "updated Test Description",
	}

	mockRepo.On("UpdateTeam", 1, mock.AnythingOfType("*teams.Team")).Return(nil)

	err := usecase.UpdateTeam(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTeam(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo)

	mockRepo.On("DeleteTeam", 1).Return(nil)

	err := usecase.DeleteTeam(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddMember(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo)

	mockRepo.On("AddMember", 1, 1).Return(nil)

	err := usecase.AddMember(1, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRemoveMember(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo)

	mockRepo.On("RemoveMember", 1, 1).Return(nil)

	err := usecase.RemoveMember(1, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTeams(t *testing.T) {
	mockRepo := new(MockTeamRepository)
	usecase := NewTeamUsecase(mockRepo)

	teams := []*teams.Team{
		{
			ID:          1,
			Name:        "Test Team",
			Description: "Test Description",
		},
		{
			ID:          2,
			Name:        "Test Team 2",
			Description: "Test Description 2",
		},
	}

	mockRepo.On("GetAllTeams").Return(teams, nil)

	result, err := usecase.GetAllTeams()

	assert.NoError(t, err)
	assert.Equal(t, teams, result)
	mockRepo.AssertExpectations(t)
}
