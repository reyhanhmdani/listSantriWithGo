package entity

type Jurusan struct {
	ID      int64  `gorm:"primaryKey" json:"id"`
	Jurusan string `gorm:"type:varchar(255)" json:"jurusan"`
}

func (Jurusan) TableName() string {
	return "Jurusan"
}
