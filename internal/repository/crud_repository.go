package repository

import (
	"auth-service/internal/database"

	"gorm.io/gorm"
)

type CRUD[T any] interface {
	Create(entity *T) error
	Get(id any) (*T, error)
	Update(entity *T) error
	Delete(id any) error
	FirstByField(field string, value any) (*T, error)
}

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{db: database.Get()}
}

func (r *Repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *Repository[T]) Get(id any) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	return &entity, err
}

func (r *Repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repository[T]) Delete(id any) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}

func (r *Repository[T]) FirstByField(field string, value any) (*T, error) {
	var entity T
	err := r.db.Where(field+" = ?", value).First(&entity).Error
	return &entity, err
}
