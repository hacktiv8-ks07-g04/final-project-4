package repository

import (
	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
)

type UsersRepository interface {
	Register(user entity.User) (entity.User, error)
}

type UsersRepositoryImpl struct {
	db *gorm.DB
}

func UsersRepositoryInit(db *gorm.DB) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{db}
}

func (ur *UsersRepositoryImpl) Register(user entity.User) (entity.User, error) {
	err := ur.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&user).Error
		if err != nil {
			return err
		}

		return nil
	})

	return user, err
}
