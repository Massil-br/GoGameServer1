package auth

type CreateUserRequest struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword	string `json:"confirmPassword"`
}

type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"password"`
}