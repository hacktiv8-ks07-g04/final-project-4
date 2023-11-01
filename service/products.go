package service

import "github.com/hacktiv8-ks07-g04/final-project-4/repository"

type ProductsService interface{}

type ProductsServiceImpl struct {
	productsRepo repository.ProductsRepository
}

func ProductsServiceInit(repository repository.ProductsRepository) *ProductsServiceImpl {
	return &ProductsServiceImpl{repository}
}
