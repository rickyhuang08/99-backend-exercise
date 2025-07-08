package entity

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type UserPayload struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
}

type GetUserParam struct {
	ID       int    `form:"id"`
	Name     string `form:"name"`
	Email    string `form:"email"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}
