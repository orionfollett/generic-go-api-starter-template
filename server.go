package main

import (
	"log"
	"net/http"
	"os"

	generated "orion/generic-api-starter/api/generated"
	resolvers "orion/generic-api-starter/api/resolvers"

	"context"

	sqlc "orion/generic-api-starter/db/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx/v5"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()

	conn, _ := pgx.Connect(ctx, "postgres://postgres:example@localhost:5432")

	defer conn.Close(ctx)

	queries := sqlc.New(conn)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{QUERIES: queries}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
