package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetSuppliersController(c echo.Context) error {
	suppliers, e := database.GetSuppliers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":            "success",
		"productCategories": suppliers,
	})
}

func GetSupplierController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a supplier, product category with ID " + c.Param("id") + " is not found",
		})
	}

	supplier, getErr := database.GetSupplier(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"supplier": supplier,
	})
}

func DeleteSupplierController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a supplier, user with ID " + c.Param("id") + " is not found",
		})
	}

	if _, deleteErr := database.DeleteSupplier(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"supplier": "success",
	})
}

func UpdateSupplierController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a product category, supplier with ID " + c.Param("id") + " is not found",
		})
	}

	var UpdateSupplier models.Supplier
	c.Bind(&UpdateSupplier)
	supplier, updateErr := database.UpdateSupplier(id, &UpdateSupplier)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new supplier",
		"supplier": supplier,
	})
}

func CreateSupplierController(c echo.Context) error {
	supplier := models.Supplier{}
	c.Bind(&supplier)

	suppliers, e := database.CreateSupplier(&supplier)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new supplier",
		"supplier": suppliers,
	})
}
