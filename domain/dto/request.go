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
