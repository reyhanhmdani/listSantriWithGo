package entity

type Status struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func (Status) TableName() string {
	return "Status"
}
