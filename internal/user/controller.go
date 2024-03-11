package user

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/bio426/gendor/internal/core"
)

type AuthCtl core.Controller

type CtlListRow struct {
	Id        int32     `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}
type CtlListResponse struct {
	Rows []CtlListRow `json:"rows"`
}

func (ctl *AuthCtl) List(c echo.Context) error {
	res, err := Service.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *AuthCtl) Create(c echo.Context) error {
	body := struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := Service.Create(c.Request().Context(), SvcCreateParams{
		Username: body.Username,
		Password: body.Password,
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

var Controller = &AuthCtl{}
