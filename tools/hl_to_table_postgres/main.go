package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var tableName = "test"

func main() {
	fmt.Println("start hl:")

	connPgx, err := pgxpool.New(context.Background(), "postgres://postgres:123456@postgres-16-1-master:5432/postgres")

	if err != nil {
		log.Printf("Unable to connect to database master: %v\n", err)

	}

	countSqlInsert := 100000

	for i := 0; i < countSqlInsert; i++ {
		query := `INSERT INTO ` + tableName + `(text) values('testweqeeqeqweeasw9w0w3545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn89738973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf8973545ydfggbn8973545ydfggbn8973545ydf545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn8973545ydfggbn897') RETURNING id;`

		var id int
		err = connPgx.QueryRow(context.Background(), query).Scan(&id)

		if err != nil {
			log.Printf("db error: %v\n", err)

		}

		fmt.Println("id:", id)
	}

}
