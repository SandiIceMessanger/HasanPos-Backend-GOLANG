package models

import (
	"github.com/jinzhu/gorm"
)

type Supplier struct {
	gorm.Model
	SupplierCode        string `json:"supplierCode" form:"supplier_code"`
	SupplierName        string `json:"supplierName" form:"supplier_name"`
	SupplierPhoneNumber string `json:"supplierPhoneNumber" form:"customer_phone_number"`
}
