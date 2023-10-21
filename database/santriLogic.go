package database

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"project1/model/entity"
	"time"
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
		Preload("Attachments").
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

// //// ALL DATA END
func (S *SantriRepository) UploadFileLocal(file *multipart.FileHeader, Id int64) (*entity.Attachments, error) {
	santri := &entity.Santri{}
	if err := S.DB.Where("id = ?", Id).First(santri).Error; err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// bikin nama file yang uniq untuk menghindari konflik
	uniqueFilename := fmt.Sprintf("%s%s", uuid.NewString(), filepath.Ext(file.Filename))

	// kirimkan upload file nya ke folder local
	uploadDir := "uploads"
	// buat direktori nya kalau beum ada
	err = os.MkdirAll(uploadDir, 0755)
	if err != nil {
		return nil, err
	}

	// buat file tujuan
	dest, err := os.Create(filepath.Join(uploadDir, uniqueFilename))
	if err != nil {
		return nil, err
	}
	defer dest.Close()

	// Copy file nya ke file tujuan
	_, err = io.Copy(dest, src)
	if err != nil {
		return nil, err
	}
	// Return the local file path
	localFilePath := filepath.Join(uploadDir, uniqueFilename)

	var attachmentOrder int64 = 1 // Set the initial attachment_order
	// Get the count of existing attachments for the data santri
	existingAttachmentCount := int64(0)
	attachmentOrder = existingAttachmentCount + 1 // Set attachment_order dynamically

	// Membuat catatan lampiran di database
	attachment := &entity.Attachments{
		UserID:          Id,
		Path:            localFilePath,
		AttachmentOrder: attachmentOrder, // atur order
		Timestamp:       time.Now(),
	}
	err = S.DB.Create(attachment).Error
	if err != nil {
		return nil, err
	}

	return attachment, nil
}

func (S *SantriRepository) UpdateToAtch(santri *entity.Santri) error {
	err := S.DB.Save(santri).Error
	return err
}

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

// Upload

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
