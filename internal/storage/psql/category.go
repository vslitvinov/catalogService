package psql

import "github.com/jackc/pgx/v4/pgxpool"


type Account struct {
	pool  *pgxpool.Pool
}