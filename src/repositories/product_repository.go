package repositories

import (
	"database/sql"
	"pg-conn/src/models"
)

type ProductRepository interface {
	GetByTitle(name string) (models.Product, error)
	GetById(id string) (models.Product, error)
	All() ([]models.Product, error)
	InserProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product, id string) error
	DeleteProduct(id string) error
}

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (pr *ProductRepo) GetByTitle(name string, product *models.Product) error {
	stmt, err := pr.db.Prepare("SELECT * FROM product WHERE title = $1")
	if err != nil {
		return err
	}

	result, err := stmt.Query(name)
	if err != nil {
		return err
	}

	err = result.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
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
	stmt, err := pr.db.Prepare("SELECT q.id, q.title, q.description, q.price, pi.stock, q.created_at, q.updated_at FROM product AS q JOIN product_inventory AS pi ON q.id = pi.product_id")
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
		result.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepo) InserProduct(product models.Product, quantity int) (models.Product, error) {
	id := 0
	err := pr.db.QueryRow("INSERT INTO product (title, description, price) VALUES ($1, $2, $3) RETURNING id", product.Title, product.Description, product.Price).Scan(&id)
	if err != nil {
		return models.Product{}, err
	}

	pir := NewProductInventoryRepo(pr.db)

	_, err = pir.InserProductInInventory(models.ProductInventory{
		Stock: quantity,
	}, id)

	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (pr *ProductRepo) UpdateProduct(product models.Product, id string) error {
	_, err := pr.GetById(id)
	if err != nil {
		return err
	}
	stmt, err := pr.db.Prepare("UPDATE product SET title = $1, description = $2, price = $3 WHERE id = $4")
	if err != nil {
		return err
	}

	_, err = stmt.Query(product.Title, product.Description, product.Price, id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductRepo) DeleteProduct(id string) error {
	stmt, err := pr.db.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Query(id)
	if err != nil {
		return err
	}

	return nil
}
