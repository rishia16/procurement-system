package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name    string
	Email   string
	Address string
}
