package models

import (
	"time"

	"github.com/google/uuid"
)

type ExperienceStatus string

const (
	StatusPlanned    ExperienceStatus = "planned"
	StatusInProgress ExperienceStatus = "in_progress"
	StatusCompleted  ExperienceStatus = "completed"
	StatusDropped    ExperienceStatus = "dropped"
	StatusOnHold     ExperienceStatus = "on_hold"
)

type Experience struct {
	ID         uuid.UUID        `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID     uuid.UUID        `gorm:"not null" json:"user_id"`
	MediaID    uuid.UUID        `gorm:"not null" json:"media_id"`
	Media      Media            `gorm:"foreignKey:MediaID" json:"media,omitempty"`
	Status     ExperienceStatus `gorm:"not null" json:"status"`
	Rating     *float32         `gorm:"type:decimal(2,1)" json:"rating,omitempty"`
	StatedAt   *time.Time       `gorm:"type:timestamp" json:"stated_at,omitempty"`
	FinishedAt *time.Time       `gorm:"type:timestamp" json:"finished_at,omitempty"`
	CreatedAt  time.Time        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
}
