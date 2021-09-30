package models

import "time"

type ProductInventory struct {
	ID        uint      `json:"id"`
	ProductId string    `json:"product_id"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
