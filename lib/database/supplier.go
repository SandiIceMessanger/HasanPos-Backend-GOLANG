package database

import (
	"project/config"
	"project/models"
)

func GetSuppliers() (interface{}, error) {
	var suppliers []models.Supplier

	if err := config.DB.Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func GetSupplier(id int) (interface{}, error) {
	var supplier models.Supplier

	if err := config.DB.First(&supplier, id).Error; err != nil {
		return nil, err
	}
	return supplier, nil
}

func DeleteSupplier(id int) (interface{}, error) {
	var supplier models.Supplier
	if err := config.DB.First(&supplier, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&supplier).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdateSupplier(id int, supplier *models.Supplier) (interface{}, error) {
	var existingSupplier models.Supplier
	if err := config.DB.First(&existingSupplier, id).Error; err != nil {
		return nil, err
	}

	existingSupplier.SupplierName = supplier.SupplierName
	if updateErr := config.DB.Save(&existingSupplier).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingSupplier, nil
}

func CreateSupplier(supplier *models.Supplier) (interface{}, error) {

	if err := config.DB.Save(supplier).Error; err != nil {
		return nil, err
	}
	return supplier, nil
}
