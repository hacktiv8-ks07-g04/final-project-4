package repository

import (
	"gorm.io/gorm"
)

type Transactions interface{}

type TransactionsImpl struct {
	db *gorm.DB
}

func InitTransactions(db *gorm.DB) *TransactionsImpl {
	return &TransactionsImpl{db}
}
