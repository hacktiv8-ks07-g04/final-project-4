package repository

import (
	"gorm.io/gorm"
)

type ProductsRepository interface{}

type ProductsRepositoryImpl struct {
	db *gorm.DB
}

func ProductsRepositoryInit(db *gorm.DB) *ProductsRepositoryImpl {
	return &ProductsRepositoryImpl{db}
}
