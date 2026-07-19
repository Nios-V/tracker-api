package services

import (
	"github.com/Nios-V/tracker-api/internal/models"
	"github.com/Nios-V/tracker-api/internal/repository"
)

type Services struct {
	Tag        *GenericService[models.Tag]
	NoteLink   *GenericService[models.NoteLink]
	MediaTag   *GenericService[models.MediaTag]
	Experience *GenericService[models.Experience]
	Session    *GenericService[models.Session]
	Note       *GenericService[models.Note]
	Media      *GenericService[models.Media]
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Tag:        NewGenericService[models.Tag](repos.Tag),
		NoteLink:   NewGenericService[models.NoteLink](repos.NoteLink),
		MediaTag:   NewGenericService[models.MediaTag](repos.MediaTag),
		Experience: NewGenericService[models.Experience](repos.Experience),
		Session:    NewGenericService[models.Session](repos.Session),
		Note:       NewGenericService[models.Note](repos.Note),
		Media:      NewGenericService[models.Media](repos.Media),
	}
}
