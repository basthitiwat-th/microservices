package repositories

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"total_price"`
	User       User    `json:"user" gorm:"foreignKey:UserID"`
	Product    Product `json:"product" gorm:"foreignKey:ProductID"`
}
