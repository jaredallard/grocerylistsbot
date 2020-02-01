// Code generated by entc, DO NOT EDIT.

package grocerylist

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/jaredallard/grocerylistsbot/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.GroceryList {
	return predicate.GroceryList(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	},
	)
}

// ModifiedAt applies equality check predicate on the "modified_at" field. It's identical to ModifiedAtEQ.
func ModifiedAt(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModifiedAt), v))
	},
	)
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	},
	)
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	},
	)
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	},
	)
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GroceryList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroceryList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	},
	)
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GroceryList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroceryList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	},
	)
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	},
	)
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	},
	)
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	},
	)
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	},
	)
}

// ModifiedAtEQ applies the EQ predicate on the "modified_at" field.
func ModifiedAtEQ(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModifiedAt), v))
	},
	)
}

// ModifiedAtNEQ applies the NEQ predicate on the "modified_at" field.
func ModifiedAtNEQ(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModifiedAt), v))
	},
	)
}

// ModifiedAtIn applies the In predicate on the "modified_at" field.
func ModifiedAtIn(vs ...time.Time) predicate.GroceryList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroceryList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldModifiedAt), v...))
	},
	)
}

// ModifiedAtNotIn applies the NotIn predicate on the "modified_at" field.
func ModifiedAtNotIn(vs ...time.Time) predicate.GroceryList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroceryList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldModifiedAt), v...))
	},
	)
}

// ModifiedAtGT applies the GT predicate on the "modified_at" field.
func ModifiedAtGT(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModifiedAt), v))
	},
	)
}

// ModifiedAtGTE applies the GTE predicate on the "modified_at" field.
func ModifiedAtGTE(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModifiedAt), v))
	},
	)
}

// ModifiedAtLT applies the LT predicate on the "modified_at" field.
func ModifiedAtLT(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModifiedAt), v))
	},
	)
}

// ModifiedAtLTE applies the LTE predicate on the "modified_at" field.
func ModifiedAtLTE(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModifiedAt), v))
	},
	)
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	},
	)
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	},
	)
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.GroceryList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroceryList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	},
	)
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.GroceryList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroceryList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	},
	)
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	},
	)
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	},
	)
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	},
	)
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	},
	)
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	},
	)
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.GroceryList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroceryList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	},
	)
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.GroceryList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroceryList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	},
	)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	},
	)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	},
	)
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	},
	)
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	},
	)
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MembersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.User) predicate.GroceryList {
	return predicate.GroceryList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MembersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.GroceryList) predicate.GroceryList {
	return predicate.GroceryList(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.GroceryList) predicate.GroceryList {
	return predicate.GroceryList(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GroceryList) predicate.GroceryList {
	return predicate.GroceryList(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
