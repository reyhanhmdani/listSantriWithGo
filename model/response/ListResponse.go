package response

import "project1/model/entity"

//import "ListCharacterGI/model/entity"

type ListResponse struct {
	Status  interface{}         `json:"status"`
	Message interface{}         `json:"message"`
	Data    entity.CreateSantri `json:"data"`
}

type GetSantriByID struct {
	Status  interface{}   `json:"status"`
	Message interface{}   `json:"message"`
	Data    entity.Santri `json:"data"`
}

type ListResponseForUpdate struct {
	Status  interface{} `json:"status"`
	Message interface{} `json:"message"`
	Data    entity.User `json:"data"`
}

type ResponseToGetAll struct {
	Message string `json:"message"`
	Data    int    `json:"data"`
	//MHS     []entity.Characters `json:"todos"`
}

type IDResponse struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type DeleteResponse struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

type UpdateResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"data"`
	Data    interface{} `json:"todos"`
}

type SantriProfileResponse struct {
	Nama     string `json:"nama"`
	UserID   int64  `json:"user_id"`
	HP       int64  `json:"hp"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Alamat   string `json:"alamat"`
	Angkatan int    `json:"angkatan"`
	Jurusan  string `json:"jurusan"`
	Minat    string `json:"minat"`
	Status   string `json:"status"`
}
