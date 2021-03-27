package main

import (
	"demographql/database"
	"demographql/graph"
	"demographql/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/gorm"
)

const defaultPort = "8080"

var db *gorm.DB

func InitDB() {
	dbOut, err := database.OpenGormOutput()

	if err != nil {
		log.Panic("Connect Database Error.")
	}

	db = dbOut
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	InitDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
