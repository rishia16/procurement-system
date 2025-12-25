package models

import "gorm.io/gorm"

type Purchasing struct {
	gorm.Model
	Date       string
	SupplierID uint
	UserID     uint
	GrandTotal float64

	Details []PurchasingDetail
}
