package models

import (
	"github.com/jinzhu/gorm"
)

type ProductCategory struct {
	gorm.Model
	CategoryCode 	string `json:"categoryCode" form:"category_code"`
	CategoryName	string `json:"categoryName" form:"category_name"`
}