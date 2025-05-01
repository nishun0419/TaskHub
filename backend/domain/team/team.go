package team

import (
	"time"
)

type Team struct {
	TeamID      int       `json:"team_id" gorm:"primary_key;auto_increment"`
	Name        string    `json:"name" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"size:255;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type TeamWithRole struct {
	Team
	Role string `json:"role"`
}
