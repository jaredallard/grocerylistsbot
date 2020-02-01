// Code generated by entc, DO NOT EDIT.

package useridmapping

import (
	"fmt"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/jaredallard/grocerylistsbot/ent/schema"
)

const (
	// Label holds the string label denoting the useridmapping type in the database.
	Label = "user_id_mapping"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldModifiedAt holds the string denoting the modified_at vertex property in the database.
	FieldModifiedAt = "modified_at"
	// FieldDeletedAt holds the string denoting the deleted_at vertex property in the database.
	FieldDeletedAt = "deleted_at"
	// FieldPlatformType holds the string denoting the platform_type vertex property in the database.
	FieldPlatformType = "platform_type"
	// FieldPlatformID holds the string denoting the platform_id vertex property in the database.
	FieldPlatformID = "platform_id"

	// Table holds the table name of the useridmapping in the database.
	Table = "user_id_mappings"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "user_id_mappings"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id_mapping_user_id"
)

// Columns holds all SQL columns for useridmapping fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldModifiedAt,
	FieldDeletedAt,
	FieldPlatformType,
	FieldPlatformID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the UserIDMapping type.
var ForeignKeys = []string{
	"user_id_mapping_user_id",
}

var (
	mixin       = schema.UserIDMapping{}.Mixin()
	mixinFields = [...][]ent.Field{
		mixin[0].Fields(),
	}
	fields = schema.UserIDMapping{}.Fields()

	// descCreatedAt is the schema descriptor for created_at field.
	descCreatedAt = mixinFields[0][0].Descriptor()
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt = descCreatedAt.Default.(func() time.Time)

	// descModifiedAt is the schema descriptor for modified_at field.
	descModifiedAt = mixinFields[0][1].Descriptor()
	// DefaultModifiedAt holds the default value on creation for the modified_at field.
	DefaultModifiedAt = descModifiedAt.Default.(func() time.Time)
)

// PlatformType defines the type for the platform_type enum field.
type PlatformType string

// PlatformType values.
const (
	PlatformTypeTelegram PlatformType = "telegram"
	PlatformTypeSms      PlatformType = "sms"
)

func (s PlatformType) String() string {
	return string(s)
}

// PlatformTypeValidator is a validator for the "pt" field enum values. It is called by the builders before save.
func PlatformTypeValidator(pt PlatformType) error {
	switch pt {
	case PlatformTypeTelegram, PlatformTypeSms:
		return nil
	default:
		return fmt.Errorf("useridmapping: invalid enum value for platform_type field: %q", pt)
	}
}