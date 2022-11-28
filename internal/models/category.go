package models

type Category struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Parent *Category `json:"parent"`
}  