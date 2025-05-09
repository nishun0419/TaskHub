package customer

type RegisterInput struct {
	Username string `json:"username" binding:"required,min=3,max=10"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}
