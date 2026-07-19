package services

import (
	"github.com/Nios-V/tracker-api/internal/repository"
	"github.com/google/uuid"
)

type GenericService[T any] struct {
	repo *repository.GenericRepository[T]
}

func NewGenericService[T any](repo *repository.GenericRepository[T]) *GenericService[T] {
	return &GenericService[T]{repo: repo}
}

func (s *GenericService[T]) Create(entity *T) error {
	return s.repo.Create(entity)
}

func (s *GenericService[T]) GetByID(id uuid.UUID) (*T, error) {
	return s.repo.GetByID(id)
}

func (s *GenericService[T]) GetAll() ([]T, error) {
	return s.repo.GetAll()
}

func (s *GenericService[T]) Update(entity *T) error {
	return s.repo.Update(entity)
}

func (s *GenericService[T]) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
