package model

type Job struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DepartmentID int64  `json:"department"`
}
