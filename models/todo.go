package models

type Todo struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Task   string `gorm:"not null" json:"task"`
	Status string `gorm:"not null" json:"status"`
}
