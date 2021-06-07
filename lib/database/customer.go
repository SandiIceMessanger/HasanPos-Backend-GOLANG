package database

import (
	"project/config"
	"project/models"
)

func GetCustomers() (interface{}, error) {
	var customers []models.Customer

	if err := config.DB.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomer(id int) (interface{}, error) {
	var customer models.Customer

	if err := config.DB.First(&customer, id).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func DeleteCustomer(id int) (interface{}, error) {
	var customer models.Customer
	if err := config.DB.First(&customer, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&customer).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdateCustomer(id int, customer *models.Customer) (interface{}, error) {
	var existingCustomer models.Customer
	if err := config.DB.First(&existingCustomer, id).Error; err != nil {
		return nil, err
	}

	existingCustomer.CustomerName = customer.CustomerName
	if updateErr := config.DB.Save(&existingCustomer).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingCustomer, nil
}

func CreateCustomer(customer *models.Customer) (interface{}, error) {

	if err := config.DB.Save(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}
