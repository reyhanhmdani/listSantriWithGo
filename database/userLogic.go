package database

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"project1/config"
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

func (S *UserRepository) UpdateUserForAdmin(updateSantri *entity.UpdateUserForAdmin, userID int64) error {
	// Pastikan hanya kolom-kolom yang perlu diupdate yang diperbarui
	updateColumns := make(map[string]interface{})

	if updateSantri.Username != "" {
		updateColumns["username"] = updateSantri.Username
	}

	if updateSantri.NewPassword != "" {
		updateColumns["password"] = updateSantri.NewPassword
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
func (S *UserRepository) UpdateUserOnlyPassword(updateData *entity.UpdatePasswordUser, userID int64) error {
	// Pastikan hanya kolom-kolom yang perlu diupdate yang diperbarui
	updateColumns := make(map[string]interface{})

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
func (S *UserRepository) UpdateUserOnlyUsername(updateUsername *entity.UpdateUsernameUser, userID int64) error {
	// Pastikan hanya kolom-kolom yang perlu diupdate yang diperbarui
	updateColumns := make(map[string]interface{})

	if updateUsername.Username != "" {
		updateColumns["username"] = updateUsername.Username
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
func (S *UserRepository) CheckUsernameUserOremail(username, email string) (*entity.User, error) {
	var admin entity.User
	result := S.DB.Where("username = ? OR email = ?", username, email).First(&admin)
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

// FORGOT PASSWORD

func (S *UserRepository) GeneratePasswordResetToken(userID int64) (string, error) {
	// Generate a unique token (e.g., a random string)
	token, _ := config.GenerateUniqueToken()

	// Simpan token reset password ke dalam database
	passwordReset := &entity.TokenReset{
		UserID: userID,
		Token:  token,
	}
	if err := S.DB.Create(passwordReset).Error; err != nil {
		return "", err
	}

	// Hapus token reset password yang sudah digunakan (jika ada)
	S.DB.Where("UserID = ?", userID).Delete(&entity.TokenReset{})

	return token, nil
}

// VerifyPasswordResetToken memeriksa apakah token reset password valid
func (S *UserRepository) VerifyPasswordResetToken(token string) (int64, error) {
	var passwordReset entity.TokenReset
	if err := S.DB.Where("token = ?", token).First(&passwordReset).Error; err != nil {
		return 0, err
	}
	return passwordReset.UserID, nil
}

// DeletePasswordResetToken menghapus token reset password yang sudah digunakan
func (S *UserRepository) DeletePasswordResetToken(token string) error {
	return S.DB.Where("token = ?", token).Delete(&entity.TokenReset{}).Error
}

func (S *UserRepository) UpdatePassword(userID int64, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Lakukan pembaruan password dalam database
	if err := S.DB.Model(&entity.User{}).Where("id = ?", userID).Update("password", string(hashedPassword)).Error; err != nil {
		return err
	}

	return nil
}

///////////////////////////
