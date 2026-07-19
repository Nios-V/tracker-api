package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DurationType string

const (
	DurationTypeMinutes  DurationType = "minutes"
	DurationTypeHours    DurationType = "hours"
	DurationTypeDays     DurationType = "days"
	DurationTypeChapters DurationType = "chapters"
	DurationTypePages    DurationType = "pages"
	DurationTypeSections DurationType = "sections"
	DurationTypeTracks   DurationType = "tracks"
	DurationTypeEpisodes DurationType = "episodes"
	DurationTypeOther    DurationType = "other"
)

type Session struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ExperienceID uuid.UUID      `gorm:"not null" json:"experience_id"`
	Date         time.Time      `gorm:"not null" json:"date"`
	DurationMin  int            `gorm:"not null" json:"duration_min"`
	DurationType DurationType   `gorm:"not null" json:"duration_type"`
	Progress     string         `gorm:"type:text" json:"progress"`
	Notes        string         `gorm:"type:text" json:"notes"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
