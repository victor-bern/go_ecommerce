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
	Update(product models.ProductInventory, id string) (models.ProductInventory, error)
	Delete(id string) error
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
