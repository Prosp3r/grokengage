package models

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	Type        string `json:"type"` // e.g., "follow", "like", "repost", "visit", "comment"
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id"`
}
