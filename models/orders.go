package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Order struct {
	Order_ID      uint      `gorm:"primary_key;auto_increment;not_null" json:"order_id"`
	Ordered_At    time.Time `gorm:"not_null"`
	Customer_Name string    `gorm:"not_null" json:"customerName"`
	Items         []Item    `json:"items"`
}
