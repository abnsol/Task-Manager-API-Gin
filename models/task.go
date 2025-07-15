package models

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Time        string `json:"time"`
	Status      bool   `json:"status"`
}
