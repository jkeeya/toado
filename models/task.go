package models

type Task struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"size:255;not null" json:"name"`
	ExpDate string `gorm:"not null" json:"exp_date"`
	Done    bool   `json:"done"`
}
