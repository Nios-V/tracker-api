package models

import "time"

type ExperienceStatus string

const (
	StatusPlanned    ExperienceStatus = "planned"
	StatusInProgress ExperienceStatus = "in_progress"
	StatusCompleted  ExperienceStatus = "completed"
	StatusDropped    ExperienceStatus = "dropped"
	StatusOnHold     ExperienceStatus = "on_hold"
)

type Experience struct {
	ID         uint             `gorm:"primaryKey" json:"id"`
	UserID     uint             `gorm:"not null" json:"user_id"`
	MediaID    uint             `gorm:"not null" json:"media_id"`
	Media      Media            `gorm:"foreignKey:MediaID" json:"media,omitempty"`
	Status     ExperienceStatus `gorm:"not null" json:"status"`
	Rating     *float32         `gorm:"type:decimal(2,1)" json:"rating,omitempty"`
	StatedAt   *time.Time       `gorm:"type:timestamp" json:"stated_at,omitempty"`
	FinishedAt *time.Time       `gorm:"type:timestamp" json:"finished_at,omitempty"`
	CreatedAt  time.Time        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
}
