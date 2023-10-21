package entity

type Santri struct {
	ID          int64         `json:"id"`
	Nama        string        `json:"nama"`
	HP          int64         `json:"hp"`
	Email       string        `json:"email"`
	Gender      string        `json:"gender"`
	Alamat      string        `json:"alamat"`
	Angkatan    int           `json:"angkatan"`
	Jurusan     string        `json:"jurusan"`
	Minat       string        `json:"minat"`
	Status      string        `json:"status"`
	Attachments []Attachments `gorm:"foreignKey:user_id" json:"attachments"`
}

func (Santri) TableName() string {
	return "Santri"
}

// Create Santri

type ForCreateSantri struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Nama     string `json:"nama" binding:"required"`
	HP       int64  `json:"hp"`
	Email    string `json:"email"`
	Gender   string `json:"gender" binding:"required"`
	Alamat   string `json:"alamat" binding:"required"`
	Angkatan int    `json:"angkatan"`
	Jurusan  int    `json:"jurusan"`
	Minat    int    `json:"minat"`
	Status   int    `json:"status"`
}

type CreateSantri struct {
	Nama     string `json:"nama"`
	UserID   int64  `json:"user_id"`
	HP       int64  `json:"hp"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Alamat   string `json:"alamat"`
	Angkatan int    `json:"angkatan"`
	Jurusan  *int   `json:"jurusan"`
	Minat    *int   `json:"minat"`
	Status   *int   `json:"status"`
}

func (CreateSantri) TableName() string {
	return "Santri"
}

// UPDATE

//type Gender string
//
//const (
//	Pria   Gender = "Pria"
//	Wanita Gender = "Wanita"
//)
