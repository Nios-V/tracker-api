package models

import (
	"time"

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
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Type      MediaType `gorm:"not null" json:"type"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TODO: Add abstraction fields for different media types (e.g., author, director, etc.) and implement polymorphic behavior if needed.
