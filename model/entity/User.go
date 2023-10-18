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

type ForgotPassword struct {
	Email    string `json:"email" binding:"required,email"`
	ResetURL string `json:"reset_url"`
}

type ResetPasswordRequest struct {
	ResetToken      string `form:"resetToken"`
	NewPassword     string `form:"newPassword"`
	ConfirmPassword string `form:"confirmPassword"`
}

type ResetPasswordToken struct {
	ID             int64     `gorm:"primaryKey" json:"id"`
	UserID         int64     `json:"user_id"`
	ResetToken     string    `json:"reset_token"`
	ExpirationTime time.Time `json:"expiration_time"`
}
