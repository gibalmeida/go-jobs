package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/gibalmeida/go-jobs/ent"
	"github.com/gibalmeida/go-jobs/graph/generated"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
