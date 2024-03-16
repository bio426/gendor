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
		"select count(*) from workshop_orders where name ilike $1",
		fmt.Sprintf("%%%s%%", params.Search),
	)
	if err := row.Scan(&totalRows); err != nil {
		return nil, err
	}

	rows, err := datasource.Postgres.QueryContext(
		c,
		"select id,name,plate,created_at from workshop_orders where plate ilike $1 order by created_at desc limit 20 offset $2",
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
			&row.Name,
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
		"select id,name,address,dni,ruc,created_at,brand,model,color,plate,mileage,observation from workshop_orders where id = $1",
		id,
	)
	if err := row.Scan(
		&res.Id,
		&res.Name,
		&res.Address,
		&res.Dni,
		&res.Ruc,
		&res.CreatedAt,
		&res.Brand,
		&res.Model,
		&res.Color,
		&res.Plate,
		&res.Mileage,
		&res.Observation,
	); err != nil {
		return nil, err
	}

	rows, err := datasource.Postgres.QueryContext(
		c,
		"select code,quantity,price,description from workshop_order_items where orderId = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row = CtlDetailItem{}
		if err = rows.Scan(
			&row.Code,
			&row.Quantity,
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
	Name        string
	Address     string
	Dni         string
	Ruc         string
	Brand       string
	Model       string
	Color       string
	Plate       string
	Mileage     int32
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
          color, plate, mileage, observation, userId
        ) 
        values 
          ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
        returning id
        `,
		params.Name, params.Address, params.Dni,
		params.Ruc, params.Brand, params.Model,
		params.Color, params.Plate, params.Mileage,
		params.Observation, params.UserId,
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
          model, color, mileage 
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
		&res.Name, &res.Address, &res.Dni, &res.Ruc,
		&res.Brand, &res.Model, &res.Color, &res.Mileage,
	); err != nil {
		if err == sql.ErrNoRows {
			return &CtlSearchByPlateResponse{}, nil
		}
		return nil, err
	}

	return res, nil

}

var Service = &WorkshopSvc{}
