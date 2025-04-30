package team_member

import (
	"backend/domain/team_member"
)

type TeamMemberUsecase struct {
	TeamMemberRepository team_member.TeamMemberRepository
}

func NewTeamMemberUsecase(repo team_member.TeamMemberRepository) *TeamMemberUsecase {
	return &TeamMemberUsecase{repo}
}

func (u *TeamMemberUsecase) AddTeamMember(teamMember *team_member.TeamMember) error {
	return u.TeamMemberRepository.AddTeamMember(teamMember)
}

func (u *TeamMemberUsecase) DeleteTeamMember(teamMemberDelInput *team_member.TeamMemberDelInput) error {
	return u.TeamMemberRepository.DeleteTeamMember(teamMemberDelInput)
}
