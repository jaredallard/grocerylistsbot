package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// GroceryList holds the schema definition for the GroceryList entity.
type GroceryList struct {
	ent.Schema
}

// Fields of the GroceryList.
func (GroceryList) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("Unset"),
	}
}

// Edges of the GroceryList.
func (GroceryList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("members", User.Type).Ref("grocerylist"),
	}
}

func (GroceryList) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CRIFields{},
	}
}
