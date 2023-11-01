package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
)

type ProductsRepository interface {
	Add(product entity.Product) (entity.Product, error)
	GetAll() ([]entity.Product, error)
}

type ProductsRepositoryImpl struct {
	db *gorm.DB
}

func ProductsRepositoryInit(db *gorm.DB) *ProductsRepositoryImpl {
	return &ProductsRepositoryImpl{db}
}

func (r *ProductsRepositoryImpl) Add(product entity.Product) (entity.Product, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Category").Create(&product).Error; err != nil {
			return err
		}

		return nil
	})

	return product, err
}

func (r *ProductsRepositoryImpl) GetAll() ([]entity.Product, error) {
	var products []entity.Product

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&products).Error; err != nil {
			return err
		}

		return nil
	})

	return products, err
}

func (r *ProductsRepositoryImpl) checkCategory(categoryID uint) error {
	var category entity.Category

	if err := r.db.Where("id = ?", categoryID).First(&category).Error; err != nil {
		return errors.New("category not found")
	}

	return nil
}
