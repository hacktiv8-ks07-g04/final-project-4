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
