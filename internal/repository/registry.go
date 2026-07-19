package repository

import (
	"github.com/Nios-V/tracker-api/internal/models"
	"gorm.io/gorm"
)

type Repositories struct {
	Tag        *GenericRepository[models.Tag]
	NoteLink   *GenericRepository[models.NoteLink]
	MediaTag   *GenericRepository[models.MediaTag]
	Experience *GenericRepository[models.Experience]
	Session    *GenericRepository[models.Session]
	Note       *GenericRepository[models.Note]
	Media      *GenericRepository[models.Media]
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Tag:        NewGenericRepository[models.Tag](db),
		NoteLink:   NewGenericRepository[models.NoteLink](db),
		MediaTag:   NewGenericRepository[models.MediaTag](db),
		Experience: NewGenericRepository[models.Experience](db),
		Session:    NewGenericRepository[models.Session](db),
		Note:       NewGenericRepository[models.Note](db),
		Media:      NewGenericRepository[models.Media](db),
	}
}
