package controllers

import (
	"net/http"
	"strconv"
	
	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetProductsController(c echo.Context) error {
	products, e := database.GetProducts()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": products,
	})
}

func GetProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{} {
			"message": "failed to get a product, product with ID " + c.Param("id") + " is not found",
		})
	}

	product, getErr := database.GetProduct(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":          "success",
		"product": product,
	})
}

func CreateProductController(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	products, e := database.CreateProduct(&product)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new product",
		"product": products,
	})
}
