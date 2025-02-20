package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Username    string    `gorm:"unique" json:"username"`
	Points      int       `json:"points"`
	Streak      int       `json:"streak"`
	LastCheckIn time.Time `json:"last_check_in"`
	CreatedAt   time.Time `json:"created_at"`
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	u.Points = 0
	u.Streak = 0
	return
}
