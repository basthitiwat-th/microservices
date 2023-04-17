package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface {
	CreateOrder(order Order) error
	CancelOrder(id uint) error
	FindOrderByOrderID(id uint) (*Order, error)
	FindOrderByUserID(userid uint) ([]Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// สร้าง Order
func (r *orderRepository) CreateOrder(order Order) error {
	return r.db.Create(&order).Error
}

// ยกเลิก Order
func (r *orderRepository) CancelOrder(orderID uint) error {
	var order Order

	tx := r.db.Begin()
	if err := r.db.Model(&order).Where("id = ?", orderID).Update("status", "cancel").Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// ค้นหาด้วย OrderId
func (r *orderRepository) FindOrderByOrderID(orderID uint) (*Order, error) {
	var order Order
	if err := r.db.Preload(clause.Associations).First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) FindOrderByUserID(userid uint) ([]Order, error) {
	var orders []Order
	if err := r.db.Preload(clause.Associations).Find(&orders, "user_id = ?", userid).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
