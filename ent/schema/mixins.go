package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type CRIFields struct{}

func (CRIFields) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("modified_at").Default(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

func (CRIFields) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (CRIFields) Hooks() []ent.Hook {
	return []ent.Hook{}
}

func (CRIFields) Indexes() []ent.Index {
	return []ent.Index{}
}
