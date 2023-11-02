package handler

import "github.com/hacktiv8-ks07-g04/final-project-4/service"

type Transactions interface{}

type TransactionsImpl struct {
	transactionsService service.Transactions
}

func InitTransactions(service service.Transactions) *TransactionsImpl {
	return &TransactionsImpl{service}
}
