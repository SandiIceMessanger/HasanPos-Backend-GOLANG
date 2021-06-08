package models

import (
	"github.com/jinzhu/gorm"
)

type PurchaseOrder struct {
	gorm.Model
	PurchaseOrderDetailNumber string `json:"purchaseOrderNumber" form:"purchase_order_number"`
	SupplierCode              string `json:"supplierCode" form:"supplier_code"`
}
