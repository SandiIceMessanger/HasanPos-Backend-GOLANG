package database

import (
	"project/config"
	"project/models"
)

func GetProducts() (interface{}, error) {
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProduct(id int) (interface{}, error) {
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func CreateProduct(product *models.Product) (interface{}, error) {

	if err := config.DB.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
