package services

import (
	"microservices/repositories"
)

type OrderService interface {
	CreateOrder(order OrderRequest) error
	CancelOrder(orderID uint) error
	GetOrderByOrderID(orderID uint) (*OrderRespone, error)
}

type orderService struct {
	orderRepository   repositories.OrderRepository
	productRepository repositories.ProductRepository
}

func NewOrderService(orderRepository repositories.OrderRepository, productRepository repositories.ProductRepository) OrderService {
	return &orderService{orderRepository: orderRepository, productRepository: productRepository}
}

func (s *orderService) CreateOrder(orderReq OrderRequest) error {

	product, err := s.productRepository.FindProductById(orderReq.ProductID)
	if err != nil {
		return err
	}

	order := repositories.Order{
		UserID:     orderReq.UserID,
		ProductID:  orderReq.ProductID,
		Quantity:   orderReq.Quantity,
		Status:     "pending",
		TotalPrice: float64(orderReq.Quantity) * product.UnitPrice,
	}

	err = s.orderRepository.CreateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

// ยกเลิก order
func (s *orderService) CancelOrder(orderID uint) error {
	//ค้นหาก่อนว่ามี orderid นี้จริงๆหรือไม่
	_, err := s.orderRepository.FindOrderByOrderID(orderID)
	if err != nil {
		return err
	}

	err = s.orderRepository.CancelOrder(orderID)
	if err != nil {
		return err
	}
	return nil
}

// ดึง order ตาม ID
func (s *orderService) GetOrderByOrderID(orderID uint) (*OrderRespone, error) {

	orderDB, err := s.orderRepository.FindOrderByOrderID(orderID)
	if err != nil {
		return nil, err
	}
	order := OrderRespone{
		OrderID:      orderDB.ID,
		Quantity:     orderDB.Quantity,
		Status:       orderDB.Status,
		TotalPrice:   orderDB.TotalPrice,
		PricePerunit: orderDB.Product.UnitPrice,
		ProductName:  orderDB.Product.ProductName,
	}

	return &order, nil
}
