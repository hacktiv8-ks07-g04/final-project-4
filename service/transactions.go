package service

import "github.com/hacktiv8-ks07-g04/final-project-4/repository"

type TransactionsService interface{}

type TransactionsServiceImpl struct {
	transactionsRepo repository.TransactionsRepository
}

func TransactionsServiceInit(repository repository.TransactionsRepository) *TransactionsServiceImpl {
	return &TransactionsServiceImpl{repository}
}
