package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetProductCategoriesController(c echo.Context) error {
	productsCategories, e := database.GetProductCategories()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":            "success",
		"productCategories": productsCategories,
	})
}

func GetProductCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a product category, product category with ID " + c.Param("id") + " is not found",
		})
	}

	productCategory, getErr := database.GetProductCategory(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":          "success",
		"productCategory": productCategory,
	})
}

func DeleteProductCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a product category, user with ID " + c.Param("id") + " is not found",
		})
	}

	if _, deleteErr := database.DeleteProductCategory(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"productCategory": "success",
	})
}

func UpdateProductCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a product category, product category with ID " + c.Param("id") + " is not found",
		})
	}

	var UpdateProductCategory models.ProductCategory
	c.Bind(&UpdateProductCategory)
	productCategory, updateErr := database.UpdateProductCategory(id, &UpdateProductCategory)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success create new productCategory",
		"productCategory": productCategory,
	})
}

func CreateProductCategoryController(c echo.Context) error {
	productCategory := models.ProductCategory{}
	c.Bind(&productCategory)

	productCategories, e := database.CreateProductCategory(&productCategory)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success create new productCategory",
		"productCategory": productCategories,
	})
}
