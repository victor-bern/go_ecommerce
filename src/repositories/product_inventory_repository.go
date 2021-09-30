package repositories

import (
	"database/sql"
	"pg-conn/src/models"
)

type ProductInventoryRepository interface {
	GetById(id string) (models.ProductInventory, error)
	GetByProductId(productid string) models.ProductInventory
	All() ([]models.ProductInventory, error)
	InserProductInInventory(product models.ProductInventory) (models.ProductInventory, error)
	UpdateProductStock(productid string, rq int) error
	DeleteProductInventory(id string) error
}

type ProductInventoryRepo struct {
	db *sql.DB
}

func NewProductInventoryRepo(db *sql.DB) *ProductInventoryRepo {
	return &ProductInventoryRepo{
		db: db,
	}
}

func (pir *ProductInventoryRepo) GetById(id string) (models.ProductInventory, error) {
	var productInventory models.ProductInventory
	stmt, err := pir.db.Prepare("SELECT * FROM product_inventory WHERE id = $1")

	if err != nil {
		return models.ProductInventory{}, err
	}

	result, err := stmt.Query(id)
	if err != nil {
		return models.ProductInventory{}, err
	}

	err = result.Scan(&productInventory.ID, &productInventory.ProductId, &productInventory.Stock, &productInventory.CreatedAt, &productInventory.UpdatedAt)
	if err != nil {
		return models.ProductInventory{}, err
	}

	return productInventory, nil
}

func (pir *ProductInventoryRepo) GetByProductId(id string) (models.ProductInventory, error) {
	var productInventory models.ProductInventory
	stmt, err := pir.db.Prepare("SELECT * FROM product_inventory WHERE product_id = $1")

	if err != nil {
		return models.ProductInventory{}, err
	}

	result, err := stmt.Query(id)
	if err != nil {
		return models.ProductInventory{}, err
	}

	err = result.Scan(&productInventory.ID, &productInventory.ProductId, &productInventory.Stock, &productInventory.CreatedAt, &productInventory.UpdatedAt)
	if err != nil {
		return models.ProductInventory{}, err
	}

	return productInventory, nil
}

func (pir *ProductInventoryRepo) All() ([]models.ProductInventory, error) {
	var productsInInventory []models.ProductInventory
	stmt, err := pir.db.Prepare("SELECT * FROM product_inventory")
	if err != nil {
		return []models.ProductInventory{}, err
	}

	result, err := stmt.Query()
	if err != nil {
		return []models.ProductInventory{}, err
	}

	for result.Next() {
		var productInventory models.ProductInventory
		result.Scan(&productInventory.ID, &productInventory.ProductId, &productInventory.Stock, &productInventory.CreatedAt, &productInventory.UpdatedAt)
		productsInInventory = append(productsInInventory, productInventory)
	}

	return productsInInventory, nil
}

func (pir *ProductInventoryRepo) InserProductInInventory(pi models.ProductInventory, id string) (models.ProductInventory, error) {
	stmt, err := pir.db.Prepare("INSERT INTO product_inventory (product_id, stock) VALUES ($1, $2)")

	if err != nil {
		return models.ProductInventory{}, err
	}

	_, err = stmt.Query(pi.ProductId, pi.Stock)
	if err != nil {
		return models.ProductInventory{}, err
	}

	return pi, nil
}

func (pir *ProductInventoryRepo) UpdateProductStock(productid string, rq int) error {
	productInventory, err := pir.GetByProductId(productid)
	if err != nil {
		return err
	}
	stmt, err := pir.db.Prepare("UPDATE product_inventory SET stock = $1 WHERE product_id = $2")
	if err != nil {
		return err
	}

	_, err = stmt.Query(productInventory.Stock-rq, productid)
	if err != nil {
		return err
	}

	return nil
}

func (pir *ProductInventoryRepo) DeleteProductInventory(id string) error {
	stmt, err := pir.db.Prepare("DELETE FROM product_inventory WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Query(id)
	if err != nil {
		return err
	}
	return nil
}
