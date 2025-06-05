package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/Yakumo-zi/web-terminal/ent"
	"github.com/Yakumo-zi/web-terminal/ent/hook"
	"github.com/google/uuid"
	"time"
)

// AssetGroupAttribute holds the schema definition for the AssetGroupAttribute entity.
type AssetGroupAttribute struct {
	ent.Schema
}

// Fields of the AssetGroupAttribute.
func (AssetGroupAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("key"),
		field.String("value"),
		field.String("type"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
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
func (AssetGroupAttribute) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.AssetGroupAttributeFunc(func(ctx context.Context, m *gen.AssetGroupAttributeMutation) (gen.Value, error) {
				m.SetUpdatedAt(time.Now())
				return next.Mutate(ctx, m)
			})
		}, ent.OpUpdate|ent.OpUpdateOne),
	}
}
