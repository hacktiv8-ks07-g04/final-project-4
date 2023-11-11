package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/repository"
)

type Transactions interface {
	Create(userID uint, req dto.CreateTransactionRequest) (dto.CreateTransactionResponse, error)
	GetUserTransactions(userID uint) ([]dto.TransactionHistory, error)
	GetAll() ([]dto.GetAllTransactionsResponse, error)
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

func (s TransactionsImpl) GetAll() ([]dto.GetAllTransactionsResponse, error) {
	var response []dto.GetAllTransactionsResponse

	transactions, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	for _, transaction := range transactions {
		response = append(response, dto.GetAllTransactionsResponse{
			ID:         transaction.ID,
			ProductID:  transaction.ProductID,
			UserID:     transaction.UserID,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: dto.Product{
				ID:         transaction.Product.ID,
				Title:      transaction.Product.Title,
				Price:      transaction.Product.Price,
				Stock:      transaction.Product.Stock,
				CategoryID: transaction.Product.CategoryID,
				CreatedAt:  transaction.Product.CreatedAt.String(),
			},
			User: dto.User{
				ID:        transaction.User.ID,
				FullName:  transaction.User.FullName,
				Email:     transaction.User.Email,
				Balance:   transaction.User.Balance,
				CreatedAt: transaction.User.CreatedAt.String(),
				UpdatedAt: transaction.User.UpdatedAt.String(),
			},
		})
	}

	return response, err
}
