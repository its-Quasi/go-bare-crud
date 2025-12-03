package service

import (
	"bare-crud/internal/models"
	"bare-crud/internal/store"
)

type Service struct {
	store store.Store
}

func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) GetAll() ([]models.Book, error) {
	return s.store.GetAll()
}

func (s *Service) GetById(id int) (models.Book, error) {
	return s.store.GetById(id)
}

func (s *Service) Create(bookData models.Book) (models.Book, error) {
	return s.store.Create(bookData)
}
