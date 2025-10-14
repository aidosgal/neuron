package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/aidosgal/neuron/pkg/gen"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default((func() uuid.UUID)(gen.UUID())),
		field.String("device_name").
			NotEmpty(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return nil
}
