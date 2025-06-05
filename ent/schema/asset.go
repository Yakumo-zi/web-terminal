package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Asset holds the schema definition for the Asset entity.
type Asset struct {
	ent.Schema
}

// Fields of the Asset.
func (Asset) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("type"),
		field.String("name").Unique(),
		field.String("ip").Unique(),
		field.Int16("port"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the Asset.
func (Asset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", AssetGroup.Type).
			Ref("assets"),
		edge.To("credentials", Credential.Type),
	}
}
