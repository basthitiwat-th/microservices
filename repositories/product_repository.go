package repositories

import (
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product Product) error
	FindAllProducts() ([]Product, error)
	FindProductById(productCode uint) (*Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// สร้างใหม่
func (r productRepository) CreateProduct(product Product) error {
	return r.db.Create(&product).Error
}

// ค้นหาทั้งหมด
func (r productRepository) FindAllProducts() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

// ค้นหาด้วย ID
func (r productRepository) FindProductById(productID uint) (*Product, error) {
	var product Product
	if err := r.db.First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
