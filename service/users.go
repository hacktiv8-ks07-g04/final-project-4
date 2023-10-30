package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
	"github.com/hacktiv8-ks07-g04/final-project-4/repository"
)

var err error

type UsersService interface {
	Register(user entity.User) (*dto.RegisterResponse, error)
}

type UsersServiceImpl struct {
	usersRepository repository.UsersRepository
}

func UsersServiceInit(repository repository.UsersRepository) *UsersServiceImpl {
	return &UsersServiceImpl{repository}
}

func (u *UsersServiceImpl) Register(user entity.User) (*dto.RegisterResponse, error) {
	user, err = u.usersRepository.Register(user)

	response := dto.RegisterResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}

	if err != nil {
		return nil, err
	}

	return &response, nil
}
