package storage

import "model"

type ProductStorage interface {
	//GetAll() (*[]model.Product)
	Get(productId int) (*model.Product, error)
	//Add(product *model.Product) (*model.Product)
	//Update(product *model.Product) (*model.Product)
	//Delete(productId string)

	Destroy()
}
