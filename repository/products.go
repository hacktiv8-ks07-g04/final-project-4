package repository

import (
	"gorm.io/gorm"
)

type ProductRepository interface{}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func ProductRepositoryInit(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}
