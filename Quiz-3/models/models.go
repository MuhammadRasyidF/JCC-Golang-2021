package models

import (
	"time"
)

type (
	Book struct {
		ID                 int       `json:"id"`
		Title              string    `json:"title"`
		Description        string    `json:"description"`
		Image_url          string    `json:"image_url"`
		Release_year       int       `json:"release_year"`
		Price              string    `json:"price"`
		Total_page         string    `json:"total_page"`
		Kategori_ketebalan string    `json:"kategori_ketebalan"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
	}
)
