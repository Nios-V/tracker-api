package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaType string

const (
	MediaTypeBook   MediaType = "book"
	MediaTypeMovie  MediaType = "movie"
	MediaTypeTVShow MediaType = "tvshow"
	MediaTypeMusic  MediaType = "album"
	MediaTypeGame   MediaType = "videogame"
	MediaTypeOther  MediaType = "other"
)

type Media struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Type      MediaType      `gorm:"not null" json:"type"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TODO: Add abstraction fields for different media types (e.g., author, director, etc.) and implement polymorphic behavior if needed.
