package storage

import "model"

type ProductStorage interface {
	ReadAll(start, count int) ([]model.Product, error)
	Read(productId int) (*model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) (*model.Product, error)
	Delete(productId int) error

	Destroy()
}
