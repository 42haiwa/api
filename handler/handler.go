package handler

import (
	user "api/data"
	"api/manager"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func SignupHandler(c echo.Context) error {
	u := new(user.User)

	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "Error Post")
	}
	connString := "postgres://postgres:password@localhost:5432/utrade"

	log.Println("before parsing")

	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("after conn")
	defer conn.Close()

	manager.SignupRequest(conn, *u, context.Background())

	return c.String(http.StatusOK, "Success Post")
}
