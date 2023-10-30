package entity

import (
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

	return nil
}
