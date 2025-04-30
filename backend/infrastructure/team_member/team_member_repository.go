package team_member

import (
	"backend/domain/team_member"

	"gorm.io/gorm"
)

type TeamMemberRepository struct {
	db *gorm.DB
}

func NewTeamMemberRepository(db *gorm.DB) *TeamMemberRepository {
	return &TeamMemberRepository{db: db}
}

func (r *TeamMemberRepository) AddTeamMember(teamMember *team_member.TeamMember) error {
	return r.db.Create(teamMember).Error
}

func (r *TeamMemberRepository) DeleteTeamMember(teamMemberDelInput *team_member.TeamMemberDelInput) error {
	return r.db.Where("team_id = ? AND customer_id = ?", teamMemberDelInput.TeamID, teamMemberDelInput.CustomerID).Delete(&team_member.TeamMember{}).Error
}
