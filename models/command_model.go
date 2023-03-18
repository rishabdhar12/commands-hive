package models

type Command struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Paltform    string `json:"paltform"`
	Description string `json:"description"`
}
