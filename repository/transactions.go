package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/hacktiv8-ks07-g04/final-project-4/domain/dto"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
)

type Transactions interface {
	Create(userID uint, req dto.CreateTransactionRequest) (dto.TransactionBill, error)
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
		transactionHistory entity.TransactionHistory
		totalPrice         int
		transactionBill    dto.TransactionBill
	)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&product, req.ProductID).Error; err != nil {
			return errors.New("product not found")
		}

		totalPrice = product.Price * req.Quantity

		transactionHistory = entity.TransactionHistory{
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
