package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/repository"
)

type Transactions interface {
	Create(userID uint, req dto.CreateTransactionRequest) (dto.CreateTransactionResponse, error)
	GetUserTransactions(userID uint) ([]dto.TransactionHistory, error)
}

type TransactionsImpl struct {
	repository repository.Transactions
}

func InitTransactions(repository repository.Transactions) *TransactionsImpl {
	return &TransactionsImpl{repository}
}

func (s TransactionsImpl) Create(
	userID uint,
	req dto.CreateTransactionRequest,
) (dto.CreateTransactionResponse, error) {
	transactionBill, err := s.repository.Create(userID, req)
	if err != nil {
		return dto.CreateTransactionResponse{}, err
	}

	response := dto.CreateTransactionResponse{
		Message: "You have successfully purchased the product",
		TransactionBill: dto.TransactionBill{
			TotalPrice:   transactionBill.TotalPrice,
			Quantity:     transactionBill.Quantity,
			ProductTitle: transactionBill.ProductTitle,
		},
	}

	return response, nil
}

func (s TransactionsImpl) GetUserTransactions(userID uint) ([]dto.TransactionHistory, error) {
	transactions, err := s.repository.GetUserTransactions(userID)
	if err != nil {
		return nil, err
	}

	return transactions, err
}
