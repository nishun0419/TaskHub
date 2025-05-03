package team_member

type TeamMember struct {
	TeamMemberID int    `json:"team_member_id" gorm:"primaryKey;autoIncrement"`
	TeamID       int    `json:"team_id" gorm:"not null"`
	CustomerID   int    `json:"customer_id" gorm:"not null"`
	Role         string `json:"role" gorm:"not null"`
}
