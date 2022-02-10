package main

import (
	"log"
	"net/http"
	"os"

	"feldrise.com/inventory-exercice/graph"
	"feldrise.com/inventory-exercice/graph/generated"
	"feldrise.com/inventory-exercice/internal/config"
	"feldrise.com/inventory-exercice/internal/database"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config.Init("config.yml")
	database.Init()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
