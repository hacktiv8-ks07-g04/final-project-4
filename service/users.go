package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-4/repository"
	"github.com/hacktiv8-ks07-g04/final-project-4/utils"
)

var err error

type UsersService interface {
	Register(req dto.RegisterRequest) (*entity.User, error)
	Login(email, password string) (string, error)
	TopUp(id uint, balance int) (int, error)
}

type UsersServiceImpl struct {
	usersRepository repository.UsersRepository
}

func UsersServiceInit(repository repository.UsersRepository) *UsersServiceImpl {
	return &UsersServiceImpl{repository}
}

func (u *UsersServiceImpl) Register(req dto.RegisterRequest) (*entity.User, error) {
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

func (u *UsersServiceImpl) Login(email, password string) (string, error) {
	user, err := u.usersRepository.Login(email, password)
	if err != nil {
		return "", err
	}

	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UsersServiceImpl) TopUp(id uint, balance int) (int, error) {
	user, err := u.usersRepository.TopUp(id, balance)
	if err != nil {
		return 0, err
	}

	return user.Balance, nil
}
