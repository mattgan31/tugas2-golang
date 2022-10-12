package models

import (
	_ "github.com/jinzhu/gorm"
	_ "gorm.io/gorm"
)

type Item struct {
	// gorm.Model
	Item_ID     uint   `gorm:"primary_key;auto_increment;not_null" json:"item_id"`
	Item_Code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_ID    uint   `json:"order_id"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
}
