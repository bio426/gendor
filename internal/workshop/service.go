package workshop

import (
	"context"
	"database/sql"

	"github.com/bio426/gendor/datasource"
	"github.com/bio426/gendor/internal/core"
)

type WorkshopSvc core.Service

func (svc *WorkshopSvc) List(c context.Context) error {
	datasource.Postgres.Query("")
	return nil
}

type SvcCreateParams struct {
	Name        string
	Address     string
	Dni         string
	Ruc         string
	Brand       string
	Model       string
	Color       string
	Plate       string
	Mileage     string
	Observation string
	Items       []SvcCreateItem
	UserId      int32
}
type SvcCreateItem struct {
	Code        string
	Quantity    int32
	Price       float32
	Description string
}

func (svc *WorkshopSvc) Create(c context.Context, params SvcCreateParams) error {
	tx, err := datasource.Postgres.BeginTx(c, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// insert order
	row := tx.QueryRowContext(
		c,
		`
        insert into workshop_orders(
          name, address, dni, ruc, brand, model, 
          color, plate, mileage, observation
        ) 
        values 
          ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
        returning id
        `,
		params.Name, params.Address, params.Dni,
		params.Ruc, params.Brand, params.Model,
		params.Color, params.Plate, params.Mileage,
		params.Observation,
	)
	var orderId int32
	if err := row.Scan(&orderId); err != nil {
		return err
	}

	// insert items
	stmt, err := tx.PrepareContext(
		c,
		`
        insert into workshop_order_items(
          code, quantity, price, description, orderId
        ) 
        values 
          ($1,$2,$3,$4,$5)
        `,
	)
	if err != nil {
		return err
	}
	for _, item := range params.Items {
		_, err = stmt.ExecContext(c, item.Code, item.Quantity, item.Price, item.Description, orderId)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (svc *WorkshopSvc) SearchByPlate(c context.Context, plate string) (*CtlSearchByPlateResponse, error) {
	res := &CtlSearchByPlateResponse{}
	row := datasource.Postgres.QueryRowContext(c,
		`
        select 
          name, address, dni, ruc, brand, 
          model, color, mileage, observation 
        from 
          workshop_orders 
        where 
          plate = $1 
        order by 
          created_at desc 
        limit 
          1
        `,
		plate,
	)
	if err := row.Scan(
		&res.Name, &res.Address, &res.Dni, &res.Ruc,
		&res.Brand, &res.Model, &res.Color, &res.Mileage,
		&res.Observation,
	); err != nil {
		if err == sql.ErrNoRows {
			return &CtlSearchByPlateResponse{}, nil
		}
		return nil, err
	}

	return res, nil

}

var Service = &WorkshopSvc{}
