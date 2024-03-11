package workshop

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bio426/gendor/internal/core"
)

type AuthCtl core.Controller

func (ctl *AuthCtl) List(c echo.Context) error {
	err := Service.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "")
}

var Controller = &AuthCtl{}
