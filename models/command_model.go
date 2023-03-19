package models

type Command struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Platform    string `json:"platform"`
	Description string `json:"description"`
}
