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

func DeleteProduct(id int) (interface{}, error) {
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&product).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdateProduct(id int, product *models.Product) (interface{}, error) {
	var existingProduct models.Product
	if err := config.DB.First(&existingProduct, id).Error; err != nil {
		return nil, err
	}

	existingProduct.ProductCode = product.ProductCode
	existingProduct.CategoryCode = product.CategoryCode
	existingProduct.ProductsName = product.ProductsName
	existingProduct.BuyingPrice = product.BuyingPrice
	existingProduct.SellingPrice = product.SellingPrice
	existingProduct.PriceOne = product.PriceOne
	existingProduct.PriceTwo = product.PriceTwo
	existingProduct.PriceThree = product.PriceThree

	if updateErr := config.DB.Save(&existingProduct).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingProduct, nil
}

func CreateProduct(product *models.Product) (interface{}, error) {

	if err := config.DB.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
