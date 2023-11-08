package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-4/repository"
)

type Products interface {
	Create(req dto.CreateProductRequest) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	Update(id string, req dto.UpdateProductRequest) (entity.Product, error)
	Delete(id string) error
}

type ProductsImpl struct {
	repository repository.Products
}

func InitProducts(repository repository.Products) *ProductsImpl {
	return &ProductsImpl{repository}
}

func (s *ProductsImpl) Create(req dto.CreateProductRequest) (entity.Product, error) {
	product := entity.Product{
		Title:      req.Title,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
	}

	product, err := s.repository.Create(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductsImpl) GetAll() ([]entity.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *ProductsImpl) Update(
	id string,
	req dto.UpdateProductRequest,
) (entity.Product, error) {
	product, err := s.repository.Update(id, req)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductsImpl) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
