package repository

import (
	"mime/multipart"
	"project1/model/entity"
)

type SantriRepository interface {
	// ALL DATA
	AllSantriData() ([]entity.Santri, error)
	// CRUD
	CreateSantri(createSantriData *entity.CreateSantri) (*entity.CreateSantri, error)
	DeleteSantri(userID int64) error
	// tambahan untuk createSantri
	GetJurusanIDByName(jurusanName string) (int64, error)
	GetMinatIDByName(minatName string) (int64, error)
	GetStatusIDByName(statusName string) (int64, error)

	// uploadFile
	UploadFileLocal(file *multipart.FileHeader, Id int64) (*entity.Attachments, error)
	UpdateToAtch(santri *entity.Santri) error

	// TAMBAHAN
	GetSantriByID(santriID int64) (*entity.Santri, error)

	// SEARCH
	SearchSantri(searchQuery string, page, perPage int) ([]entity.Santri, int64, error)
}
