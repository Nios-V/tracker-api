package models

type MediaTag struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	MediaID uint `gorm:"not null" json:"media_id"`
	TagID   uint `gorm:"not null" json:"tag_id"`
	Tag     Tag  `gorm:"foreignKey:TagID" json:"tag,omitempty"`
}
