package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-4/utils"
)

// Role is a custom type for user role
// It is used to create enum type in database
type Role string

const (
	ADMIN    Role = "admin"
	CUSTOMER Role = "customer"
)

type User struct {
	Base
	FullName     string               `gorm:"not null;type:varchar(255)"                    json:"full_name,omitempty"    valid:"required~full name is required,type(string)"`
	Email        string               `gorm:"not null;type:varchar(255);unique"             json:"email,omitempty"        valid:"required~email is required,type(string),email~email must be a valid email address"`
	Password     string               `gorm:"not null;type:varchar(255)"                    json:"password,omitempty"     valid:"required~password is required,minstringlength(6)~password must be at least 6 characters long"`
	Role         Role                 `gorm:"not null;type:role;default:'customer'"         json:"role,omitempty"`
	Balance      int                  `gorm:"not null;type:int"                             json:"balance,omitempty"      valid:"type(int),range(0|100000000)~balance must be between 0 and 100.000.000"`
	Transactions []TransactionHistory `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE" json:"transactions,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return nil
}
