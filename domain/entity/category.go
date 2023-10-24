package entity

type Category struct {
	Base
	Type              string    `gorm:"not null;type:varchar(255)"                        json:"type,omitempty"`
	SoldProductAmount int       `gorm:"not null;type:int"                                 json:"sold_product_amount,omitempty"`
	Products          []Product `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE" json:"products,omitempty"`
}
