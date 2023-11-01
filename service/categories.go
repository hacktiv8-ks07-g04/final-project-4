package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-4/repository"
)

type CategoriesService interface {
	Create(req dto.CreateCategoryRequest) (entity.Category, error)
	GetAll() ([]entity.Category, error)
}

type CategoriesServiceImpl struct {
	categoriesRepo repository.CategoriesRepository
}

func CategoriesServiceInit(repository repository.CategoriesRepository) *CategoriesServiceImpl {
	return &CategoriesServiceImpl{repository}
}

func (c *CategoriesServiceImpl) Create(req dto.CreateCategoryRequest) (entity.Category, error) {
	category := entity.Category{
		Type: req.Type,
	}

	category, err := c.categoriesRepo.Create(category)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (c *CategoriesServiceImpl) GetAll() ([]entity.Category, error) {
	categories, err := c.categoriesRepo.GetAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}
