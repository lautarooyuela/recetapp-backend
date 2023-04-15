package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Email    string `gorm:"not null;uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Token    string `json:"token"`
}
