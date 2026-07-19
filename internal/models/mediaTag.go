package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaTag struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	MediaID   uuid.UUID      `gorm:"not null" json:"media_id"`
	TagID     uuid.UUID      `gorm:"not null" json:"tag_id"`
	Tag       Tag            `gorm:"foreignKey:TagID" json:"tag,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
