package entity

type Product struct {
	Base
	Title        string               `gorm:"not null;type:varchar(255)"                       json:"title,omitempty"`
	Price        int                  `gorm:"not null;type:int"                                json:"price,omitempty"`
	Stock        int                  `gorm:"not null;type:int"                                json:"stock,omitempty"`
	CategoryID   uint                 `gorm:"not null;type:int"                                json:"category_id,omitempty"`
	Transactions []TransactionHistory `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE" json:"transactions,omitempty"`
}
