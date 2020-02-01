// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/jaredallard/grocerylistsbot/ent/groceryitem"
)

// GroceryItem is the model entity for the GroceryItem schema.
type GroceryItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Status holds the value of the "status" field.
	Status groceryitem.Status `json:"status,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroceryItemQuery when eager-loading is set.
	Edges                       GroceryItemEdges `json:"edges"`
	grocery_item_grocerylist_id *int
}

// GroceryItemEdges holds the relations/edges for other nodes in the graph.
type GroceryItemEdges struct {
	// Grocerylist holds the value of the grocerylist edge.
	Grocerylist *GroceryList
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroceryItem) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullTime{},    // created_at
		&sql.NullTime{},    // modified_at
		&sql.NullTime{},    // deleted_at
		&sql.NullString{},  // name
		&sql.NullString{},  // status
		&sql.NullFloat64{}, // price
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*GroceryItem) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // grocery_item_grocerylist_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroceryItem fields.
func (gi *GroceryItem) assignValues(values ...interface{}) error {
	if m, n := len(values), len(groceryitem.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	gi.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		gi.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field modified_at", values[1])
	} else if value.Valid {
		gi.ModifiedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field deleted_at", values[2])
	} else if value.Valid {
		gi.DeletedAt = new(time.Time)
		*gi.DeletedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[3])
	} else if value.Valid {
		gi.Name = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[4])
	} else if value.Valid {
		gi.Status = groceryitem.Status(value.String)
	}
	if value, ok := values[5].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field price", values[5])
	} else if value.Valid {
		gi.Price = value.Float64
	}
	values = values[6:]
	if len(values) == len(groceryitem.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field grocery_item_grocerylist_id", value)
		} else if value.Valid {
			gi.grocery_item_grocerylist_id = new(int)
			*gi.grocery_item_grocerylist_id = int(value.Int64)
		}
	}
	return nil
}

// QueryGrocerylist queries the grocerylist edge of the GroceryItem.
func (gi *GroceryItem) QueryGrocerylist() *GroceryListQuery {
	return (&GroceryItemClient{gi.config}).QueryGrocerylist(gi)
}

// Update returns a builder for updating this GroceryItem.
// Note that, you need to call GroceryItem.Unwrap() before calling this method, if this GroceryItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (gi *GroceryItem) Update() *GroceryItemUpdateOne {
	return (&GroceryItemClient{gi.config}).UpdateOne(gi)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gi *GroceryItem) Unwrap() *GroceryItem {
	tx, ok := gi.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroceryItem is not a transactional entity")
	}
	gi.config.driver = tx.drv
	return gi
}

// String implements the fmt.Stringer.
func (gi *GroceryItem) String() string {
	var builder strings.Builder
	builder.WriteString("GroceryItem(")
	builder.WriteString(fmt.Sprintf("id=%v", gi.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(gi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", modified_at=")
	builder.WriteString(gi.ModifiedAt.Format(time.ANSIC))
	if v := gi.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(gi.Name)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", gi.Status))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", gi.Price))
	builder.WriteByte(')')
	return builder.String()
}

// GroceryItems is a parsable slice of GroceryItem.
type GroceryItems []*GroceryItem

func (gi GroceryItems) config(cfg config) {
	for _i := range gi {
		gi[_i].config = cfg
	}
}