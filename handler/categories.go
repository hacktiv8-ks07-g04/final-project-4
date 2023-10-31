package handler

import "github.com/hacktiv8-ks07-g04/final-project-4/service"

type CategoriesHandler interface{}

type CategoriesHandlerImpl struct {
	categoriesService service.CategoriesService
}

func CategoriesHandlerInit(service *service.CategoriesService) *CategoriesHandlerImpl {
	return &CategoriesHandlerImpl{service}
}
