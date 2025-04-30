package teams

import (
	"time"
)

type Team struct {
	TeamID      int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
