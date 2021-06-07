package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUsersController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)

	r.GET("/products", controllers.GetProductsController)
	r.GET("/products/:id", controllers.GetProductController)
	r.POST("/products", controllers.CreateProductController)

	r.GET("/product_categories", controllers.GetProductCategoriesController)
	r.GET("/product_categories/:id", controllers.GetProductCategoryController)
	r.DELETE("/product_categories/:id", controllers.DeleteProductCategoryController)
	r.PUT("/product_categories/:id", controllers.UpdateProductCategoryController)
	r.POST("/product_categories", controllers.CreateProductCategoryController)

	r.GET("/customers", controllers.GetCustomersController)
	r.GET("/customers/:id", controllers.GetCustomerController)
	r.DELETE("/customers/:id", controllers.DeleteCustomerController)
	r.PUT("/customers/:id", controllers.UpdateCustomerController)
	r.POST("/customers", controllers.CreateCustomerController)

	return e
}
