package repository

import "project1/model/entity"

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

	// TAMBAHAN
	GetSantriByID(santriID int64) (*entity.Santri, error)

	// SEARCH
	SearchSantri(searchQuery string, page, perPage int) ([]entity.Santri, int64, error)
}
