package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenericRepository[T any] struct {
	db *gorm.DB
}

func NewGenericRepository[T any](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{db: db}
}

func (r *GenericRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *GenericRepository[T]) GetByID(id uuid.UUID) (*T, error) {
	var entity T
	err := r.db.First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GenericRepository[T]) GetAll() ([]T, error) {
	var entities []T
	err := r.db.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *GenericRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *GenericRepository[T]) Delete(id uuid.UUID) error {
	// softdelete
	return r.db.Delete(new(T), "id = ?", id).Error
}
