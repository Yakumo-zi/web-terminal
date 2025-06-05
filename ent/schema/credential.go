package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Credential holds the schema definition for the Credential entity.
type Credential struct {
	ent.Schema
}

// Fields of the Credential.
func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("username"),
		field.String("secret"),
		field.String("type"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the Credential.
func (Credential) Edges() []ent.Edge {
	return nil
}
