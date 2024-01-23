package route

import (
	"github.com/bio426/gendor/controller"
	"github.com/labstack/echo/v4"
)

func RegisterApi(app *echo.Echo) {
	apiRouter := app.Group("/api")

	// Auth
	authRouter := apiRouter.Group("/auth")
	authRouter.POST("", controller.AuthLogin)

	// Products
	productsRouter := apiRouter.Group("/product")
	productsRouter.GET("", controller.Products.List)
	productsRouter.POST("", controller.Products.Create)
	productsRouter.GET("/category", controller.Products.CategoryList)
	productsRouter.POST("/category", controller.Products.CategoryCreate)
}
