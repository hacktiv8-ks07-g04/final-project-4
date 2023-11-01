package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-4/repository"
)

type ProductsService interface {
	Add(req dto.CreateProductRequest) (entity.Product, error)
}

type ProductsServiceImpl struct {
	productsRepo repository.ProductsRepository
}

func ProductsServiceInit(repository repository.ProductsRepository) *ProductsServiceImpl {
	return &ProductsServiceImpl{repository}
}

func (s *ProductsServiceImpl) Add(req dto.CreateProductRequest) (entity.Product, error) {
	product := entity.Product{
		Title:      req.Title,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
	}

	product, err := s.productsRepo.Add(product)
	if err != nil {
		return product, err
	}

	return product, nil
}
