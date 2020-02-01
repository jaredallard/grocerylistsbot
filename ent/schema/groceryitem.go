package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// GroceryItem holds the schema definition for the GroceryItem entity.
type GroceryItem struct {
	ent.Schema
}

// Fields of the GroceryItem.
func (GroceryItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("status").Values("purchased", "unpurchased").Default("unpurchased"),
		field.Float("price").Optional(),
	}
}

// Edges of the GroceryItem.
func (GroceryItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("grocerylist", GroceryList.Type).Unique(),
	}
}

func (GroceryItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CRIFields{},
	}
}
