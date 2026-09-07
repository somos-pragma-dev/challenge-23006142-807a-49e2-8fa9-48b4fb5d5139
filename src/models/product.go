package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string  `gorm:"unique;not null"`
	Price   float64 `gorm:"not null"`
	Stock   int     `gorm:"not null"`
	Category string  `gorm:"not null"`
}