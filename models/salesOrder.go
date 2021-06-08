package models

import (
	"github.com/jinzhu/gorm"
)

type SalesOrder struct {
	gorm.Model
	SalesOrderDetailNumber string `json:"salesOrderNumber" form:"sales_order_number"`
	CustomerNumber         string `json:"supplierNumber" form:"supplier_number"`
}
