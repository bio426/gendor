package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/bio426/gendor/datasource"
	"github.com/bio426/gendor/internal/core"
)

type UserSvc core.Service

func (svc *UserSvc) List(c context.Context) (*CtlListResponse, error) {
	rows, err := datasource.Postgres.QueryContext(
		c,
		"select id,username,role,active,created_at from users where role != 'super'",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := &CtlListResponse{Rows: []CtlListRow{}}
	for rows.Next() {
		var row = CtlListRow{}
		if err = rows.Scan(
			&row.Id,
			&row.Username,
			&row.Role,
			&row.Active,
			&row.CreatedAt,
		); err != nil {
			return nil, err
		}
		res.Rows = append(res.Rows, row)
	}
	return res, nil
}

type SvcCreateParams struct {
	Username string
	Password string
}

func (svc *UserSvc) Create(c context.Context, params SvcCreateParams) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		return err
	}
	_, err = datasource.Postgres.ExecContext(
		c,
		"insert into users(username,password,role) values ($1,$2,$3)",
		params.Username, hashedPassword, "admin",
	)
	if err != nil {
		return err
	}

	return nil
}

var Service = &UserSvc{}
