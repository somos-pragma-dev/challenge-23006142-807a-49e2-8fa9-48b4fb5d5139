package services

import (
	"errors"
	"gorm.io/gorm"
	"src/models"
)

type CreateProductDTO struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	Category string  `json:"category"`
}

type UpdateProductDTO struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	Category string  `json:"category"`
}

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

func (ps *ProductService) CreateProduct(productDTO CreateProductDTO) (models.Product, error) {
	product := models.Product{
		Name:    productDTO.Name,
		Price:   productDTO.Price,
		Stock:   productDTO.Stock,
		Category: productDTO.Category,
	}
	if err := ps.db.Create(&product).Error; err!= nil {
		return models.Product{}, err
	}
	return product, nil
}

func (ps *ProductService) GetProduct(id string) (models.Product, error) {
	var product models.Product
	if err := ps.db.First(&product, id).Error; err!= nil {
		return models.Product{}, err
	}
	return product, nil
}

func (ps *ProductService) UpdateProduct(id string, productDTO UpdateProductDTO) (models.Product, error) {
	var product models.Product
	if err := ps.db.First(&product, id).Error; err!= nil {
		return models.Product{}, err
	}
	product.Name = productDTO.Name
	product.Price = productDTO.Price
	product.Stock = productDTO.Stock
	product.Category = productDTO.Category
	if err := ps.db.Save(&product).Error; err!= nil {
		return models.Product{}, err
	}
	return product, nil
}

func (ps *ProductService) DeleteProduct(id string) error {
	var product models.Product
	if err := ps.db.First(&product, id).Error; err!= nil {
		return err
	}
	if err := ps.db.Delete(&product).Error; err!= nil {
		return err
	}
	return nil
}