package repository

import (
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
)

type ProductsRepository interface {
	Create(product entity.Product) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	Update(id string, updatedProduct dto.UpdateProductRequest) (entity.Product, error)
}

type ProductsRepositoryImpl struct {
	db *gorm.DB
}

func ProductsRepositoryInit(db *gorm.DB) *ProductsRepositoryImpl {
	return &ProductsRepositoryImpl{db}
}

func (r *ProductsRepositoryImpl) Create(product entity.Product) (entity.Product, error) {
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

func (r *ProductsRepositoryImpl) Update(
	id string,
	updatedProduct dto.UpdateProductRequest,
) (entity.Product, error) {
	var product entity.Product

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).First(&product).Error; err != nil {
			return err
		}

		if err := tx.Model(&product).Updates(updatedProduct).Error; err != nil {
			return err
		}

		if err := tx.First(&product).Error; err != nil {
			return err
		}

		return nil
	})

	return product, err
}
