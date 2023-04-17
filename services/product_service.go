package services

import (
	"microservices/repositories"
)

type ProductServices interface {
	CreateProduct(product ProductRequest) error
	GetAllProducts() ([]ProductResponse, error)
	GetProductById(productCode uint) (*ProductResponse, error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) ProductServices {
	return &productService{productRepository: productRepository}
}

// สร้างรายการใหม่
func (s productService) CreateProduct(productReq ProductRequest) error {
	product := repositories.Product{
		ProductName: productReq.ProductName,
		ProductCode: productReq.ProductCode,
		UnitPrice:   productReq.UnitPrice,
		Inventory:   productReq.Inventory,
	}
	err := s.productRepository.CreateProduct(product)
	return err
}

// ดึงข้อมูลทั้งหมด
func (s *productService) GetAllProducts() ([]ProductResponse, error) {
	var products []ProductResponse
	productsDB, err := s.productRepository.FindAllProducts()
	if err != nil {
		return nil, err
	}

	for _, product := range productsDB {
		products = append(products, ProductResponse{
			Id:          product.ID,
			ProductName: product.ProductName,
			ProductCode: product.ProductCode,
			UnitPrice:   product.UnitPrice,
			Inventory:   product.Inventory,
		})
	}

	return products, nil
}

// ดึง product  ตาม ID
func (s *productService) GetProductById(productID uint) (*ProductResponse, error) {
	productDB, err := s.productRepository.FindProductById(productID)
	if err != nil {
		return nil, err
	}
	product := ProductResponse{
		Id:          productDB.ID,
		ProductName: productDB.ProductName,
		ProductCode: productDB.ProductCode,
		UnitPrice:   productDB.UnitPrice,
		Inventory:   productDB.Inventory,
	}
	return &product, nil
}
