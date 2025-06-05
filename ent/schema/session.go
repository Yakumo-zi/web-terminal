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

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("status"),
		field.String("type"),
		field.Time("created_at"),
		field.Time("updated_at").Default(time.Now),
		field.Time("stoped_at").Default(time.Now),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("asset", Asset.Type).Unique(),
		edge.To("credential", Credential.Type).Unique(),
	}
}

func (Session) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.SessionFunc(func(ctx context.Context, m *gen.SessionMutation) (gen.Value, error) {
				m.SetUpdatedAt(time.Now())
				return next.Mutate(ctx, m)
			})
		}, ent.OpUpdate|ent.OpUpdateOne),
	}
}
