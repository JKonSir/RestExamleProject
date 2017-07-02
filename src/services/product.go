package services

import (
	"model"
	"storage"
)

type ProductService interface {
	Get(productId int) (*model.Product, error)
}

type ProductServiceImpl struct {
	ProductService
	productStorage storage.ProductStorage
}

func NewProductService(productStorage storage.ProductStorage) (ProductService) {
	return &ProductServiceImpl{productStorage: productStorage}
}

func (productService *ProductServiceImpl) Get(productId int) (*model.Product, error) {
	return productService.productStorage.Get(productId)
}
