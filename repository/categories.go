package repository

import (
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
)

type CategoriesRepository interface {
	Create(category entity.Category) (entity.Category, error)
	GetAll() ([]entity.Category, error)
}

type CategoriesRepositoryImpl struct {
	db *gorm.DB
}

func CategoriesRepositoryInit(db *gorm.DB) *CategoriesRepositoryImpl {
	return &CategoriesRepositoryImpl{db}
}

func (cr *CategoriesRepositoryImpl) Create(category entity.Category) (entity.Category, error) {
	err := cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&category).Error
		if err != nil {
			return err
		}

		return nil
	})

	return category, err
}

func (cr *CategoriesRepositoryImpl) GetAll() ([]entity.Category, error) {
	var categories []entity.Category

	err := cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Preload("Products").Find(&categories).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return categories, err
	}

	return categories, nil
}
