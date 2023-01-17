package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID    string `gorm:"primaryKey"`
	Title string `json:"title" `
	Genre string `json:"genre"`
}
