package entity

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token"`
}