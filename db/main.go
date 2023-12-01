package main

import (
	"context"
	"log"
	sqlc "orion/generic-api-starter/db/generated"
	"reflect"

	"github.com/jackc/pgx/v5"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:example@localhost:5432")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := sqlc.New(conn)

	// create a run
	insertedRun, err := queries.CreateRun(ctx, "run1")
	if err != nil {
		return err
	}
	log.Println(insertedRun)

	// get the run we just inserted
	fetchedRun, err := queries.GetRun(ctx, insertedRun.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedRun, fetchedRun))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
