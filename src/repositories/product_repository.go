package repositories

import (
	"database/sql"
	"fmt"
	"pg-conn/src/models"
)

type ProductRepository interface {
	GetByTitle(name string) (models.Product, error)
	GetById(id string) (models.Product, error)
	All() ([]models.Product, error)
	InserProduct(product models.Product) (models.Product, error)
	Update(product models.Product, id string) (models.Product, error)
	Delete(id string) error
}

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (pr *ProductRepo) GetByTitle(name string) (models.Product, error) {
	stmt, err := pr.db.Prepare("SELECT * FROM product WHERE title = $1")
	var product models.Product
	if err != nil {
		return models.Product{}, err
	}

	result, err := stmt.Query(name)
	if err != nil {
		return models.Product{}, err
	}

	err = result.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepo) GetById(id string) (models.Product, error) {
	stmt, err := pr.db.Prepare("SELECT * FROM product WHERE id = $1")
	var product models.Product
	if err != nil {
		return models.Product{}, err
	}

	result, err := stmt.Query(id)
	if err != nil {
		return models.Product{}, err
	}

	err = result.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepo) All() ([]models.Product, error) {
	stmt, err := pr.db.Prepare("SELECT * FROM product")
	var products []models.Product
	if err != nil {
		return []models.Product{}, err
	}

	result, err := stmt.Query()
	if err != nil {
		return []models.Product{}, err
	}
	for result.Next() {
		var product models.Product
		result.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepo) InserProduct(product models.Product, quantity int) (models.Product, error) {
	stmt, err := pr.db.Prepare("INSERT INTO product (title, description, price) VALUES ($1, $2, $3)")
	if err != nil {
		return models.Product{}, err
	}

	result, err := stmt.Exec(product.Title, product.Description, product.Price)
	if err != nil {
		return models.Product{}, err
	}

	pir := NewProductInventoryRepo(pr.db)
	lid, err := result.LastInsertId()
	if err != nil {
		return models.Product{}, err
	}
	_, err = pir.InserProductInInventory(models.ProductInventory{
		Stock: quantity,
	}, fmt.Sprint(lid))

	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
