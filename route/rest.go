package route

import (
	"github.com/labstack/echo/v4"

	"github.com/bio426/gendor/internal/auth"
	"github.com/bio426/gendor/internal/product"
	"github.com/bio426/gendor/internal/user"
	"github.com/bio426/gendor/internal/workshop"
)

func RegisterRest(app *echo.Echo) {
	restRouter := app.Group("/rest")

	authRouter := restRouter.Group("/auth")
	authRouter.POST("/login", auth.Controller.Login)
	authRouter.POST("/logout", auth.Controller.Logout)

	userRouter := restRouter.Group("/user")
	userRouter.Use(auth.Middleware, auth.MiddlewareWithRoles([]string{"super"}))
	userRouter.GET("", user.Controller.List)
	userRouter.POST("", user.Controller.Create)

	workshopRouter := restRouter.Group("/workshop")
	workshopRouter.Use(auth.Middleware)
	workshopRouter.GET("", workshop.Controller.List)
	workshopRouter.POST("", workshop.Controller.Create)

	productsRouter := restRouter.Group("/product")

	productsRouter.GET("", product.Controller.List)
}
