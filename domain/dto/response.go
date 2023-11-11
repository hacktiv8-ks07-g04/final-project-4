package dto

import (
	"time"
)

type RegisterResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type TopUpResponse struct {
	Message string `json:"message"`
}

type CreateCategoryResponse struct {
	ID                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

type GetCategoryResponse struct {
	ID                uint                 `json:"id"`
	Type              string               `json:"type"`
	SoldProductAmount int                  `json:"sold_product_amount"`
	CreatedAt         time.Time            `json:"created_at"`
	UpdatedAt         time.Time            `json:"updated_at"`
	Products          []GetProductResponse `json:"products"`
}

type DeleteCategoryResponse struct {
	Message string `json:"message"`
}

type GetProductResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type UpdateProductResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type TransactionBill struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type CreateTransactionResponse struct {
	Message         string          `json:"message"`
	TransactionBill TransactionBill `json:"transaction_bill"`
}

type User struct {
	ID        uint   `json:"id"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Product struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID uint   `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type TransactionHistory struct {
	ID         uint    `json:"id"`
	ProductID  uint    `json:"product_id"`
	UserID     uint    `json:"user_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice int     `json:"total_price"`
	Product    Product `json:"product"`
}

type GetAllTransactionsResponse struct {
	ID         uint    `json:"id"`
	ProductID  uint    `json:"product_id"`
	UserID     uint    `json:"user_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice int     `json:"total_price"`
	Product    Product `json:"product"`
	User       User    `json:"user"`
}
