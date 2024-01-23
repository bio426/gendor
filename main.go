package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bio426/gendor/datasource"
	"github.com/bio426/gendor/route"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	datasource.InitConfig()
	datasource.InitPostgres()
	datasource.InitCloudinary()

	e := echo.New()
	e.HideBanner = true
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middlewares
	if datasource.Config.PROD {
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:  "./dist",
			HTML5: true,
		}))
	} else {
		e.Debug = true
		e.Use(middleware.CORS())
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		}))
	}

	// Routes
	route.RegisterApi(e)
	e.GET("", func(c echo.Context) error {
		return c.String(200, "main")
	})

	// Start server
	if datasource.Config.PROD {
		e.Logger.Fatal(e.Start(":8080"))
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}
