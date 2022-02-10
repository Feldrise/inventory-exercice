package main

import (
	"log"
	"net/http"
	"os"

	"feldrise.com/inventory-exercice/graph"
	"feldrise.com/inventory-exercice/graph/generated"
	"feldrise.com/inventory-exercice/internal/auth"
	"feldrise.com/inventory-exercice/internal/config"
	"feldrise.com/inventory-exercice/internal/database"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	config.Init("config.yml")
	database.Init()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
