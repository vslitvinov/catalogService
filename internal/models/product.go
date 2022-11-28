package models

import "time"

type Product struct {
	ID              string  `json:"id"`
	ShopID          string  `json:"shop_id"`
	Category        string  `json:"category"`
	ProductTy       string  `json:"product_type"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Status          string  `json:"status"`
	Price           float64 `json:"price"`
	Count           int64   `json:"count"`
	SendingDeadline int64   `json:"sending_deadline"`
	Materials       string  `json:"materials"` ///?
	Colors          string  `json:"colors"`
	Tags            []Tag   `json:"tags"`
	Size            string  `json:"size"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

