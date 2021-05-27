package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/lessbutter/momeal-api/config"
	"github.com/lessbutter/momeal-api/database"
	"github.com/lessbutter/momeal-api/src/generated"
	"github.com/lessbutter/momeal-api/src/resolver"
)

func main() {
	conf := config.GetConfiguration()

	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(conf.PORT)
	}

	database.InitDB(conf)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
