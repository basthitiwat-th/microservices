package services

import (
	"microservices/repositories"
)

type UserService interface {
	GetProfile(userid uint) (*UserProfile, error)
	GetOrderHistoryByUserID(userid uint) ([]OrderRespone, error)
}

type userServices struct {
	userRepository  repositories.UserRepository
	orderRepository repositories.OrderRepository
}

func NewUserServices(userRepository repositories.UserRepository, orderRepository repositories.OrderRepository) UserService {
	return &userServices{userRepository: userRepository, orderRepository: orderRepository}
}

// ดึง Profile
func (s *userServices) GetProfile(userid uint) (*UserProfile, error) {
	userDB, err := s.userRepository.FindUserByID(userid)
	if err != nil {
		return nil, err
	}
	user := UserProfile{
		Username:  userDB.Username,
		FirstName: userDB.FirstName,
		LastName:  userDB.LastName,
		Phone:     userDB.Phone,
		Email:     userDB.Email,
	}
	return &user, nil
}

// ดึงรายการซือของ User
func (s *userServices) GetOrderHistoryByUserID(id uint) ([]OrderRespone, error) {
	odersDB, err := s.orderRepository.FindOrderByUserID(id)
	if err != nil {
		return nil, err
	}
	var orders []OrderRespone
	for _, orderDB := range odersDB {
		orders = append(orders, OrderRespone{
			OrderID:      orderDB.ID,
			Quantity:     orderDB.Quantity,
			Status:       orderDB.Status,
			TotalPrice:   orderDB.TotalPrice,
			PricePerunit: orderDB.Product.UnitPrice,
			ProductName:  orderDB.Product.ProductName,
		})
	}
	return orders, nil
}
