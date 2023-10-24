package entity

type TransactionHistory struct {
	Base
	Quantity   int  `gorm:"not null;type:int" json:"quantity,omitempty"`
	TotalPrice int  `gorm:"not null;type:int" json:"total_price,omitempty"`
	UserID     uint `gorm:"not null;type:int" json:"user_id,omitempty"`
	ProductID  uint `gorm:"not null;type:int" json:"product_id,omitempty"`
}
