package todo

import "time"

type Todo struct {
	TodoID      int       `json:"todo_id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Completed   bool      `json:"completed" gorm:"not null"`
	TeamID      int       `json:"team_id" gorm:"not null"`
	CustomerID  int       `json:"customer_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null;autoUpdateTime"`
}
