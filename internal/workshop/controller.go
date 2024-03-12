package workshop

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bio426/gendor/internal/auth"
	"github.com/bio426/gendor/internal/core"
)

type WorkshopCtl core.Controller

func (ctl *WorkshopCtl) List(c echo.Context) error {
	err := Service.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "")
}

func (ctl *WorkshopCtl) Create(c echo.Context) error {
	body := struct {
		Name        string `json:"name" validate:"required"`
		Address     string `json:"address" validate:"required"`
		Dni         string `json:"dni" validate:"required"`
		Ruc         string `json:"ruc" validate:"required"`
		Brand       string `json:"brand" validate:"required"`
		Model       string `json:"model" validate:"required"`
		Color       string `json:"color" validate:"required"`
		Plate       string `json:"plate" validate:"required"`
		Mileage     string `json:"mileage" validate:"required"`
		Observation string `json:"observation" validate:"required"`
		// Items       []SvcCreateItem `json:"Items" validate:"required"`
		Items []struct {
			Code        string  `json:"code" validate:"required"`
			Quantity    int32   `json:"quantity" validate:"required"`
			Price       float32 `json:"price" validate:"required"`
			Description string  `json:"description" validate:"required"`
		} `json:"items" validate:"required,min=1"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	userId := c.Get(auth.CtxUserIdKey).(int32)

	parsedItems := []SvcCreateItem{}
	for _, item := range body.Items {
		parsedItems = append(parsedItems, SvcCreateItem{
			Code:        item.Code,
			Quantity:    item.Quantity,
			Price:       item.Price,
			Description: item.Description,
		})
	}

	err := Service.Create(c.Request().Context(), SvcCreateParams{
		Name:        body.Name,
		Address:     body.Address,
		Dni:         body.Dni,
		Ruc:         body.Ruc,
		Brand:       body.Brand,
		Model:       body.Model,
		Color:       body.Color,
		Plate:       body.Plate,
		Mileage:     body.Mileage,
		Observation: body.Observation,
		Items:       parsedItems,
		UserId:      userId,
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

type CtlSearchByPlateResponse struct {
	Id          string `json:"id,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	Address     string `json:"address,omitempty" validate:"required"`
	Dni         string `json:"dni,omitempty" validate:"required"`
	Ruc         string `json:"ruc,omitempty" validate:"required"`
	Brand       string `json:"brand,omitempty" validate:"required"`
	Model       string `json:"model,omitempty" validate:"required"`
	Color       string `json:"color,omitempty" validate:"required"`
	Mileage     string `json:"mileage,omitempty" validate:"required"`
	Observation string `json:"observation,omitempty" validate:"required"`
}

func (ctl *WorkshopCtl) SearchByPlate(c echo.Context) error {
	query := struct {
		Plate string `query:"plate" validate:"required"`
	}{}
	if err := c.Bind(&query); err != nil {
		return err
	}
	if err := c.Validate(query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := Service.Create(c.Request().Context(), SvcCreateParams{})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

var Controller = &WorkshopCtl{}
