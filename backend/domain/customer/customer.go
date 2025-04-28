package customer

import (
	"time"
)

type Customer struct {
	CustomerID int    `json:"customer_id" gorm:"primary_key;auto_increment"`
	Username   string `json:"username" gorm:"size:255;not null"`
	Email      string `json:"email" gorm:"size:255;not null"`
	Password   string `json:"password" gorm:"size:255;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
