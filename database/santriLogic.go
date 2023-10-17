package database

import (
	"errors"
	"gorm.io/gorm"
	"project1/model/entity"
)

type SantriRepository struct {
	DB *gorm.DB
}

func NewSantriRepository(DB *gorm.DB) *SantriRepository {
	return &SantriRepository{
		DB: DB,
	}
}

// //// ALL DATA
func (S *SantriRepository) AllSantriData() ([]entity.Santri, error) {
	var santris []entity.Santri
	if err := S.DB.
		Table("Santri").
		Select("Santri.*, Jurusan.Jurusan, Minat.Minat, Status.Status").
		Joins("LEFT JOIN Jurusan ON Santri.jurusan = Jurusan.id").
		Joins("LEFT JOIN Minat ON Santri.minat = Minat.id").
		Joins("LEFT JOIN Status ON Santri.status = Status.id").
		Find(&santris).Error; err != nil {
		return nil, err
	}
	return santris, nil
}

////// ALL DATA END

// //// CRUD

// Metode untuk membuat data Santri baru
func (S *SantriRepository) CreateSantri(createSantriData *entity.CreateSantri) (*entity.CreateSantri, error) {
	result := S.DB.Create(createSantriData)
	return createSantriData, result.Error
}
func (S *SantriRepository) DeleteSantri(santriID int64) error {
	result := S.DB.Where("id", santriID).Delete(&entity.Santri{})
	return result.Error
}

// JURUSA, MINAT, STATUS
// GetJurusanIDByName mengambil ID Jurusan berdasarkan nama Jurusan.
func (S *SantriRepository) GetJurusanIDByName(jurusanName string) (int64, error) {
	var jurusanID int64
	err := S.DB.Model(&entity.Jurusan{}).
		Where("jurusan = ?", jurusanName).
		Pluck("id", &jurusanID).
		Error
	return jurusanID, err
}

// GetMinatIDByName mengambil ID Minat berdasarkan nama Minat.
func (S *SantriRepository) GetMinatIDByName(minatName string) (int64, error) {
	var minatID int64
	err := S.DB.Model(&entity.Minat{}).
		Where("minat = ?", minatName).
		Pluck("id", &minatID).
		Error
	return minatID, err
}

// GetStatusIDByName mengambil ID Status berdasarkan nama Status.
func (S *SantriRepository) GetStatusIDByName(statusName string) (int64, error) {
	var statusID int64
	err := S.DB.Model(&entity.Status{}).
		Where("status = ?", statusName).
		Pluck("id", &statusID).
		Error
	return statusID, err
}

////// CRUD

// TAMBAHAN
func (S *SantriRepository) GetSantriByID(santriID int64) (*entity.Santri, error) {
	var santri entity.Santri
	if err := S.DB.Where("id = ?", santriID).First(&santri).Error; err != nil {
		// Tangani kasus jika data Santri tidak ditemukan
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &santri, nil
}

func (S *SantriRepository) SearchSantri(searchQuery string, page, perPage int) ([]entity.Santri, int64, error) {
	var Santri []entity.Santri

	var total int64
	db := S.DB.Model(&entity.Santri{}).
		Joins("JOIN Jurusan ON Santri.jurusan = Jurusan.id").
		Joins("JOIN Minat ON Santri.minat = Minat.id").
		Joins("JOIN Status ON Santri.status = Status.id").
		Where("nama LIKE ? OR gender LIKE ? OR angkatan LIKE ? OR Jurusan.jurusan LIKE ? OR Minat.minat LIKE ? OR Status.status LIKE ?",
			"%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%",
			"%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").
		Count(&total)

	// mengambil data dengan paginasi
	offset := (page - 1) * perPage
	err := db.
		Select("Santri.*").
		Offset(offset).
		Limit(perPage).
		Find(&Santri).Error

	return Santri, total, err
}
