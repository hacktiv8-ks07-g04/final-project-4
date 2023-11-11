package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
)

type Transactions interface {
	Create(userID uint, req dto.CreateTransactionRequest) (dto.TransactionBill, error)
	GetUserTransactions(userID uint) ([]dto.TransactionHistory, error)
	GetAll() ([]entity.TransactionHistory, error)
}

type TransactionsImpl struct {
	db *gorm.DB
}

func InitTransactions(db *gorm.DB) *TransactionsImpl {
	return &TransactionsImpl{db}
}

func (r *TransactionsImpl) Create(
	userID uint,
	req dto.CreateTransactionRequest,
) (dto.TransactionBill, error) {
	var (
		product            entity.Product
		transactionHistory dto.TransactionHistory
		totalPrice         int
		transactionBill    dto.TransactionBill
	)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&product, req.ProductID).Error; err != nil {
			return errors.New("product not found")
		}

		totalPrice = product.Price * req.Quantity

		transactionHistory = dto.TransactionHistory{
			Quantity:   req.Quantity,
			TotalPrice: totalPrice,
			UserID:     userID,
			ProductID:  uint(req.ProductID),
		}

		if err := tx.Create(&transactionHistory).Error; err != nil {
			return err
		}

		transactionBill.ProductTitle = product.Title
		transactionBill.TotalPrice = totalPrice
		transactionBill.Quantity = req.Quantity

		return nil
	})

	return transactionBill, err
}

func (r *TransactionsImpl) GetUserTransactions(userID uint) ([]dto.TransactionHistory, error) {
	var transactions []dto.TransactionHistory

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Product").Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
			return errors.New("transactions not found")
		}

		if len(transactions) == 0 {
			return errors.New("transactions not found")
		}

		return nil
	})

	return transactions, err
}

func (r *TransactionsImpl) GetAll() ([]entity.TransactionHistory, error) {
	var transactions []entity.TransactionHistory

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Product").Preload("User").Find(&transactions).Error; err != nil {
			return errors.New("transactions not found")
		}

		if len(transactions) == 0 {
			return errors.New("transactions are empty")
		}

		return nil
	})

	return transactions, err
}
