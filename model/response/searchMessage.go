package response

import "project1/model/entity"

type SearchResponse struct {
	Status  int             `json:"status"`
	Total   int64           `json:"total"`
	Page    int             `json:"page"`
	PerPage int             `json:"per_Page"`
	Data    []entity.Santri `json:"data"`
}
