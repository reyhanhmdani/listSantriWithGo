package entity

import "time"

type TokenReset struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	UserID    int64     `json:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (TokenReset) TableName() string {
	return "TokenReset"
}

type SendEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// PasswordResetRequest digunakan untuk mengganti kata sandi
type PasswordResetRequest struct {
	Token           string `json:"token"`
	NewPassword     string `json:"new_password" binding:"required"`
	ReenterPassword string `json:"reenter_password" binding:"required"`
}
