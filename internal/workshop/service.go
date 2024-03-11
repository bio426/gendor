package workshop

import (
	"context"

	"github.com/bio426/gendor/datasource"
	"github.com/bio426/gendor/internal/core"
)

type AuthSvc core.Service

func (svc *AuthSvc) List(c context.Context) error {
	datasource.Postgres.Query("")
	return nil
}

var Service = &AuthSvc{}
