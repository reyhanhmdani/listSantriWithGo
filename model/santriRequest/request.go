package santriRequest

type CreateUsersRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	ID       int64  `json:"id"`
	Role     string `json:"role"`
}
