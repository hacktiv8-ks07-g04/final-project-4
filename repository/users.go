package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
)

type Users interface {
	Register(user entity.User) (entity.User, error)
	Login(email, password string) (entity.User, error)
	TopUp(id uint, balance int) (entity.User, error)
}

type UsersImpl struct {
	db *gorm.DB
}

func InitUsers(db *gorm.DB) *UsersImpl {
	return &UsersImpl{db}
}

func (r *UsersImpl) Register(user entity.User) (entity.User, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&user).Error
		if err != nil {
			return err
		}

		return nil
	})

	return user, err
}

func (r *UsersImpl) Login(email, password string) (entity.User, error) {
	var user entity.User

	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("email = ?", email).First(&user).Error
		if err != nil {
			return errors.New("user not found")
		}

		return nil
	})

	return user, err
}

func (r *UsersImpl) TopUp(id uint, balance int) (entity.User, error) {
	var user entity.User

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
			return errors.New("user not found")
		}

		user.Balance += balance

		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		return nil
	})

	return user, err
}
