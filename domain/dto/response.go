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
