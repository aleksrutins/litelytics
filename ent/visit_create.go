// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/aleksrutins/litelytics/ent/site"
	"github.com/aleksrutins/litelytics/ent/visit"
)

// VisitCreate is the builder for creating a Visit entity.
type VisitCreate struct {
	config
	mutation *VisitMutation
	hooks    []Hook
}

// SetPath sets the "path" field.
func (vc *VisitCreate) SetPath(s string) *VisitCreate {
	vc.mutation.SetPath(s)
	return vc
}

// SetReferrer sets the "referrer" field.
func (vc *VisitCreate) SetReferrer(s string) *VisitCreate {
	vc.mutation.SetReferrer(s)
	return vc
}

// SetTimestamp sets the "timestamp" field.
func (vc *VisitCreate) SetTimestamp(t time.Time) *VisitCreate {
	vc.mutation.SetTimestamp(t)
	return vc
}

// SetIP sets the "ip" field.
func (vc *VisitCreate) SetIP(s string) *VisitCreate {
	vc.mutation.SetIP(s)
	return vc
}

// SetSiteID sets the "site" edge to the Site entity by ID.
func (vc *VisitCreate) SetSiteID(id int) *VisitCreate {
	vc.mutation.SetSiteID(id)
	return vc
}

// SetSite sets the "site" edge to the Site entity.
func (vc *VisitCreate) SetSite(s *Site) *VisitCreate {
	return vc.SetSiteID(s.ID)
}

// Mutation returns the VisitMutation object of the builder.
func (vc *VisitCreate) Mutation() *VisitMutation {
	return vc.mutation
}

// Save creates the Visit in the database.
func (vc *VisitCreate) Save(ctx context.Context) (*Visit, error) {
	var (
		err  error
		node *Visit
	)
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VisitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, vc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Visit)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from VisitMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VisitCreate) SaveX(ctx context.Context) *Visit {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VisitCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VisitCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VisitCreate) check() error {
	if _, ok := vc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Visit.path"`)}
	}
	if _, ok := vc.mutation.Referrer(); !ok {
		return &ValidationError{Name: "referrer", err: errors.New(`ent: missing required field "Visit.referrer"`)}
	}
	if _, ok := vc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "Visit.timestamp"`)}
	}
	if _, ok := vc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "Visit.ip"`)}
	}
	if _, ok := vc.mutation.SiteID(); !ok {
		return &ValidationError{Name: "site", err: errors.New(`ent: missing required edge "Visit.site"`)}
	}
	return nil
}

func (vc *VisitCreate) sqlSave(ctx context.Context) (*Visit, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vc *VisitCreate) createSpec() (*Visit, *sqlgraph.CreateSpec) {
	var (
		_node = &Visit{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: visit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: visit.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.Path(); ok {
		_spec.SetField(visit.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := vc.mutation.Referrer(); ok {
		_spec.SetField(visit.FieldReferrer, field.TypeString, value)
		_node.Referrer = value
	}
	if value, ok := vc.mutation.Timestamp(); ok {
		_spec.SetField(visit.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	if value, ok := vc.mutation.IP(); ok {
		_spec.SetField(visit.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if nodes := vc.mutation.SiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   visit.SiteTable,
			Columns: []string{visit.SiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.site_visits = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VisitCreateBulk is the builder for creating many Visit entities in bulk.
type VisitCreateBulk struct {
	config
	builders []*VisitCreate
}

// Save creates the Visit entities in the database.
func (vcb *VisitCreateBulk) Save(ctx context.Context) ([]*Visit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Visit, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VisitMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VisitCreateBulk) SaveX(ctx context.Context) []*Visit {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VisitCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VisitCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
