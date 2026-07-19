package handlers

import (
	"github.com/Nios-V/tracker-api/internal/models"
	"github.com/Nios-V/tracker-api/internal/services"
)

type Handlers struct {
	Tag        *GenericHandler[models.Tag]
	Media      *GenericHandler[models.Media]
	MediaTag   *GenericHandler[models.MediaTag]
	Experience *GenericHandler[models.Experience]
	Session    *GenericHandler[models.Session]
	Note       *GenericHandler[models.Note]
	NoteLink   *GenericHandler[models.NoteLink]
}

func NewHandlers(s *services.Services) *Handlers {
	return &Handlers{
		Tag:        NewGenericHandler(s.Tag),
		Media:      NewGenericHandler(s.Media),
		MediaTag:   NewGenericHandler(s.MediaTag),
		Experience: NewGenericHandler(s.Experience),
		Session:    NewGenericHandler(s.Session),
		Note:       NewGenericHandler(s.Note),
		NoteLink:   NewGenericHandler(s.NoteLink),
	}
}
