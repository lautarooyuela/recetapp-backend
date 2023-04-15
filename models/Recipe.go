package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title        string `gorm:"not null;unique_index" json:"title"`
	Ingredients  string `json:"ingredients"`
	Steps        string `json:"steps"`
	Type         string `json:"type"`
	Group        string `json:"group"`
	Done         bool   `gorm:"default:false" json:"done"`
	Calification int    `json:"calification"`
	Image        string `json:"image"`
	UserID       uint   `json:"user_id"`
	Email        string `json:"email"`
}
