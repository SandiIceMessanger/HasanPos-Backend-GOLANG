package database

import (
	"project/config"
	"project/models"
)

func GetProductCategories() (interface{}, error) {
	var productCategories []models.ProductCategory

	if err := config.DB.Find(&productCategories).Error; err != nil {
		return nil, err
	}
	return productCategories, nil
}

func GetProductCategory(id int) (interface{}, error) {
	var productCategory []models.ProductCategory

	if err := config.DB.First(&productCategory, id).Error; err != nil {
		return nil, err
	}
	return productCategory, nil
}

func DeleteProductCategory(id int) (interface{}, error) {
	var productCategory models.ProductCategory
	if err := config.DB.First(&productCategory, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&productCategory).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdateProductCategory(id int, productCategory *models.ProductCategory) (interface{}, error) {
	var existingProductCategory models.ProductCategory
	if err := config.DB.First(&existingProductCategory, id).Error; err != nil {
		return nil, err
	}

	existingProductCategory.CategoryCode = productCategory.CategoryCode
	existingProductCategory.CategoryName = productCategory.CategoryName
	if updateErr := config.DB.Save(&existingProductCategory).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingProductCategory, nil
}

func CreateProductCategory(productCategory *models.ProductCategory) (interface{}, error) {
	if err := config.DB.Save(productCategory).Error; err != nil {
		return nil, err
	}
	return productCategory, nil
}
