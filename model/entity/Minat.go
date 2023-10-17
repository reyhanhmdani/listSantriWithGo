package entity

type Minat struct {
	ID    int64  `gorm:"primaryKey" json:"id"`
	Minat string `gorm:"type:varchar(255)" json:"minat"`
}

func (Minat) TableName() string {
	return "Minat"
}
