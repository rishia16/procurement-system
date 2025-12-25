package models

import "gorm.io/gorm"

type PurchasingDetail struct {
	gorm.Model
	PurchasingID uint
	ItemID       uint
	Qty          int
	SubTotal     float64
}
