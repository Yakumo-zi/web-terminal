package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AssetGroup holds the schema definition for the AssetGroup entity.
type AssetGroup struct {
	ent.Schema
}

// Fields of the AssetGroup.
func (AssetGroup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("name"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the AssetGroup.
func (AssetGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("assets", Asset.Type),
		edge.To("attrs", AssetGroupAttribute.Type),
	}
}
