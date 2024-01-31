package connect

import "github.com/jackc/pgx/v5/pgxpool"

type Connect struct {
	DBPostr *pgxpool.Pool
}
