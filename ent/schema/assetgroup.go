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

// AssetGroup holds the schema definition for the AssetGroup entity.
type AssetGroup struct {
	ent.Schema
}

// Fields of the AssetGroup.
func (AssetGroup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("name").Unique(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the AssetGroup.
func (AssetGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("assets", Asset.Type),
		edge.To("attrs", AssetGroupAttribute.Type),
	}
}
func (AssetGroup) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.AssetGroupFunc(func(ctx context.Context, m *gen.AssetGroupMutation) (gen.Value, error) {
				m.SetUpdatedAt(time.Now())
				return next.Mutate(ctx, m)
			})
		}, ent.OpUpdate|ent.OpUpdateOne),
	}
}
