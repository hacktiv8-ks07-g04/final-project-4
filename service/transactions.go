package service

import "github.com/hacktiv8-ks07-g04/final-project-4/repository"

type Transactions interface{}

type TransactionsImpl struct {
	repository repository.Transactions
}

func InitTransactions(repository repository.Transactions) *TransactionsImpl {
	return &TransactionsImpl{repository}
}
