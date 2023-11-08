package dto

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email"     binding:"required,email"`
	Password string `json:"password"  binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type TopUpRequest struct {
	Balance int `json:"balance" binding:"number,min=0,max=100000000"`
}

type CreateCategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

type CreateProductRequest struct {
	Title      string `json:"title"           binding:"required"`
	Price      int    `json:"price,omitempty" binding:"omitempty,number,min=0,max=50000000"`
	Stock      int    `json:"stock,omitempty" binding:"omitempty,number,min=5,max=1000000"`
	CategoryID uint   `json:"category_id"     binding:"required"`
}

type UpdateProductRequest struct {
	Title      string `json:"title,omitempty"`
	Price      int    `json:"price,omitempty"       binding:"omitempty,number,min=0,max=50000000"`
	Stock      int    `json:"stock,omitempty"       binding:"omitempty,number,min=5,max=1000000"`
	CategoryID uint   `json:"category_id,omitempty"`
}
