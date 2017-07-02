package services

import (
	"model"
	"storage"
)

type ProductService interface {
	FindOne(productId int) (*model.Product, error)
	Find(start, count int) ([]model.Product, error)
	Add(product *model.Product) error
	Update(product *model.Product) (*model.Product, error)
	Delete(productId int) error

	Destroy()
}

type ProductServiceImpl struct {
	ProductService
	productStorage storage.ProductStorage
}

func NewProductService(productStorage storage.ProductStorage) (ProductService) {
	return &ProductServiceImpl{productStorage: productStorage}
}

func (productService *ProductServiceImpl) FindOne(productId int) (*model.Product, error) {
	return productService.productStorage.Read(productId)
}

func (productService *ProductServiceImpl) Find(start, count int) ([]model.Product, error) {
	return productService.productStorage.ReadAll(start, count)
}

func (productService *ProductServiceImpl) Add(product *model.Product) error {
	return productService.productStorage.Create(product)
}

func (productService *ProductServiceImpl) Update(product *model.Product) (*model.Product, error) {
	return productService.productStorage.Update(product)
}

func (productService *ProductServiceImpl) Delete(productId int) error {
	return productService.productStorage.Delete(productId)
}