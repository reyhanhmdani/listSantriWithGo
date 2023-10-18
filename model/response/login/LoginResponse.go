package responseForSantri

type LoginResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
	Token   string `json:"token"`
	Admin   bool   `json:"-"`
}

type UserLogin struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
	Token    string `json:"token"`
}
