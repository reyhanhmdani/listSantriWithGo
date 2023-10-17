package database

import (
	"errors"
	"gorm.io/gorm"
	"project1/model/entity"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

// //// ALL DATA
func (S *UserRepository) AllUsersData() ([]entity.Users, error) {
	var users []entity.Users
	if err := S.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

////// ALL DATA END

// //// CRUD
func (S UserRepository) CreateUserAdmin(users *entity.User) error {
	if err := S.DB.Create(users).Error; err != nil {
		return err
	}
	return nil
}
func (S *UserRepository) DeleteUser(userID int64) error {
	result := S.DB.Where("id", userID).Delete(&entity.User{})
	return result.Error
}

// Metode untuk membuat data Santri baru
func (S *UserRepository) UpdateUser(updateData *entity.UpdateUser, userID int64) error {
	// Pastikan hanya kolom-kolom yang perlu diupdate yang diperbarui
	updateColumns := make(map[string]interface{})

	if updateData.Username != "" {
		updateColumns["username"] = updateData.Username
	}

	if updateData.NewPassword != "" {
		updateColumns["password"] = updateData.NewPassword
	}

	// Gunakan GORM untuk melakukan pembaruan di database
	result := S.DB.Model(&entity.User{}).Where("id = ?", userID).Updates(updateColumns)

	// Periksa apakah ada kesalahan saat melakukan pembaruan
	if result.Error != nil {
		return result.Error
	}

	// Periksa apakah pengguna yang sesuai ditemukan untuk menghindari kasus error yang tidak diharapkan
	if result.RowsAffected == 0 {
		return errors.New("User not found")
	}

	return nil
}

// TOKEN
func (S *UserRepository) UpdateUserToken(user *entity.User) error {
	if err := S.DB.Model(user).Update("token", user.Token).Error; err != nil {
		return err
	}
	return nil
}

////// CRUD

// TAMBAHAN
func (S *UserRepository) CheckUsernameUser(username string) (*entity.User, error) {
	var admin entity.User
	result := S.DB.Where("username = ?", username).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &admin, nil

}
func (S *UserRepository) GetUserByID(userId int64) (*entity.User, error) {
	var user entity.User
	err := S.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (S *UserRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	if err := S.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (S *UserRepository) UpdateUserPasswordAndProfile(newUsername, newPassword string, userID int64) error {
	// Mulai transaksi database
	tx := S.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Perbarui username dan password pengguna
	if err := tx.Model(&entity.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"username": newUsername,
		"password": newPassword,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaksi
	tx.Commit()

	return nil
}

// //// PROFILE
func (S *UserRepository) GetSantriByUserID(userID int64) (*entity.CreateSantri, error) {
	var santri entity.CreateSantri
	if err := S.DB.Where("user_id = ?", userID).First(&santri).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Santri data not found for the user
		}
		return nil, err // Other database error
	}
	return &santri, nil
}

// JURUSAN, MINAT, STATUS
func (S *UserRepository) GetJurusanByID(id int) (*entity.Jurusan, error) {
	var jurusan entity.Jurusan
	if err := S.DB.First(&jurusan, id).Error; err != nil {
		return nil, err
	}
	return &jurusan, nil
}
func (S *UserRepository) GetMinatByID(id int) (*entity.Minat, error) {
	var minat entity.Minat
	if err := S.DB.First(&minat, id).Error; err != nil {
		return nil, err
	}
	return &minat, nil
}
func (S *UserRepository) GetStatusByID(id int) (*entity.Status, error) {
	var status entity.Status
	if err := S.DB.First(&status, id).Error; err != nil {
		return nil, err
	}
	return &status, nil
}

////// PROFILE END

// SEARCH
//func (S *UserRepository) SearchSantri(searchQuery string, page, perPage int) ([]entity.Santri, int64, error) {
//	var Santri []entity.Santri
//
//	var total int64
//	S.DB.Model(&entity.Santri{}).
//		Where("nama LIKE ? OR gender LIKE ? OR angkatan LIKE ? OR jurusan LIKE ? OR minat LIKE ? OR status LIKE ?",
//			"%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%",
//			"%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").
//		Count(&total)
//
//	// mengambil data dengan paginasi
//	offset := (page - 1) * perPage
//	err := S.DB.Where("nama LIKE ? OR gender LIKE ? OR angkatan LIKE ? OR jurusan LIKE ? OR minat LIKE ? OR status LIKE ?",
//		"%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%",
//		"%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").
//		Offset(offset).Limit(perPage).Find(&Santri).Error
//
//	return Santri, total, err
//}

// GET REFERENCE DATA
func (S *UserRepository) GetMinatList() ([]entity.Minat, error) {
	var minatList []entity.Minat
	if err := S.DB.Find(&minatList).Error; err != nil {
		return nil, err
	}
	return minatList, nil
}

func (S *UserRepository) GetJurusanList() ([]entity.Jurusan, error) {
	var jurusanList []entity.Jurusan
	if err := S.DB.Find(&jurusanList).Error; err != nil {
		return nil, err
	}
	return jurusanList, nil
}

func (S *UserRepository) GetStatusList() ([]entity.Status, error) {
	var statusList []entity.Status
	if err := S.DB.Find(&statusList).Error; err != nil {
		return nil, err
	}
	return statusList, nil
}