package models

import (
	"github.com/jinzhu/gorm"
)

type PurchaseOrderDetail struct {
	gorm.Model
	ProductCode string `json:"productCode" form:"product_code"`
	Quantity    string `json:"quantity" form:"quantity"`
}
