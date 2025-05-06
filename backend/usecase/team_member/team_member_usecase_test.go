package team_member

import (
	"backend/domain/team_member"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTeamMemberRepository struct {
	mock.Mock
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
	usecase := NewTeamMemberUsecase(mockRepo)

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
	usecase := NewTeamMemberUsecase(mockRepo)

	input := team_member.TeamMemberDelInput{
		TeamID:     1,
		CustomerID: 1,
	}

	mockRepo.On("DeleteTeamMember", mock.AnythingOfType("*team_member.TeamMemberDelInput")).Return(nil)

	err := usecase.DeleteTeamMember(&input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
