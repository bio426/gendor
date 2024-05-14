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
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Plate     string    `json:"plate"`
	CreatedAt time.Time `json:"createdAt"`
}

// type CtlListResponse struct {
// 	Rows []CtlListRow `json:"rows"`
// }

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
	Code        string  `json:"code"`
	Quantity    int32   `json:"quantity"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}
type CtlDetailResponse struct {
	Id          int32           `json:"id"`
	Name        string          `json:"name"`
	Address     string          `json:"address"`
	Dni         string          `json:"dni"`
	Ruc         string          `json:"ruc"`
	CreatedAt   time.Time       `json:"createdAt"`
	Brand       string          `json:"brand"`
	Model       string          `json:"model"`
	Color       string          `json:"color"`
	Plate       string          `json:"plate"`
	Mileage     int32           `json:"mileage"`
	Observation string          `json:"observation"`
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
		Name        string `json:"name"`
		Address     string `json:"address"`
		Dni         string `json:"dni"`
		Ruc         string `json:"ruc"`
		Brand       string `json:"brand"`
		Model       string `json:"model"`
		Color       string `json:"color"`
		Plate       string `json:"plate" validate:"required"`
		Mileage     int32  `json:"mileage"`
		Observation string `json:"observation"`
		Items       []struct {
			Code        string  `json:"code"`
			Quantity    int32   `json:"quantity"`
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
	Name    string `json:"name,omitempty" validate:"required"`
	Address string `json:"address,omitempty" validate:"required"`
	Dni     string `json:"dni,omitempty" validate:"required"`
	Ruc     string `json:"ruc,omitempty" validate:"required"`
	Brand   string `json:"brand,omitempty" validate:"required"`
	Model   string `json:"model,omitempty" validate:"required"`
	Color   string `json:"color,omitempty" validate:"required"`
	Mileage int32  `json:"mileage,omitempty" validate:"required"`
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
