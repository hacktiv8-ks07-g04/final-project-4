package dto

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email"     binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TopUpRequest struct {
	Balance int `json:"balance" valid:"type(int),range(0|100000000)~balance must be between 0 and 100.000.000"`
}
