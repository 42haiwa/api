package manager

import (
	user "api/data"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func SignupRequest(conn *pgxpool.Pool, u user.User, ctx context.Context) error {
	query := `INSERT INTO users (uuid, username, email, first_name, last_name, password, subscription_status) VALUES (@uuid, @username, @email, @first_name, @last_name, @password, @subscription_status)`

	log.Println("val", u)
	args := pgx.NamedArgs{
		"uuid":                uuid.New(),
		"username":            u.Username,
		"email":               u.Email,
		"first_name":          u.Firstname,
		"last_name":           u.Lastname,
		"password":            u.Password,
		"subscription_status": 1,
	}

	_, err := conn.Exec(ctx, query, args)
	if err != nil {
		log.Fatal(err)
	}
	fetchDB(conn)
	return nil
}

func fetchDB(conn *pgxpool.Pool) bool {
	query := `SELECT * FROM users`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("rows len", rows)

	return false
}
