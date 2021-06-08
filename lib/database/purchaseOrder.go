package database

import (
	"project/config"
	"project/models"
)

func GetPurchaseOrders() (interface{}, error) {
	var purchaseOrders []models.PurchaseOrder

	if err := config.DB.Find(&purchaseOrders).Error; err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func GetPurchaseOrder(id int) (interface{}, error) {
	var purchaseOrder models.PurchaseOrder

	if err := config.DB.First(&purchaseOrder, id).Error; err != nil {
		return nil, err
	}
	return purchaseOrder, nil
}

func DeletePurchaseOrder(id int) (interface{}, error) {
	var purchaseOrder models.PurchaseOrder
	if err := config.DB.First(&purchaseOrder, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&purchaseOrder).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdatePurchaseOrder(id int, purchaseOrder *models.PurchaseOrder) (interface{}, error) {
	var existingPurchaseOrder models.PurchaseOrder
	if err := config.DB.First(&existingPurchaseOrder, id).Error; err != nil {
		return nil, err
	}

	existingPurchaseOrder.PurchaseOrderDetailNumber = purchaseOrder.PurchaseOrderDetailNumber
	existingPurchaseOrder.SupplierCode = purchaseOrder.SupplierCode
	if updateErr := config.DB.Save(&existingPurchaseOrder).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingPurchaseOrder, nil
}

func CreatePurchaseOrder(purchaseOrder *models.PurchaseOrder) (interface{}, error) {

	if err := config.DB.Save(purchaseOrder).Error; err != nil {
		return nil, err
	}
	return purchaseOrder, nil
}
