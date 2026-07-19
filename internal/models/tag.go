package models

import "github.com/google/uuid"

type Tag struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"not null;unique" json:"name"`
}
