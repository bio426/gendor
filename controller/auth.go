package controller

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/bio426/gendor/datasource"
	"github.com/labstack/echo/v4"
)

const tokenName = "session-token"

func AuthLogin(c echo.Context) error {
	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	// check user and password
	var (
		id       int32
		password string
	)
	if err := datasource.Postgres.QueryRowContext(
		c.Request().Context(),
		"select u.id, u.username from users u where u.username = $1",
		body.Username,
	).Scan(&id, &password); err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return err
	}

	// delete existent sessions of user
	if _, err := datasource.Postgres.ExecContext(
		c.Request().Context(),
		"delete from users u where u.user_id = $1",
		id); err != nil {
		return err
	}

	// create session
	expiresAt := time.Now().Add(time.Second * 10)
	cookie := new(http.Cookie)
	cookie.Name = tokenName
	cookie.Expires = expiresAt
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func AuthLogout(c echo.Context) error {

	// create session
	// cookie := new(http.Cookie)
	cookie := &http.Cookie{}
	cookie.Name = tokenName
	cookie.Value = ""
	cookie.Expires = time.Now()
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
