package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string   `gorm:"not null" json:"name"`
	Email   string   `gorm:"not null" json:"email"`
	Recipes []Recipe `json:"recipes"`
}
