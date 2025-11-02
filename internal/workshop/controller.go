package workshop

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/bio426/gendor/internal/auth"
	"github.com/bio426/gendor/internal/core"
)

type WorkshopCtl core.Controller

type CtlListRow struct {
	Id         int32     `json:"id"`
	Propietary string    `json:"propietary"`
	Plate      string    `json:"plate"`
	CreatedAt  time.Time `json:"createdAt"`
}

type CtlListResponse core.PaginatedResponse[CtlListRow]

func (ctl *WorkshopCtl) List(c echo.Context) error {
	query := struct {
		Search string `query:"search" `
		Page   int32  `query:"page" validate:"required"`
	}{}
	if err := c.Bind(&query); err != nil {
		return err
	}
	if err := c.Validate(query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := Service.List(c.Request().Context(), SvcListParams{Search: query.Search, Page: query.Page})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

type CtlDetailItem struct {
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}
type CtlDetailResponse struct {
	Id          int32           `json:"id"`
	Propietary  string          `json:"propietary"`
	CreatedAt   time.Time       `json:"createdAt"`
	Brand       string          `json:"brand"`
	Model       string          `json:"model"`
	Year        int32           `json:"year"`
	Plate       string          `json:"plate"`
	Mileage     string          `json:"mileage"`
	Observation string          `json:"observation"`
	Discount    float32         `json:"discount"`
	Items       []CtlDetailItem `json:"items,omitempty"`
}

func (ctl *WorkshopCtl) Detail(c echo.Context) error {
	params := struct {
		Id int32 `param:"id" validate:"required"`
	}{}
	if err := c.Bind(&params); err != nil {
		return err
	}
	if err := c.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := Service.Detail(c.Request().Context(), params.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *WorkshopCtl) Create(c echo.Context) error {
	body := struct {
		Propietary  string  `json:"propietary"`
		Brand       string  `json:"brand"`
		Model       string  `json:"model"`
		Year        int32   `json:"year"`
		Plate       string  `json:"plate" validate:"required"`
		Mileage     int32   `json:"mileage"`
		Observation string  `json:"observation"`
		Discount    float32 `json:"discount"`
		Items       []struct {
			Price       float32 `json:"price" validate:"required"`
			Description string  `json:"description" validate:"required"`
		} `json:"items" validate:"required,min=1"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	userId := c.Get(auth.CtxUserIdKey).(int32)

	parsedItems := []SvcCreateItem{}
	for _, item := range body.Items {
		parsedItems = append(parsedItems, SvcCreateItem{
			Price:       item.Price,
			Description: item.Description,
		})
	}

	err := Service.Create(c.Request().Context(), SvcCreateParams{
		Propietary:  body.Propietary,
		Brand:       body.Brand,
		Model:       body.Model,
		Year:        body.Year,
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
	Propietary string `json:"propietary,omitempty" validate:"required"`
	Brand      string `json:"brand,omitempty" validate:"required"`
	Model      string `json:"model,omitempty" validate:"required"`
	Year       int32  `json:"year,omitempty" validate:"required"`
	Mileage    int32  `json:"mileage,omitempty" validate:"required"`
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

	res, err := Service.SearchByPlate(c.Request().Context(), query.Plate)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

var Controller = &WorkshopCtl{}
