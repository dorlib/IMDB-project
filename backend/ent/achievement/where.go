// Code generated by entc, DO NOT EDIT.

package achievement

import (
	"imdbv2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
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
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
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
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImage), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Achievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Achievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Achievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Achievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImage), v))
	})
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImage), v))
	})
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.Achievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Achievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImage), v...))
	})
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.Achievement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Achievement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImage), v...))
	})
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImage), v))
	})
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImage), v))
	})
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImage), v))
	})
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImage), v))
	})
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImage), v))
	})
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImage), v))
	})
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImage), v))
	})
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImage), v))
	})
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImage), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Achievement) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Achievement) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Achievement) predicate.Achievement {
	return predicate.Achievement(func(s *sql.Selector) {
		p(s.Not())
	})
}
