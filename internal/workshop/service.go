package workshop

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bio426/gendor/datasource"
	"github.com/bio426/gendor/internal/core"
	"github.com/bio426/gendor/util"
)

type WorkshopSvc core.Service

type SvcListParams struct {
	Search string
	Page   int32
}

func (svc *WorkshopSvc) List(c context.Context, params SvcListParams) (*CtlListResponse, error) {
	var totalRows int32
	row := datasource.Postgres.QueryRowContext(
		c,
		"select count(*) from workshop_orders where propietary ilike $1",
		fmt.Sprintf("%%%s%%", params.Search),
	)
	if err := row.Scan(&totalRows); err != nil {
		return nil, err
	}

	rows, err := datasource.Postgres.QueryContext(
		c,
		"select id,propietary,plate,created_at from workshop_orders where plate ilike $1 order by created_at desc limit 20 offset $2",
		fmt.Sprintf("%%%s%%", params.Search),
		(params.Page-1)*20,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resRows := []CtlListRow{}
	for rows.Next() {
		var row = CtlListRow{}
		if err = rows.Scan(
			&row.Id,
			&row.Propietary,
			&row.Plate,
			&row.CreatedAt,
		); err != nil {
			return nil, err
		}
		resRows = append(resRows, row)
	}

	from, to := util.DataFromTo(totalRows, params.Page, 20)
	res := &CtlListResponse{
		Total: totalRows,
		From:  from,
		To:    to,
		Rows:  resRows,
	}

	return res, nil
}

func (svc *WorkshopSvc) Detail(c context.Context, id int32) (*CtlDetailResponse, error) {
	tx, err := datasource.Postgres.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res := &CtlDetailResponse{Items: []CtlDetailItem{}}
	// get order
	row := datasource.Postgres.QueryRowContext(
		c,
		"select id,propietary,created_at,brand,model,car_year,plate,mileage,observation,discount from workshop_orders where id = $1",
		id,
	)
	if err := row.Scan(
		&res.Id,
		&res.Propietary,
		&res.CreatedAt,
		&res.Brand,
		&res.Model,
		&res.Year,
		&res.Plate,
		&res.Mileage,
		&res.Observation,
		&res.Discount,
	); err != nil {
		return nil, err
	}

	rows, err := datasource.Postgres.QueryContext(
		c,
		"select price,description from workshop_order_items where orderId = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row = CtlDetailItem{}
		if err = rows.Scan(
			&row.Price,
			&row.Description,
		); err != nil {
			return nil, err
		}
		res.Items = append(res.Items, row)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return res, nil
}

type SvcCreateParams struct {
	Propietary  string
	Brand       string
	Model       string
	Year        int32
	Plate       string
	Mileage     int32
	Observation string
	Discount    float32
	Items       []SvcCreateItem
	UserId      int32
}
type SvcCreateItem struct {
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
          propietary, brand, model, car_year, 
          plate, mileage, observation, discount, userId
        ) 
        values 
          ($1,$2,$3,$4,$5,$6,$7,$8,$9)
        returning id
        `,
		params.Propietary, params.Brand, params.Model,
		params.Year, params.Plate, params.Mileage,
		params.Observation, params.Discount, params.UserId,
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
          price, description, orderId
        ) 
        values 
          ($1,$2,$3)
        `,
	)
	if err != nil {
		return err
	}
	for _, item := range params.Items {
		_, err = stmt.ExecContext(c, item.Price, item.Description, orderId)
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
          propietary, brand, 
          model, car_year, mileage 
        from 
          workshop_orders 
        where 
          plate ilike $1
        order by 
          created_at desc 
        limit 
          1
        `,
		fmt.Sprintf("%%%s%%", plate),
	)
	if err := row.Scan(
		&res.Propietary,
		&res.Brand, &res.Model, &res.Year, &res.Mileage,
	); err != nil {
		if err == sql.ErrNoRows {
			return &CtlSearchByPlateResponse{}, nil
		}
		return nil, err
	}

	return res, nil

}

var Service = &WorkshopSvc{}
