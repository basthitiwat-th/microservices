package repositories

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string  `json:"product_name" gorm:"idx_productname;unique;not null"`
	ProductCode string  `json:"product_code" gorm:"idx_productcode;unique;not null;"`
	UnitPrice   float64 `json:"unit_price" gorm:"not null"`
	Inventory   int     `json:"inventory"`
}
