package handler

import "github.com/hacktiv8-ks07-g04/final-project-4/service"

type ProductsHandler interface{}

type ProductsHandlerImpl struct {
	productsService service.ProductsService
}

func ProductsHandlerInit(service service.ProductsService) *ProductsHandlerImpl {
	return &ProductsHandlerImpl{service}
}
