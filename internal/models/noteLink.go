package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteLink struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	NoteID uuid.UUID `gorm:"not null" json:"note_id"`
	// LinkedNoteID is the ID of the note that this note is linked to.
	LinkedNoteID uuid.UUID      `gorm:"not null" json:"linked_note_id"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
