package santriRequest

type CreateUsersRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `gorm:"type:varchar(255);uniqueIndex:idx_email_name" json:"email"`
	Password string `json:"password" binding:"required"`
	ID       int64  `json:"id"`
	Role     string `json:"role"`
}
