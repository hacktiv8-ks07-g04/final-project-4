package dto

type RegisterRequest struct {
	FullName string `json:"full_name" valid:"required~full name is required,type(string)"`
	Email    string `json:"email"     valid:"required~email is required,type(string),email~email must be a valid email address"`
	Password string `json:"password"  valid:"required~password is required,minstringlength(6)~password must be at least 6 characters long"`
}

type LoginRequest struct {
	Email    string `json:"email"    valid:"required~email is required,type(string),email~email must be a valid email address"`
	Password string `json:"password" valid:"required~password is required,minstringlength(6)~password must be at least 6 characters long"`
}

type TopUpRequest struct {
	Balance int `json:"balance" valid:"type(int),range(0|100000000)~balance must be between 0 and 100.000.000"`
}

type CreateCategoryRequest struct {
	Type string `json:"type" valid:"required~type is required,type(string)"`
}

type CreateProductRequest struct {
	Title      string `json:"title"       valid:"required~title is required,type(string)"`
	Price      int    `json:"price"       valid:"required~price is required,type(int),range(0|50000000)~price must be between 0 and 50.000.000"`
	Stock      int    `json:"stock"       valid:"required~stock is required,type(int),range(5|1000000)~stock must be at least 5"`
	CategoryID uint   `json:"category_id"`
}

type UpdateProductRequest struct {
	Title      string `json:"title,omitempty"`
	Price      int    `json:"price,omitempty"       binding:"omitempty,number,min=0,max=50000000"`
	Stock      int    `json:"stock,omitempty"       binding:"omitempty,number,min=5,max=1000000"`
	CategoryID uint   `json:"category_id,omitempty"`
}
