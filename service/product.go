package service

import (
	"context"
	"database/sql"

	"github.com/bio426/gendor/datasource"
	"github.com/cloudinary/cloudinary-go/v2"
)

type ProductService struct {
	database *sql.DB
	storage  *cloudinary.Cloudinary
	config   datasource.Configuration
}

func (s ProductService) CreateProductCategory(c context.Context, name string) error {
	if _, err := datasource.Postgres.ExecContext(
		c,
		"insert into product_categories(name) values ($1)",
		name,
	); err != nil {
		return err
	}
	return nil
}
