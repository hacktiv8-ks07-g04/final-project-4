package entity

// Role is a custom type for user role
// It is used to create enum type in database
type Role string

const (
	ADMIN    Role = "admin"
	CUSTOMER Role = "customer"
)

type User struct {
	Base
	FullName     string               `gorm:"not null;type:varchar(255)"                    json:"full_name,omitempty"`
	Email        string               `gorm:"not null;type:varchar(255);unique"             json:"email,omitempty"`
	Password     string               `gorm:"not null;type:varchar(255)"                    json:"password,omitempty"`
	Role         Role                 `gorm:"not null;type:role;default:'customer'"         json:"role,omitempty"`
	Balance      int                  `gorm:"not null;type:int"                             json:"balance,omitempty"`
	Transactions []TransactionHistory `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE" json:"transactions,omitempty"`
}
