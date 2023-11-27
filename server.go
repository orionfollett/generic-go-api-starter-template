package main

import (
	"log"
	"net/http"
	"os"

	graph "orion/generic-api-starter/api/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

//this could be good and modified so codegen automatically runs: go:generate go run golang.org/x/tools/cmd/stringer -type=Pill

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
