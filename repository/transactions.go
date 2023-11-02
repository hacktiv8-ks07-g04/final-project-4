package repository

import (
	"gorm.io/gorm"
)

type TransactionsRepository interface{}

type TransactionsRepositoryImpl struct {
	db *gorm.DB
}

func TransactionsRepositoryInit(db *gorm.DB) *TransactionsRepositoryImpl {
	return &TransactionsRepositoryImpl{db}
}
