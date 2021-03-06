// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/jaredallard/grocerylistsbot/ent/grocerylist"
	"github.com/jaredallard/grocerylistsbot/ent/user"
)

// User is the model entity for the User schema.
type User struct {
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
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges            UserEdges `json:"edges"`
	user_active_list *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Grocerylist holds the value of the grocerylist edge.
	Grocerylist []*GroceryList
	// ActiveList holds the value of the active_list edge.
	ActiveList *GroceryList
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GrocerylistOrErr returns the Grocerylist value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GrocerylistOrErr() ([]*GroceryList, error) {
	if e.loadedTypes[0] {
		return e.Grocerylist, nil
	}
	return nil, &NotLoadedError{edge: "grocerylist"}
}

// ActiveListOrErr returns the ActiveList value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ActiveListOrErr() (*GroceryList, error) {
	if e.loadedTypes[1] {
		if e.ActiveList == nil {
			// The edge active_list was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: grocerylist.Label}
		}
		return e.ActiveList, nil
	}
	return nil, &NotLoadedError{edge: "active_list"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // modified_at
		&sql.NullTime{},   // deleted_at
		&sql.NullString{}, // name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_active_list
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		u.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field modified_at", values[1])
	} else if value.Valid {
		u.ModifiedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field deleted_at", values[2])
	} else if value.Valid {
		u.DeletedAt = new(time.Time)
		*u.DeletedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[3])
	} else if value.Valid {
		u.Name = value.String
	}
	values = values[4:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_active_list", value)
		} else if value.Valid {
			u.user_active_list = new(int)
			*u.user_active_list = int(value.Int64)
		}
	}
	return nil
}

// QueryGrocerylist queries the grocerylist edge of the User.
func (u *User) QueryGrocerylist() *GroceryListQuery {
	return (&UserClient{config: u.config}).QueryGrocerylist(u)
}

// QueryActiveList queries the active_list edge of the User.
func (u *User) QueryActiveList() *GroceryListQuery {
	return (&UserClient{config: u.config}).QueryActiveList(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", modified_at=")
	builder.WriteString(u.ModifiedAt.Format(time.ANSIC))
	if v := u.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
