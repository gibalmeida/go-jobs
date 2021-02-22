package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MinLen(2).
			MaxLen(50).
			Annotations(entgql.OrderField("NAME")),
		field.String("email").
			NotEmpty().
			Annotations(entgql.OrderField("EMAIL")),
		field.String("password").
			NotEmpty(),
		field.Enum("role").
			NamedValues(
				"admin", "ADMIN",
				"manager", "MANAGER",
				"applicant", "APPLICANT",
			).
			Annotations(entgql.OrderField("ROLE")),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("departments", Department.Type).
			Annotations(entgql.Bind()),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
