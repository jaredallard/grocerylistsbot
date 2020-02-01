// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/jaredallard/grocerylistsbot/ent/groceryitem"
	"github.com/jaredallard/grocerylistsbot/ent/grocerylist"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// GroceryItemsColumns holds the columns for the "grocery_items" table.
	GroceryItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"purchased", "unpurchased"}, Default: groceryitem.DefaultStatus},
		{Name: "price", Type: field.TypeFloat64, Nullable: true},
		{Name: "grocery_item_grocerylist_id", Type: field.TypeInt, Nullable: true},
	}
	// GroceryItemsTable holds the schema information for the "grocery_items" table.
	GroceryItemsTable = &schema.Table{
		Name:       "grocery_items",
		Columns:    GroceryItemsColumns,
		PrimaryKey: []*schema.Column{GroceryItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "grocery_items_grocery_lists_grocerylist",
				Columns: []*schema.Column{GroceryItemsColumns[7]},

				RefColumns: []*schema.Column{GroceryListsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroceryListsColumns holds the columns for the "grocery_lists" table.
	GroceryListsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Default: grocerylist.DefaultName},
	}
	// GroceryListsTable holds the schema information for the "grocery_lists" table.
	GroceryListsTable = &schema.Table{
		Name:        "grocery_lists",
		Columns:     GroceryListsColumns,
		PrimaryKey:  []*schema.Column{GroceryListsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_active_list_id", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_grocery_lists_active_list",
				Columns: []*schema.Column{UsersColumns[5]},

				RefColumns: []*schema.Column{GroceryListsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_name",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[4]},
			},
		},
	}
	// UserIDMappingsColumns holds the columns for the "user_id_mappings" table.
	UserIDMappingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "platform_type", Type: field.TypeEnum, Enums: []string{"telegram", "sms"}},
		{Name: "platform_id", Type: field.TypeString},
		{Name: "user_id_mapping_user_id", Type: field.TypeInt, Nullable: true},
	}
	// UserIDMappingsTable holds the schema information for the "user_id_mappings" table.
	UserIDMappingsTable = &schema.Table{
		Name:       "user_id_mappings",
		Columns:    UserIDMappingsColumns,
		PrimaryKey: []*schema.Column{UserIDMappingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_id_mappings_users_user",
				Columns: []*schema.Column{UserIDMappingsColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "useridmapping_platform_type_platform_id",
				Unique:  true,
				Columns: []*schema.Column{UserIDMappingsColumns[4], UserIDMappingsColumns[5]},
			},
		},
	}
	// UserGrocerylistColumns holds the columns for the "user_grocerylist" table.
	UserGrocerylistColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "grocery_list_id", Type: field.TypeInt},
	}
	// UserGrocerylistTable holds the schema information for the "user_grocerylist" table.
	UserGrocerylistTable = &schema.Table{
		Name:       "user_grocerylist",
		Columns:    UserGrocerylistColumns,
		PrimaryKey: []*schema.Column{UserGrocerylistColumns[0], UserGrocerylistColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_grocerylist_user_id",
				Columns: []*schema.Column{UserGrocerylistColumns[0]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "user_grocerylist_grocery_list_id",
				Columns: []*schema.Column{UserGrocerylistColumns[1]},

				RefColumns: []*schema.Column{GroceryListsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroceryItemsTable,
		GroceryListsTable,
		UsersTable,
		UserIDMappingsTable,
		UserGrocerylistTable,
	}
)

func init() {
	GroceryItemsTable.ForeignKeys[0].RefTable = GroceryListsTable
	UsersTable.ForeignKeys[0].RefTable = GroceryListsTable
	UserIDMappingsTable.ForeignKeys[0].RefTable = UsersTable
	UserGrocerylistTable.ForeignKeys[0].RefTable = UsersTable
	UserGrocerylistTable.ForeignKeys[1].RefTable = GroceryListsTable
}