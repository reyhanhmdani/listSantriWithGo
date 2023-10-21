package entity

import "time"

type Attachments struct {
	ID              int64     `gorm:"primaryKey" json:"id"`
	UserID          int64     `gorm:"index" json:"user_id"`
	Path            string    `gorm:"type:varchar(255)" json:"path"`
	AttachmentOrder int64     `json:"attachment_order"`
	Timestamp       time.Time `gorm:"default:current_timestamp" json:"timestamp"`
}

func (Attachments) TableName() string {
	return "Attachments"
}
