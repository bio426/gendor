package controller

import (
	// "database/sql"
	"net/http"
	// "time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/bio426/gendor/datasource"
	"github.com/labstack/echo/v4"
)

type controller struct {
	is string
}

var Products = controller{is: "Ctr"}

func (ctl controller) Create(c echo.Context) error {
	form := struct {
		Name  string  `form:"name" validate:"required"`
		Price float32 `form:"price" validate:"required|number"`
		// Image int32 `form:"image" validate:"required"`
	}{}
	if err := c.Bind(&form); err != nil {
		return err
	}
	if err := c.Validate(form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Upload image
	res, err := datasource.Cld.Upload.Upload(
		c.Request().Context(),
		src,
		uploader.UploadParams{
			ResourceType: "image",
			Folder:       "gendor/products",
		},
	)
	if err != nil {
		return err
	}

	// Create db register
	if _, err := datasource.Postgres.ExecContext(
		c.Request().Context(),
		"insert into products(name,price,image) values ($1,$2,$3)",
		form.Name,
		form.Price,
		res.SecureURL,
	); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (ctl controller) List(c echo.Context) error {
	query := struct {
		Page   int32  `query:"page" validate:"required|number"`
		Count  int32  `query:"count" validate:"required|number"`
		Search string `query:"search"`
	}{}
	if err := c.Bind(&query); err != nil {
		return err
	}
	if err := c.Validate(query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Query data
	var rowCount int32 = 1
	rows, err := datasource.Postgres.QueryContext(
		c.Request().Context(),
		"select p.id, p.name, p.price, p.image from products p",
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	type row = struct {
		Id    int32   `json:"id"`
		Name  string  `json:"name"`
		Price float32 `json:"price"`
		Image string  `json:"image"`
	}

	data := []row{}
	for rows.Next() {
		var row = row{}
		if err = rows.Scan(
			&row.Id,
			&row.Name,
			&row.Price,
			&row.Image,
		); err != nil {
			return err
		}
		data = append(data, row)
	}
	res := struct {
		Total int32 `json:"total"`
		Rows  []row `json:"rows"`
	}{
		Total: rowCount,
		Rows:  data,
	}

	return c.JSON(http.StatusCreated, res)
}

func (ctl controller) CategoryList(c echo.Context) error {
	rows, err := datasource.Postgres.QueryContext(
		c.Request().Context(),
		"select pc.id, pc.name from product_categories pc",
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	type row = struct {
		Id   int32  `json:"id"`
		Name string `json:"name"`
	}

	data := []row{}
	for rows.Next() {
		var row = row{}
		if err = rows.Scan(
			&row.Id,
			&row.Name,
		); err != nil {
			return err
		}
		data = append(data, row)
	}
	res := struct {
		Rows []row `json:"rows"`
	}{
		Rows: data,
	}

	return c.JSON(http.StatusCreated, res)
}

func (ctl controller) CategoryCreate(c echo.Context) error {
	body := struct {
		Name string `json:"name" validate:"required"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if _, err := datasource.Postgres.ExecContext(
		c.Request().Context(),
		"insert into product_categories(name) values ($1)",
		body.Name,
	); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
