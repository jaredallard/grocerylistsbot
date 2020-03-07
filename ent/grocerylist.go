// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/jaredallard/grocerylistsbot/ent/grocerylist"
)

// GroceryList is the model entity for the GroceryList schema.
type GroceryList struct {
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroceryListQuery when eager-loading is set.
	Edges GroceryListEdges `json:"edges"`
}

// GroceryListEdges holds the relations/edges for other nodes in the graph.
type GroceryListEdges struct {
	// Members holds the value of the members edge.
	Members []*User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e GroceryListEdges) MembersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroceryList) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // modified_at
		&sql.NullTime{},   // deleted_at
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroceryList fields.
func (gl *GroceryList) assignValues(values ...interface{}) error {
	if m, n := len(values), len(grocerylist.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	gl.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		gl.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field modified_at", values[1])
	} else if value.Valid {
		gl.ModifiedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field deleted_at", values[2])
	} else if value.Valid {
		gl.DeletedAt = new(time.Time)
		*gl.DeletedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[3])
	} else if value.Valid {
		gl.Name = value.String
	}
	return nil
}

// QueryMembers queries the members edge of the GroceryList.
func (gl *GroceryList) QueryMembers() *UserQuery {
	return (&GroceryListClient{gl.config}).QueryMembers(gl)
}

// Update returns a builder for updating this GroceryList.
// Note that, you need to call GroceryList.Unwrap() before calling this method, if this GroceryList
// was returned from a transaction, and the transaction was committed or rolled back.
func (gl *GroceryList) Update() *GroceryListUpdateOne {
	return (&GroceryListClient{gl.config}).UpdateOne(gl)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gl *GroceryList) Unwrap() *GroceryList {
	tx, ok := gl.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroceryList is not a transactional entity")
	}
	gl.config.driver = tx.drv
	return gl
}

// String implements the fmt.Stringer.
func (gl *GroceryList) String() string {
	var builder strings.Builder
	builder.WriteString("GroceryList(")
	builder.WriteString(fmt.Sprintf("id=%v", gl.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(gl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", modified_at=")
	builder.WriteString(gl.ModifiedAt.Format(time.ANSIC))
	if v := gl.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(gl.Name)
	builder.WriteByte(')')
	return builder.String()
}

// GroceryLists is a parsable slice of GroceryList.
type GroceryLists []*GroceryList

func (gl GroceryLists) config(cfg config) {
	for _i := range gl {
		gl[_i].config = cfg
	}
}
