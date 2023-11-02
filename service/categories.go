package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-4/repository"
)

type Categories interface {
	Create(req dto.CreateCategoryRequest) (entity.Category, error)
	GetAll() ([]entity.Category, error)
	Update(id, updatedType string) (entity.Category, error)
	Delete(id string) error
}

type CategoriesImpl struct {
	repository repository.Categories
}

func InitCategories(repository repository.Categories) *CategoriesImpl {
	return &CategoriesImpl{repository}
}

func (s *CategoriesImpl) Create(req dto.CreateCategoryRequest) (entity.Category, error) {
	category := entity.Category{
		Type: req.Type,
	}

	category, err := s.repository.Create(category)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *CategoriesImpl) GetAll() ([]entity.Category, error) {
	categories, err := s.repository.GetAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (s *CategoriesImpl) Update(id, updatedType string) (entity.Category, error) {
	category, err := s.repository.Update(id, updatedType)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *CategoriesImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
