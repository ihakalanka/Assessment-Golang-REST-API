package models

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name     string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}