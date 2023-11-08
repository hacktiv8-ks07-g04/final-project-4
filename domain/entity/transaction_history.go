package entity

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type TransactionHistory struct {
	Base
	Quantity   int  `gorm:"not null;type:int" json:"quantity,omitempty"    valid:"required~quantity is required,type(int)"`
	TotalPrice int  `gorm:"not null;type:int" json:"total_price,omitempty" valid:"required~total_price is required,type(int)"`
	UserID     uint `gorm:"not null;type:int" json:"user_id,omitempty"`
	ProductID  uint `gorm:"not null;type:int" json:"product_id,omitempty"`
}

func (t *TransactionHistory) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}

	var (
		user       User
		product    Product
		category   Category
		totalPrice int
	)

	// check product stock
	if err := tx.Where("id = ?", t.ProductID).First(&product).Error; err != nil {
		return errors.New("product not found")
	}

	if product.Stock < t.Quantity {
		return errors.New("stock is not enough")
	}

	totalPrice = product.Price * t.Quantity

	// check user balance
	if err := tx.Where("id = ?", t.UserID).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	if user.Balance < totalPrice {
		return errors.New("balance is not enough")
	}

	// update product stock
	if err := tx.Model(&product).Update("stock", product.Stock-t.Quantity).Error; err != nil {
		return errors.New("failed to update product stock")
	}

	// update user Balance
	if err := tx.Model(&user).Update("balance", user.Balance-totalPrice).Error; err != nil {
		return errors.New("failed to update user balance")
	}

	// update category sold
	if err := tx.Where("id = ?", product.CategoryID).First(&category).Error; err != nil {
		return errors.New("category not found")
	}

	if err := tx.Model(&category).Update("sold_product_amount", category.SoldProductAmount+t.Quantity).Error; err != nil {
		return errors.New("failed to update category sold")
	}

	return nil
}
