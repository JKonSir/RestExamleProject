package postgres

import (
	"storage"
	"model"
	"storage/postgres/support"
)

type ProductPostgresDBStorage struct {
	storage.ProductStorage
	db *support.PostgresDB
}

func NewProductStorage(db *support.PostgresDB) (storage.ProductStorage) {
	return &ProductPostgresDBStorage{db: db}
}

func (postgresDBStorage *ProductPostgresDBStorage) Read(productId int) (*model.Product, error) {
	product := model.Product{ID: productId}
	product.ID = productId
	err := postgresDBStorage.db.DB.QueryRow("SELECT name, price FROM products WHERE id=$1",
		productId).Scan(&product.Name, &product.Price)

	return &product, err
}

func (postgresDBStorage *ProductPostgresDBStorage) Update(product *model.Product) (*model.Product, error) {
	_, err := postgresDBStorage.db.DB.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
		product.Name, product.Price, product.ID)

	return product, err
}

func (postgresDBStorage *ProductPostgresDBStorage) Delete(productId int) error {
	_, err := postgresDBStorage.db.DB.Exec("DELETE FROM products WHERE id=$1", productId)

	return err
}

func (postgresDBStorage *ProductPostgresDBStorage) Create(product *model.Product) error {
	err := postgresDBStorage.db.DB.QueryRow(
		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
		product.Name, product.Price).Scan(&product.ID)

	return err
}

func (postgresDBStorage *ProductPostgresDBStorage) ReadAll(start, count int) ([]model.Product, error) {
	rows, err := postgresDBStorage.db.DB.Query(
		"SELECT id, name, price FROM products LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []model.Product{}

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
