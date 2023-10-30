package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Title        string               `gorm:"not null;type:varchar(255)"                       json:"title,omitempty"        valid:"required~title is required,type(string)"`
	Price        int                  `gorm:"not null;type:int"                                json:"price,omitempty"        valid:"required~price is required,type(int),range(0|50000000)~price must be between 0 and 50.000.000"`
	Stock        int                  `gorm:"not null;type:int"                                json:"stock,omitempty"        valid:"required~stock is required,type(int),min(5)~stock must be at least 5"`
	CategoryID   uint                 `gorm:"not null;type:int"                                json:"category_id,omitempty"`
	Transactions []TransactionHistory `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE" json:"transactions,omitempty"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}
