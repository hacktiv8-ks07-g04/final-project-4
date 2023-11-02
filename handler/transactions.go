package handler

import "github.com/hacktiv8-ks07-g04/final-project-4/service"

type TransactionsHandler interface{}

type TransactionsHandlerImpl struct {
	transactionsService service.Transactions
}

func TransactionsHandlerInit(service service.Transactions) *TransactionsHandlerImpl {
	return &TransactionsHandlerImpl{service}
}
