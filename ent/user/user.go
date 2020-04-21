// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID         = "id"          // FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt  = "created_at"  // FieldModifiedAt holds the string denoting the modified_at vertex property in the database.
	FieldModifiedAt = "modified_at" // FieldDeletedAt holds the string denoting the deleted_at vertex property in the database.
	FieldDeletedAt  = "deleted_at"  // FieldName holds the string denoting the name vertex property in the database.
	FieldName       = "name"

	// EdgeGrocerylist holds the string denoting the grocerylist edge name in mutations.
	EdgeGrocerylist = "grocerylist"
	// EdgeActiveList holds the string denoting the active_list edge name in mutations.
	EdgeActiveList = "active_list"

	// Table holds the table name of the user in the database.
	Table = "users"
	// GrocerylistTable is the table the holds the grocerylist relation/edge. The primary key declared below.
	GrocerylistTable = "user_grocerylist"
	// GrocerylistInverseTable is the table name for the GroceryList entity.
	// It exists in this package in order to avoid circular dependency with the "grocerylist" package.
	GrocerylistInverseTable = "grocery_lists"
	// ActiveListTable is the table the holds the active_list relation/edge.
	ActiveListTable = "users"
	// ActiveListInverseTable is the table name for the GroceryList entity.
	// It exists in this package in order to avoid circular dependency with the "grocerylist" package.
	ActiveListInverseTable = "grocery_lists"
	// ActiveListColumn is the table column denoting the active_list relation/edge.
	ActiveListColumn = "user_active_list"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldModifiedAt,
	FieldDeletedAt,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"user_active_list",
}

var (
	// GrocerylistPrimaryKey and GrocerylistColumn2 are the table columns denoting the
	// primary key for the grocerylist relation (M2M).
	GrocerylistPrimaryKey = []string{"user_id", "grocery_list_id"}
)

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultModifiedAt holds the default value on creation for the modified_at field.
	DefaultModifiedAt func() time.Time
)
