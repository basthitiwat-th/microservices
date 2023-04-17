package services

type ProductRequest struct {
	ProductName string  `json:"product_name" gorm:"idx_productname;unique;not null"`
	ProductCode string  `json:"product_code" gorm:"idx_productcode;unique;not null;"`
	UnitPrice   float64 `json:"unit_price" gorm:"not null"`
	Inventory   int     `json:"inventory"`
}

type ProductResponse struct {
	Id          uint    `json:"id"`
	ProductName string  `json:"product_name"`
	ProductCode string  `json:"product_code"`
	UnitPrice   float64 `json:"unit_price"`
	Inventory   int     `json:"inventory"`
}
