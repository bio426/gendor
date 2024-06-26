package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bio426/gendor/datasource"
	"github.com/bio426/gendor/route"
)

func main() {
	datasource.InitConfig()
	err := datasource.InitPostgres()
	if err != nil {
		panic(err)
	}
	// datasource.InitCloudinary()

	e := echo.New()
	e.HideBanner = true
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middlewares
	if datasource.Config.PROD {
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:  "./dist",
			HTML5: true,
		}))
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Skipper: func(c echo.Context) bool {
				path := c.Path()
				if strings.Contains(path, "/rest") {
					return false
				}
				return true
			},
		}))
	} else {
		e.Debug = true
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowCredentials: true,
		}))
		// This logger breaks the spa middleware in prod
		e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogMethod:   true,
			LogURI:      true,
			LogStatus:   true,
			LogError:    true,
			HandleError: true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				if v.Error == nil {
					fmt.Printf("%s %s %d\n", v.Method, v.URI, v.Status)
				} else {
					fmt.Printf("%s %s %d %s\n", v.Method, v.URI, v.Status, v.Error.Error())
				}
				return nil
			},
			Skipper: func(c echo.Context) bool {
				path := c.Path()
				if strings.Contains(path, "/rest") {
					return false
				}
				return true
			},
		}))
	}

	// Routes
	route.RegisterRest(e)

	// Start server
	if datasource.Config.PROD {
		e.Logger.Fatal(e.Start(":8080"))
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
