package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name         string  `json:"name"`
	Barcode      string  `json:"barcode"`
	Image        string  `json:"image"`
	Stock        int     `json:"stock"`
	SuppliedFrom string  `json:"suppliedfrom"`
	BuyPrice     float32 `json:"buyprice"`
	SellPrice    float32 `json:"sellprice"`
}
