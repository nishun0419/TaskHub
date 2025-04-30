package teams

import (
	"backend/domain/teams"

	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) CreateTeam(team *teams.Team) error {
	return r.db.Create(team).Error
}

func (r *TeamRepository) AddMember(teamID int, customerID int) error {
	return r.db.Exec("INSERT INTO team_members (team_id, customer_id) VALUES (?, ?)", teamID, customerID).Error
}

func (r *TeamRepository) GetTeam(id int) (*teams.Team, error) {
	var team teams.Team
	err := r.db.First(&team, id).Error
	return &team, err
}

func (r *TeamRepository) UpdateTeam(id int, team *teams.Team) error {
	return r.db.Model(&teams.Team{}).Where("id = ?", id).Updates(team).Error
}

func (r *TeamRepository) DeleteTeam(id int) error {
	return r.db.Delete(&teams.Team{}, id).Error
}

func (r *TeamRepository) RemoveMember(teamID int, customerID int) error {
	return r.db.Exec("DELETE FROM team_members WHERE team_id = ? AND customer_id = ?", teamID, customerID).Error
}

func (r *TeamRepository) GetTeamsByCustomerID(customerID int) ([]*teams.Team, error) {
	var teams []*teams.Team
	err := r.db.Joins("JOIN team_members ON teams.id = team_members.team_id").
		Where("team_members.customer_id = ?", customerID).
		Find(&teams).Error
	return teams, err
}
