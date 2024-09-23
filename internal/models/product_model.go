package models

type Product struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Description string `json:"description"`
}