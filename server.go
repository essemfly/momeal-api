package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"lessbutter.co/mealkit/config"
	"lessbutter.co/mealkit/src/generated"
	"lessbutter.co/mealkit/src/resolver"
)

func main() {
	conf := config.GetConfiguration()

	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(conf.PORT)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
