// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/jaredallard/grocerylistsbot/ent/groceryitem"
	"github.com/jaredallard/grocerylistsbot/ent/grocerylist"
)

// GroceryItemCreate is the builder for creating a GroceryItem entity.
type GroceryItemCreate struct {
	config
	created_at  *time.Time
	modified_at *time.Time
	deleted_at  *time.Time
	name        *string
	status      *groceryitem.Status
	price       *float64
	grocerylist map[int]struct{}
}

// SetCreatedAt sets the created_at field.
func (gic *GroceryItemCreate) SetCreatedAt(t time.Time) *GroceryItemCreate {
	gic.created_at = &t
	return gic
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (gic *GroceryItemCreate) SetNillableCreatedAt(t *time.Time) *GroceryItemCreate {
	if t != nil {
		gic.SetCreatedAt(*t)
	}
	return gic
}

// SetModifiedAt sets the modified_at field.
func (gic *GroceryItemCreate) SetModifiedAt(t time.Time) *GroceryItemCreate {
	gic.modified_at = &t
	return gic
}

// SetNillableModifiedAt sets the modified_at field if the given value is not nil.
func (gic *GroceryItemCreate) SetNillableModifiedAt(t *time.Time) *GroceryItemCreate {
	if t != nil {
		gic.SetModifiedAt(*t)
	}
	return gic
}

// SetDeletedAt sets the deleted_at field.
func (gic *GroceryItemCreate) SetDeletedAt(t time.Time) *GroceryItemCreate {
	gic.deleted_at = &t
	return gic
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (gic *GroceryItemCreate) SetNillableDeletedAt(t *time.Time) *GroceryItemCreate {
	if t != nil {
		gic.SetDeletedAt(*t)
	}
	return gic
}

// SetName sets the name field.
func (gic *GroceryItemCreate) SetName(s string) *GroceryItemCreate {
	gic.name = &s
	return gic
}

// SetStatus sets the status field.
func (gic *GroceryItemCreate) SetStatus(gr groceryitem.Status) *GroceryItemCreate {
	gic.status = &gr
	return gic
}

// SetNillableStatus sets the status field if the given value is not nil.
func (gic *GroceryItemCreate) SetNillableStatus(gr *groceryitem.Status) *GroceryItemCreate {
	if gr != nil {
		gic.SetStatus(*gr)
	}
	return gic
}

// SetPrice sets the price field.
func (gic *GroceryItemCreate) SetPrice(f float64) *GroceryItemCreate {
	gic.price = &f
	return gic
}

// SetNillablePrice sets the price field if the given value is not nil.
func (gic *GroceryItemCreate) SetNillablePrice(f *float64) *GroceryItemCreate {
	if f != nil {
		gic.SetPrice(*f)
	}
	return gic
}

// SetGrocerylistID sets the grocerylist edge to GroceryList by id.
func (gic *GroceryItemCreate) SetGrocerylistID(id int) *GroceryItemCreate {
	if gic.grocerylist == nil {
		gic.grocerylist = make(map[int]struct{})
	}
	gic.grocerylist[id] = struct{}{}
	return gic
}

// SetNillableGrocerylistID sets the grocerylist edge to GroceryList by id if the given value is not nil.
func (gic *GroceryItemCreate) SetNillableGrocerylistID(id *int) *GroceryItemCreate {
	if id != nil {
		gic = gic.SetGrocerylistID(*id)
	}
	return gic
}

// SetGrocerylist sets the grocerylist edge to GroceryList.
func (gic *GroceryItemCreate) SetGrocerylist(g *GroceryList) *GroceryItemCreate {
	return gic.SetGrocerylistID(g.ID)
}

// Save creates the GroceryItem in the database.
func (gic *GroceryItemCreate) Save(ctx context.Context) (*GroceryItem, error) {
	if gic.created_at == nil {
		v := groceryitem.DefaultCreatedAt()
		gic.created_at = &v
	}
	if gic.modified_at == nil {
		v := groceryitem.DefaultModifiedAt()
		gic.modified_at = &v
	}
	if gic.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if gic.status == nil {
		v := groceryitem.DefaultStatus
		gic.status = &v
	}
	if err := groceryitem.StatusValidator(*gic.status); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"status\": %v", err)
	}
	if len(gic.grocerylist) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"grocerylist\"")
	}
	return gic.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (gic *GroceryItemCreate) SaveX(ctx context.Context) *GroceryItem {
	v, err := gic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gic *GroceryItemCreate) sqlSave(ctx context.Context) (*GroceryItem, error) {
	var (
		gi    = &GroceryItem{config: gic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: groceryitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groceryitem.FieldID,
			},
		}
	)
	if value := gic.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: groceryitem.FieldCreatedAt,
		})
		gi.CreatedAt = *value
	}
	if value := gic.modified_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: groceryitem.FieldModifiedAt,
		})
		gi.ModifiedAt = *value
	}
	if value := gic.deleted_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: groceryitem.FieldDeletedAt,
		})
		gi.DeletedAt = value
	}
	if value := gic.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: groceryitem.FieldName,
		})
		gi.Name = *value
	}
	if value := gic.status; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: groceryitem.FieldStatus,
		})
		gi.Status = *value
	}
	if value := gic.price; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  *value,
			Column: groceryitem.FieldPrice,
		})
		gi.Price = *value
	}
	if nodes := gic.grocerylist; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groceryitem.GrocerylistTable,
			Columns: []string{groceryitem.GrocerylistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: grocerylist.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, gic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	gi.ID = int(id)
	return gi, nil
}