package models

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	CustomerName        string `json:"customerName" form:"customer_name"`
	CustomerPhoneNumber string `json:"customerPhoneNumber" form:"customer_phone_number"`
	CustomerPoint       int    `json:"CustomerPoint" form:"customer_point"`
}
