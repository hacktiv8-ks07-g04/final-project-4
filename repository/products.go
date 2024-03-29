package repository

import (
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
)

type Products interface {
	Create(product entity.Product) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	Update(id string, updatedProduct dto.UpdateProductRequest) (entity.Product, error)
	Delete(id string) error
}

type ProductsImpl struct {
	db *gorm.DB
}

func InitProducts(db *gorm.DB) *ProductsImpl {
	return &ProductsImpl{db}
}

func (r *ProductsImpl) Create(product entity.Product) (entity.Product, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Category").Create(&product).Error; err != nil {
			return err
		}

		return nil
	})

	return product, err
}

func (r *ProductsImpl) GetAll() ([]entity.Product, error) {
	var products []entity.Product

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&products).Error; err != nil {
			return err
		}

		return nil
	})

	return products, err
}

func (r *ProductsImpl) Update(
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

func (r *ProductsImpl) Delete(id string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).First(&entity.Product{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&entity.Product{}, id).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
