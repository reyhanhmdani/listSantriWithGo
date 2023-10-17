package respErr

type ErrorResponse struct {
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
}

type Error struct {
	Error string `json:"respErr"`
}
