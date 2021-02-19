package model

type Department struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ManagerID string `json:"user"`
}
