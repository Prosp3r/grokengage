package models

type Reward struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	PointsCost  int    `json:"points_cost"`
	Description string `json:"description"`
}
