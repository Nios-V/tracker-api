package models

type NoteLink struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	NoteID uint `gorm:"not null" json:"note_id"`
	// LinkedNoteID is the ID of the note that this note is linked to.
	LinkedNoteID uint `gorm:"not null" json:"linked_note_id"`
}
