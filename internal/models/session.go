package models

import "time"

type Session struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ExperienceID uint      `gorm:"not null" json:"experience_id"`
	Date         time.Time `gorm:"not null" json:"date"`
	DurationMin  int       `gorm:"not null" json:"duration_min"`
	Progress     string    `gorm:"type:text" json:"progress"`
	Notes        string    `gorm:"type:text" json:"notes"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
