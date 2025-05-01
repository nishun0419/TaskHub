package team

import (
	"backend/domain/team"

	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) CreateTeam(team *team.Team) error {
	return r.db.Create(team).Error
}

func (r *TeamRepository) GetTeam(id int) (*team.Team, error) {
	var team team.Team
	err := r.db.First(&team, id).Error
	return &team, err
}

func (r *TeamRepository) UpdateTeam(id int, team *team.Team) error {
	return r.db.Model(team).Where("id = ?", id).Updates(team).Error
}

func (r *TeamRepository) DeleteTeam(id int) error {
	return r.db.Delete(&team.Team{}, id).Error
}

func (r *TeamRepository) GetTeamsByCustomerID(customerID int) ([]*team.TeamWithRole, error) {
	var teams []*team.TeamWithRole
	err := r.db.Select("teams.*, team_members.role").
		Table("teams").
		Joins("JOIN team_members ON teams.team_id = team_members.team_id").
		Where("team_members.customer_id = ?", customerID).
		Find(&teams).Error
	return teams, err
}
