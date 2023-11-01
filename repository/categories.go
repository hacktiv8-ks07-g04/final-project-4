package repository

import (
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
)

type CategoriesRepository interface {
	Create(category entity.Category) (entity.Category, error)
	Get(id string) (entity.Category, error)
	GetAll() ([]entity.Category, error)
	Update(id, updatedType string) (entity.Category, error)
	Delete(id string) error
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

func (cr *CategoriesRepositoryImpl) Get(id string) (entity.Category, error) {
	var category entity.Category

	err := cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Preload("Products").First(&category, id).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return category, err
	}

	return category, nil
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

func (cr *CategoriesRepositoryImpl) Update(id, updatedType string) (entity.Category, error) {
	var category entity.Category

	err := cr.db.Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Model(&category).Where("id = ?", id).Update("type", updatedType).Error
		if err != nil {
			return err
		}

		err = tx.Preload("Products").Find(&category).Error
		if err != nil {
			return err
		}

		return nil
	})

	return category, err
}

func (cr *CategoriesRepositoryImpl) Delete(id string) error {
	var category entity.Category
	var err error

	category, err = cr.Get(id)
	if err != nil {
		return err
	}

	err = cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&category, id).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
