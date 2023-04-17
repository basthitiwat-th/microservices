package services

type OrderRespone struct {
	OrderID      uint    `json:"order_id"`
	Quantity     int     `json:"quantity"`
	Status       string  `json:"status"`
	TotalPrice   float64 `json:"total_price"`
	ProductName  string  `json:"product_name"`
	PricePerunit float64 `json:"price_peruint"`
}

type OrderRequest struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
