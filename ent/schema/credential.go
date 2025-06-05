package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	gen "github.com/Yakumo-zi/web-terminal/ent"
	"github.com/Yakumo-zi/web-terminal/ent/hook"
	"github.com/google/uuid"
	"time"
)

// Credential holds the schema definition for the Credential entity.
type Credential struct {
	ent.Schema
}

// Fields of the Credential.
func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("username"),
		field.String("secret"),
		field.String("type"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Credential.
func (Credential) Edges() []ent.Edge {
	return nil
}

func (Credential) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.CredentialFunc(func(ctx context.Context, m *gen.CredentialMutation) (gen.Value, error) {
				m.SetUpdatedAt(time.Now())
				return next.Mutate(ctx, m)
			})
		}, ent.OpUpdate|ent.OpUpdateOne),
	}
}
