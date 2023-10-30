package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Category struct {
	Base
	Type              string    `gorm:"not null;type:varchar(255)"                        json:"type,omitempty"                valid:"required,type(string)"`
	SoldProductAmount int       `gorm:"not null;type:int"                                 json:"sold_product_amount,omitempty"`
	Products          []Product `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE" json:"products,omitempty"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}
