package models

import "time"

type Note struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ExperienceID uint      `gorm:"not null" json:"experience_id"`
	Title        string    `gorm:"not null" json:"title"`
	Content      string    `gorm:"type:text" json:"content"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
