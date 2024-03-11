package route

import (
	"github.com/labstack/echo/v4"

	"github.com/bio426/gendor/internal/auth"
	"github.com/bio426/gendor/internal/product"
	"github.com/bio426/gendor/internal/user"
)

func RegisterRest(app *echo.Echo) {
	restRouter := app.Group("/rest")

	// Auth
	authRouter := restRouter.Group("/auth")
	authRouter.POST("/login", auth.Controller.Login)
	authRouter.POST("/logout", auth.Controller.Logout)

	userRouter := restRouter.Group("/user")
	userRouter.GET("", user.Controller.List)
	userRouter.POST("", user.Controller.Create)

	// Products
	productsRouter := restRouter.Group("/product")

	productsRouter.GET("", product.Controller.List)
}
