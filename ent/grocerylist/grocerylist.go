// Code generated by entc, DO NOT EDIT.

package grocerylist

import (
	"time"
)

const (
	// Label holds the string label denoting the grocerylist type in the database.
	Label = "grocery_list"
	// FieldID holds the string denoting the id field in the database.
	FieldID         = "id"          // FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt  = "created_at"  // FieldModifiedAt holds the string denoting the modified_at vertex property in the database.
	FieldModifiedAt = "modified_at" // FieldDeletedAt holds the string denoting the deleted_at vertex property in the database.
	FieldDeletedAt  = "deleted_at"  // FieldName holds the string denoting the name vertex property in the database.
	FieldName       = "name"

	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"

	// Table holds the table name of the grocerylist in the database.
	Table = "grocery_lists"
	// MembersTable is the table the holds the members relation/edge. The primary key declared below.
	MembersTable = "user_grocerylist"
	// MembersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	MembersInverseTable = "users"
)

// Columns holds all SQL columns for grocerylist fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldModifiedAt,
	FieldDeletedAt,
	FieldName,
}

var (
	// MembersPrimaryKey and MembersColumn2 are the table columns denoting the
	// primary key for the members relation (M2M).
	MembersPrimaryKey = []string{"user_id", "grocery_list_id"}
)

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultModifiedAt holds the default value on creation for the modified_at field.
	DefaultModifiedAt func() time.Time
	// DefaultName holds the default value on creation for the name field.
	DefaultName string
)
