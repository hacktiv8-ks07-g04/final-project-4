package service

import "github.com/hacktiv8-ks07-g04/final-project-4/repository"

type CategoriesService interface{}

type CategoriesServiceImpl struct {
	categoriesRepo repository.CategoriesRepository
}

func CategoriesServiceInit(repository repository.CategoriesRepository) *CategoriesServiceImpl {
	return &CategoriesServiceImpl{repository}
}
