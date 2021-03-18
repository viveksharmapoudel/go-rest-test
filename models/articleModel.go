package models

import(
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type Article struct{
	gorm.Model
	Title string   `json:"title"`
	Description string  `json:"description"`
}