package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Earl-Power/gqlgengin/graph"
	generated "github.com/Earl-Power/gqlgengin/graph"
	"github.com/Earl-Power/gqlgengin/internal/auth"
	database "github.com/Earl-Power/gqlgengin/internal/pkg/db/mysql"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}),
	)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

/*
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
*/
