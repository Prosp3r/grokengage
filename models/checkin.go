package models

import "time"

type CheckIn struct {
	ID     uint      `gorm:"primaryKey" json:"id"`
	UserID uint      `json:"user_id"`
	Date   time.Time `json:"date"`
	Points int       `json:"points"`
}
