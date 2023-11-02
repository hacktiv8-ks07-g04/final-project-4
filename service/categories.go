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
	categoriesRepo repository.Categories
}

func InitCategories(repository repository.Categories) *CategoriesImpl {
	return &CategoriesImpl{repository}
}

func (c *CategoriesImpl) Create(req dto.CreateCategoryRequest) (entity.Category, error) {
	category := entity.Category{
		Type: req.Type,
	}

	category, err := c.categoriesRepo.Create(category)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (c *CategoriesImpl) GetAll() ([]entity.Category, error) {
	categories, err := c.categoriesRepo.GetAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (c *CategoriesImpl) Update(id, updatedType string) (entity.Category, error) {
	category, err := c.categoriesRepo.Update(id, updatedType)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (c *CategoriesImpl) Delete(id string) error {
	err := c.categoriesRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
