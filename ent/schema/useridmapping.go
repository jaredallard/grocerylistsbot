package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// UserIDMapping holds the schema definition for the UserIDMapping entity.
type UserIDMapping struct {
	ent.Schema
}

// Fields of the UserIDMapping.
func (UserIDMapping) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("platform_type").Values("telegram", "sms"),
		field.String("platform_id"),
	}
}

// Edges of the UserIDMapping.
func (UserIDMapping) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
	}
}

func (UserIDMapping) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("platform_type", "platform_id").Unique(),
	}
}

func (UserIDMapping) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CRIFields{},
	}
}
