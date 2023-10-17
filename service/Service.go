package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"project1/config"
	"project1/model/entity"
	"project1/model/respErr"
	"project1/model/response"
	responseForSantri "project1/model/response/login"
	"project1/model/santriRequest"
	"project1/repository"
	"strconv"
)

type Handler struct {
	SantriRepository repository.SantriRepository
	UserRepository   repository.UserRepository
}

func NewSantriService(santriRepo repository.SantriRepository, userRepo repository.UserRepository) *Handler {
	return &Handler{
		SantriRepository: santriRepo,
		UserRepository:   userRepo,
	}
}

// ALL DATA
// ADMIN
func (h *Handler) ViewAllUsers(ctx *gin.Context) {
	users, err := h.UserRepository.AllUsersData()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
	}
	ctx.JSON(http.StatusOK, users)
}

// USER
func (h *Handler) Register(ctx *gin.Context) {
	request := new(santriRequest.CreateUsersRequest)

	// binding request body ke struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	// cek apakah username sudah ada di database
	existingUser, err := h.UserRepository.CheckUsernameUser(request.Username)
	if existingUser != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: "Username, or email already exist",
			Status:  http.StatusBadRequest,
		})
		return
	}
	// hash password pengguna sebelum disimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Failed to hash Password",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	//role := request.Role
	//if role == "" {
	//	role = "user"
	//}

	// simpan Pengguna ke database
	newSantri := &entity.User{
		Username: request.Username,
		Password: string(hashedPassword),
		Role:     request.Role,
	}

	err = h.UserRepository.CreateUserAdmin(newSantri)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		logrus.Error("gagal membuat")
		return
	}

	// mengembalikan pesan berhasil sebagai response
	ctx.JSON(http.StatusOK, gin.H{"message": "Santri created successfully"})
}
func (h *Handler) UpdateUserProfile(ctx *gin.Context) {
	// Dapatkan ID pengguna dari URL
	userID, _ := ctx.Params.Get("id")
	// Dapatkan peran pengguna dari token
	// Pastikan ID yang diterima adalah angka
	userIDInt64, err := strconv.ParseInt(userID, 10, 64)
	if err != nil || userIDInt64 == 0 {
		ctx.JSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: "Invalid user ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Ambil data user yang sedang login
	currentUserID, _ := ctx.Get("user_id")
	if currentUserID == nil {
		ctx.JSON(http.StatusUnauthorized, respErr.ErrorResponse{
			Message: "User not authenticated",
			Status:  http.StatusUnauthorized,
		})
		return
	}

	// Cast currentUserID ke int64
	currentUserIDInt64, ok := currentUserID.(int64)
	if !ok {
		ctx.JSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: "Invalid user_id",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Pastikan pengguna hanya dapat mengubah profil mereka sendiri
	if currentUserIDInt64 != userIDInt64 {
		ctx.JSON(http.StatusUnauthorized, respErr.ErrorResponse{
			Message: "User is not authorized to update this profile",
			Status:  http.StatusUnauthorized,
		})
		return
	}
	// Ambil data yang ingin diupdate dari JSON request
	updateData := new(entity.UpdateUser)
	if err := ctx.ShouldBindJSON(updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Ambil data pengguna yang sedang login dari database
	currentUser, err := h.UserRepository.GetUserByID(currentUserIDInt64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	// Validasi old password jika ada
	if updateData.OldPassword != "" {
		// Lakukan validasi old password dengan password yang ada di database
		passwordMatchErr := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(updateData.OldPassword))
		if passwordMatchErr != nil {
			ctx.JSON(http.StatusUnauthorized, respErr.ErrorResponse{
				Message: "Old password is incorrect",
				Status:  http.StatusUnauthorized,
			})
			return
		}
	}

	// Jika username atau password dalam JSON request kosong, gunakan data yang sudah ada di profil pengguna
	if updateData.Username == "" {
		updateData.Username = currentUser.Username
	}
	if updateData.NewPassword == "" {
		updateData.NewPassword = currentUser.Password
	}

	// Hash password baru jika ada
	if updateData.NewPassword != currentUser.Password {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
				Message: "Internal Server Error",
				Status:  http.StatusInternalServerError,
			})
			return
		}
		updateData.NewPassword = string(hashedPassword)
	}

	// Lakukan validasi dan update profil user di sini
	if err := h.UserRepository.UpdateUser(updateData, userIDInt64); err != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.UpdateResponse{
		Status:  http.StatusOK,
		Message: "User profile updated successfully",
		Data: entity.UpdateUser{
			ID:          userIDInt64,
			Username:    updateData.Username,
			OldPassword: updateData.OldPassword,
			NewPassword: updateData.NewPassword,
		},
	})
}
func (h *Handler) UpdateUserForAdmin(ctx *gin.Context) {
	// pengecekan role nya sudah di bagian middleware
	userID, _ := ctx.Params.Get("id")
	// Dapatkan peran pengguna dari token
	// Pastikan ID yang diterima adalah angka
	userIDInt64, err := strconv.ParseInt(userID, 10, 64)
	if err != nil || userIDInt64 == 0 {
		ctx.JSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: "Invalid user ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// dapatkan peran pengguna yang akan di perbarui dari database
	userToUpdate, err := h.UserRepository.GetUserByID(userIDInt64)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}

	// pastikan hanya user yang bisa di update oleh admin
	// Pastikan pengguna yang akan diperbarui memiliki peran "user"
	if userToUpdate.Role != "user" {
		ctx.JSON(http.StatusUnauthorized, respErr.ErrorResponse{
			Message: "Admin can only update users, not other admins",
			Status:  http.StatusUnauthorized,
		})
		return
	}

	// Dapatkan data pembaruan dari JSON request
	updateData := new(entity.UpdateUser)
	if err := ctx.ShouldBindJSON(updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Ambil data pengguna yang sedang login dari database
	User, err := h.UserRepository.GetUserByID(userIDInt64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	// Validasi old password jika ada
	if updateData.OldPassword != "" {
		// Lakukan validasi old password dengan password yang ada di database
		passwordMatchErr := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(updateData.OldPassword))
		if passwordMatchErr != nil {
			ctx.JSON(http.StatusUnauthorized, respErr.ErrorResponse{
				Message: "Old password is incorrect",
				Status:  http.StatusUnauthorized,
			})
			return
		}
	}

	// Jika username atau password dalam JSON request kosong, gunakan data yang sudah ada di profil pengguna
	if updateData.Username == "" {
		updateData.Username = User.Username
	}
	if updateData.NewPassword == "" {
		updateData.NewPassword = User.Password
	}

	// Hash password baru jika ada
	if updateData.NewPassword != User.Password {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
				Message: "Internal Server Error",
				Status:  http.StatusInternalServerError,
			})
			return
		}
		updateData.NewPassword = string(hashedPassword)
	}

	// Lakukan validasi dan update profil user di sini
	if err := h.UserRepository.UpdateUser(updateData, userIDInt64); err != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.UpdateResponse{
		Status:  http.StatusOK,
		Message: "User profile updated successfully",
		Data: entity.UpdateUser{
			ID:          userIDInt64,
			Username:    updateData.Username,
			OldPassword: updateData.OldPassword,
			NewPassword: updateData.NewPassword,
		},
	})

}
func (h *Handler) DeleteUser(ctx *gin.Context) {
	// Dapatkan ID pengguna yang ingin dihapus
	userIDToDelete, _ := ctx.Params.Get("id")
	userIDInt64, err := strconv.ParseInt(userIDToDelete, 10, 64)
	if err != nil || userIDInt64 == 0 {
		ctx.JSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: "Invalid user ID",
			Status:  http.StatusBadRequest,
		})
		return
	}

	currentUserID, _ := ctx.Get("user_id")
	if currentUserID == nil {
		ctx.JSON(http.StatusUnauthorized, respErr.ErrorResponse{
			Message: "User not authenticated",
			Status:  http.StatusUnauthorized,
		})
		return
	}

	// Pastikan pengguna hanya dapat menghapus akun mereka sendiri
	currentUserIDInt64, ok := currentUserID.(int64)
	if !ok || currentUserIDInt64 != userIDInt64 {
		ctx.JSON(http.StatusUnauthorized, respErr.ErrorResponse{
			Message: "User is not authorized to delete this account",
			Status:  http.StatusUnauthorized,
		})
		return
	}

	// Dapatkan data pengguna yang akan dihapus
	userToDelete, err := h.UserRepository.GetUserByID(userIDInt64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	// Hapus data santri yang terkait jika ada
	if userToDelete.ID != 0 {
		if err := h.SantriRepository.DeleteSantri(userToDelete.ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
				Message: "Failed to delete associated santri data",
				Status:  http.StatusInternalServerError,
			})
			return
		}
	}
	// Hapus akun pengguna
	if err := h.UserRepository.DeleteUser(userIDInt64); err != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Failed to delete user account",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessMessage{
		Status:  http.StatusOK,
		Message: "User account and associated santri data deleted successfully",
	})
}

func (h *Handler) Login(ctx *gin.Context) {
	var userLogin responseForSantri.UserLogin

	// binding request body ke struct UserLogin
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: "Invalid request Body",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Cek apakah pengguna ada di database berdasarkan username atau email
	storedUser, err := h.UserRepository.CheckUsernameUser(userLogin.Username)
	if err != nil || storedUser == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, respErr.ErrorResponse{
			Message: "Invalid Username or Password",
			Status:  http.StatusUnauthorized,
		})
		return
	}

	// Membandingkan password yang dimasukkan dengan hash password di database
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(userLogin.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, respErr.ErrorResponse{
			Message: "Invalid Username or Password",
			Status:  http.StatusUnauthorized,
		})
		return
	}

	// Menentukan apakah pengguna adalah admin
	isAdmin := storedUser.Role == "admin"

	// Membuat token
	token, err := config.CreateJWTToken(storedUser.Username, storedUser.ID, storedUser.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Failed to generate Token",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	// Cetak token untuk memeriksanya
	//fmt.Println("Generated Token:", token)
	// Simpan token ke dalam field "token" di tabel database
	storedUser.Token = token
	if err := h.UserRepository.UpdateUserToken(storedUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Failed to update token",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	rsp := responseForSantri.LoginResponse{
		ID: storedUser.ID,
		Message: fmt.Sprintf("Hello %s! You are%s logged in.", userLogin.Username, func() string {
			if isAdmin {
				return " an admin"
			}
			return " user"
		}()),
		Token: token,
	}

	ctx.JSON(http.StatusOK, rsp)
}

// PROFILE
func (h *Handler) GetProfile(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, respErr.ErrorResponse{
			Message: "User not authenticated",
			Status:  http.StatusUnauthorized,
		})
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: "invalid user_id",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Mengambil data Santri terhubung ke pengguna berdasarkan ID pengguna
	santri, err := h.UserRepository.GetSantriByUserID(userIDInt64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	var jurusan, minat, status string

	if santri.Jurusan != nil {
		// Mengambil rincian jurusan jika ada
		jurusanDetail, err := h.UserRepository.GetJurusanByID(*santri.Jurusan)
		if err == nil {
			jurusan = jurusanDetail.Jurusan
		}
	}

	if santri.Minat != nil {
		// Mengambil rincian minat jika ada
		minatDetail, err := h.UserRepository.GetMinatByID(*santri.Minat)
		if err == nil {
			minat = minatDetail.Minat
		}
	}

	if santri.Status != nil {
		// Mengambil rincian status jika ada
		statusDetail, err := h.UserRepository.GetStatusByID(*santri.Status)
		if err == nil {
			status = statusDetail.Status
		}
	}

	// Membuat respons dengan urutan yang diinginkan
	profileResponse := response.SantriProfileResponse{
		Nama:     santri.Nama,
		UserID:   santri.UserID,
		HP:       santri.HP,
		Email:    santri.Email,
		Gender:   santri.Gender,
		Alamat:   santri.Alamat,
		Angkatan: santri.Angkatan,
		Jurusan:  jurusan,
		Minat:    minat,
		Status:   status,
	}

	ctx.JSON(http.StatusOK, profileResponse)
}

func (h *Handler) GetSantriByID(ctx *gin.Context) {
	santriId := ctx.Param("id")
	Id, err := strconv.ParseInt(santriId, 10, 64)
	if err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: "Bad request",
			Status:  http.StatusBadRequest,
		})
		return
	}
	listSantri, err := h.SantriRepository.GetSantriByID(Id)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusNotFound, respErr.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		})
		return
	}

	logrus.Info(http.StatusOK, " Success Get By ID")
	ctx.JSON(http.StatusOK, response.GetSantriByID{
		Status:  http.StatusOK,
		Message: "Success Get Id",
		Data:    *listSantri,
	})

}

// PROFILE END

// SEARCH
func (h *Handler) SearchHandler(ctx *gin.Context) {

	searchQuery := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "5"))

	santri, total, err := h.SantriRepository.SearchSantri(searchQuery, page, perPage)
	if err != nil {
		logrus.Errorf("Error during santri search: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SearchResponse{
		Status:  http.StatusOK,
		Data:    santri,
		Total:   total,
		Page:    page,
		PerPage: perPage,
	})
}

// GET REFERENCES DATA
func (h *Handler) GetReferenceData(ctx *gin.Context) {
	minatList, errMinat := h.UserRepository.GetMinatList()
	jurusanList, errJurusan := h.UserRepository.GetJurusanList()
	statusList, errStatus := h.UserRepository.GetStatusList()

	if errMinat != nil || errJurusan != nil || errStatus != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	referenceData := map[string]interface{}{
		"minat":   minatList,
		"jurusan": jurusanList,
		"status":  statusList,
	}

	ctx.JSON(http.StatusOK, referenceData)
}

func (h *Handler) GetJurusanList(ctx *gin.Context) {
	jurusanList, errJurusan := h.UserRepository.GetJurusanList()

	if errJurusan != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, jurusanList)
}

func (h *Handler) GetMinatList(ctx *gin.Context) {
	minatList, errMinat := h.UserRepository.GetMinatList()

	if errMinat != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, minatList)
}

func (h *Handler) GetStatusList(ctx *gin.Context) {
	statusList, errStatus := h.UserRepository.GetStatusList()

	if errStatus != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, statusList)
}

// SANTRI
func (h *Handler) ViewAllSantri(ctx *gin.Context) {
	// Dapatkan data pengguna dari basis data
	santris, err := h.SantriRepository.AllSantriData()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	// Tampilkan data pengguna dalam format respons JSON
	ctx.JSON(http.StatusOK, santris)
}

func (h *Handler) CreateSantri(ctx *gin.Context) {
	santri := new(entity.CreateSantri)
	if err := ctx.ShouldBindJSON(santri); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Get the user ID from the token
	userID, _ := ctx.Get("user_id")
	role, _ := ctx.Get("role")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, respErr.ErrorResponse{
			Message: "User not authenticated",
			Status:  http.StatusUnauthorized,
		})
		return
	}

	// Cast the userID to int64
	userIDInt64, ok := userID.(int64)
	if !ok {
		ctx.JSON(http.StatusBadRequest, respErr.ErrorResponse{
			Message: "Invalid user_id",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// Check if the user is admin
	isAdmin := role == "admin"

	// If the user is not admin, check if Santri data already exists
	if !isAdmin {
		existingSantri, err := h.UserRepository.GetSantriByUserID(userIDInt64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
				Message: "Internal Server Error",
				Status:  http.StatusInternalServerError,
			})
			return
		}
		if existingSantri != nil {
			ctx.JSON(http.StatusBadRequest, respErr.ErrorResponse{
				Message: "Santri data for this user already exists",
				Status:  http.StatusBadRequest,
			})
			return
		}
	}

	santri.UserID = userIDInt64

	newSantri := &entity.CreateSantri{
		UserID:   santri.UserID,
		Nama:     santri.Nama,
		HP:       santri.HP,
		Email:    santri.Email,
		Gender:   santri.Gender,
		Alamat:   santri.Alamat,
		Angkatan: santri.Angkatan,
		Jurusan:  santri.Jurusan,
		Minat:    santri.Minat,
		Status:   santri.Status,
	}

	createdSantri, errCreate := h.SantriRepository.CreateSantri(newSantri)
	if errCreate != nil {
		ctx.JSON(http.StatusInternalServerError, respErr.ErrorResponse{
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ListResponse{
		Status:  http.StatusOK,
		Message: "New Character Created",
		Data:    *createdSantri,
	})
}
