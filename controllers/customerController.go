package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetCustomersController(c echo.Context) error {
	customers, e := database.GetCustomers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":            "success",
		"productCategories": customers,
	})
}

func GetCustomerController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a customer, product category with ID " + c.Param("id") + " is not found",
		})
	}

	customer, getErr := database.GetCustomer(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"customer": customer,
	})
}

func DeleteCustomerController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a customer, user with ID " + c.Param("id") + " is not found",
		})
	}

	if _, deleteErr := database.DeleteCustomer(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"customer": "success",
	})
}

func UpdateCustomerController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a product category, customer with ID " + c.Param("id") + " is not found",
		})
	}

	var UpdateCustomer models.Customer
	c.Bind(&UpdateCustomer)
	customer, updateErr := database.UpdateCustomer(id, &UpdateCustomer)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new customer",
		"customer": customer,
	})
}

func CreateCustomerController(c echo.Context) error {
	customer := models.Customer{}
	c.Bind(&customer)

	customers, e := database.CreateCustomer(&customer)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new customer",
		"customer": customers,
	})
}
