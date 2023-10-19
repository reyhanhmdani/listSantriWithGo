package entity

import "time"

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null;unique" json:"username"`
	Email    string `gorm:"type:varchar(255);uniqueIndex:idx_email_name" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Token    string `gorm:"type:varchar(255)" json:"token"`
	Role     string `gorm:"default:user" json:"role"`
}

func (User) TableName() string {
	return "User"
}

type Users struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null;unique" json:"username"`
	Email    string `gorm:"type:varchar(255);uniqueIndex:idx_email_name" json:"email"`
	Role     string `json:"role"`
}

func (Users) TableName() string {
	return "User"
}

// Update USERS

type UpdatePasswordUser struct {
	ID          int64  `json:"id"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
type UpdateUsernameUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
}

type UpdateUserForAdmin struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UpdateEmail struct {
	ID    int64  `json:"id"`
	Email string `gorm:"type:varchar(255);uniqueIndex:idx_email_name" json:"email"`
}

type TokenReset struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	UserID    int64     `json:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

func (TokenReset) TableName() string {
	return "PasswordReset"
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
