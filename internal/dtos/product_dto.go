package dto

type ProductInput struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
}