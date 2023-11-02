package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-4/repository"
	"github.com/hacktiv8-ks07-g04/final-project-4/utils"
)

var err error

type Users interface {
	Register(req dto.RegisterRequest) (*entity.User, error)
	Login(email, password string) (string, error)
	TopUp(id uint, balance int) (int, error)
}

type UsersImpl struct {
	usersRepository repository.UsersRepository
}

func InitUsers(repository repository.UsersRepository) *UsersImpl {
	return &UsersImpl{repository}
}

func (u *UsersImpl) Register(req dto.RegisterRequest) (*entity.User, error) {
	user := entity.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err = u.usersRepository.Register(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UsersImpl) Login(email, password string) (string, error) {
	user, err := u.usersRepository.Login(email, password)
	if err != nil {
		return "", err
	}

	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID, user.Email, string(user.Role))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UsersImpl) TopUp(id uint, balance int) (int, error) {
	user, err := u.usersRepository.TopUp(id, balance)
	if err != nil {
		return 0, err
	}

	return user.Balance, nil
}
