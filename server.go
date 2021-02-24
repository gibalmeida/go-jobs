package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gibalmeida/go-jobs/ent"
	"github.com/gibalmeida/go-jobs/ent/migrate"
	"github.com/gibalmeida/go-jobs/graph"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

// Defining the Graphql handler
func graphqlHandler(client *ent.Client) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewSchema(client))
	h.Use(entgql.Transactioner{TxOpener: client})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("mysql", "root:dbpass@tcp(localhost:3306)/jobs?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to MySQL server: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	// Run migration.
	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphqlHandler(client))
	r.GET("/", playgroundHandler())
	r.Run(":" + port)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}
