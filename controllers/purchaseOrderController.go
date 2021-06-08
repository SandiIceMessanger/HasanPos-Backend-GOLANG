package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetPurchaseOrdersController(c echo.Context) error {
	purchaseOrders, e := database.GetPurchaseOrders()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":            "success",
		"productCategories": purchaseOrders,
	})
}

func GetPurchaseOrderController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a purchaseOrder, product category with ID " + c.Param("id") + " is not found",
		})
	}

	purchaseOrder, getErr := database.GetPurchaseOrder(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":        "success",
		"purchaseOrder": purchaseOrder,
	})
}

func DeletePurchaseOrderController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a purchaseOrder, user with ID " + c.Param("id") + " is not found",
		})
	}

	if _, deleteErr := database.DeletePurchaseOrder(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"purchaseOrder": "success",
	})
}

func UpdatePurchaseOrderController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a product category, purchaseOrder with ID " + c.Param("id") + " is not found",
		})
	}

	var UpdatePurchaseOrder models.PurchaseOrder
	c.Bind(&UpdatePurchaseOrder)
	purchaseOrder, updateErr := database.UpdatePurchaseOrder(id, &UpdatePurchaseOrder)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success create new purchaseOrder",
		"purchaseOrder": purchaseOrder,
	})
}

func CreatePurchaseOrderController(c echo.Context) error {
	purchaseOrder := models.PurchaseOrder{}
	c.Bind(&purchaseOrder)

	purchaseOrders, e := database.CreatePurchaseOrder(&purchaseOrder)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success create new purchaseOrder",
		"purchaseOrder": purchaseOrders,
	})
}
