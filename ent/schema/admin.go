package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("login").
			Unique().NotEmpty(),
		field.String("password").
			Sensitive(),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return nil
}
