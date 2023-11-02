package service

import "github.com/hacktiv8-ks07-g04/final-project-4/repository"

type Transactions interface{}

type TransactionsImpl struct {
	transactionsRepo repository.TransactionsRepository
}

func InitTransactions(repository repository.TransactionsRepository) *TransactionsImpl {
	return &TransactionsImpl{repository}
}
