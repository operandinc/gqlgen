package main

import (
	"log"
	"net/http"

	"github.com/operandinc/gqlgen/_examples/selection"
	"github.com/operandinc/gqlgen/graphql/handler"
	"github.com/operandinc/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
