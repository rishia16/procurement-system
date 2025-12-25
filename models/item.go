package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name  string
	Stock int
	Price float64
}
