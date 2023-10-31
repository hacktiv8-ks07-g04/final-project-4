package repository

import (
	"gorm.io/gorm"
)

type CategoriesRepository interface{}

type CategoriesRepositoryImpl struct {
	db *gorm.DB
}

func CategoriesRepositoryInit(db *gorm.DB) *CategoriesRepositoryImpl {
	return &CategoriesRepositoryImpl{db}
}
