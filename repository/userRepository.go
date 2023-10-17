package repository

import "project1/model/entity"

type UserRepository interface {
	// ALL DATA
	AllUsersData() ([]entity.Users, error)
	//AllSantriData() ([]entity.Santri, error)
	// CRUD
	CreateUserAdmin(users *entity.User) error
	DeleteUser(userID int64) error
	//CreateSantri(createSantriData *entity.CreateSantri) (*entity.CreateSantri, error)
	UpdateUser(updateData *entity.UpdateUser, userID int64) error

	// TOKEN
	UpdateUserToken(user *entity.User) error
	// tambahan untuk createSantri
	//GetJurusanIDByName(jurusanName string) (int64, error)
	//GetMinatIDByName(minatName string) (int64, error)
	//GetStatusIDByName(statusName string) (int64, error)

	// TAMBAHAN
	CheckUsernameUser(username string) (*entity.User, error)
	//GetSantriByID(santriID int64) (*entity.Santri, error)
	GetUserByID(userId int64) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	UpdateUserPasswordAndProfile(newUsername, newPassword string, userID int64) error

	// PROFILE
	GetSantriByUserID(userID int64) (*entity.CreateSantri, error)

	GetJurusanByID(id int) (*entity.Jurusan, error)
	GetMinatByID(id int) (*entity.Minat, error)
	GetStatusByID(id int) (*entity.Status, error)

	// GET REFERENCE DATA
	GetMinatList() ([]entity.Minat, error)
	GetJurusanList() ([]entity.Jurusan, error)
	GetStatusList() ([]entity.Status, error)
}
