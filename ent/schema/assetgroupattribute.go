package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AssetGroupAttribute holds the schema definition for the AssetGroupAttribute entity.
type AssetGroupAttribute struct {
	ent.Schema
}

// Fields of the AssetGroupAttribute.
func (AssetGroupAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("key"),
		field.String("value"),
		field.String("type"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the AssetGroupAttribute.
func (AssetGroupAttribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", AssetGroup.Type).
			Ref("attrs").
			Unique(),
	}
}
