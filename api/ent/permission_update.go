// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xdorro/golang-grpc-base-project/api/ent/permission"
	"github.com/xdorro/golang-grpc-base-project/api/ent/predicate"
	"github.com/xdorro/golang-grpc-base-project/api/ent/role"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetDeleteTime sets the "delete_time" field.
func (pu *PermissionUpdate) SetDeleteTime(t time.Time) *PermissionUpdate {
	pu.mutation.SetDeleteTime(t)
	return pu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDeleteTime(t *time.Time) *PermissionUpdate {
	if t != nil {
		pu.SetDeleteTime(*t)
	}
	return pu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (pu *PermissionUpdate) ClearDeleteTime() *PermissionUpdate {
	pu.mutation.ClearDeleteTime()
	return pu
}

// SetName sets the "name" field.
func (pu *PermissionUpdate) SetName(s string) *PermissionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetSlug sets the "slug" field.
func (pu *PermissionUpdate) SetSlug(s string) *PermissionUpdate {
	pu.mutation.SetSlug(s)
	return pu
}

// SetStatus sets the "status" field.
func (pu *PermissionUpdate) SetStatus(i int32) *PermissionUpdate {
	pu.mutation.ResetStatus()
	pu.mutation.SetStatus(i)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableStatus(i *int32) *PermissionUpdate {
	if i != nil {
		pu.SetStatus(*i)
	}
	return pu
}

// AddStatus adds i to the "status" field.
func (pu *PermissionUpdate) AddStatus(i int32) *PermissionUpdate {
	pu.mutation.AddStatus(i)
	return pu
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (pu *PermissionUpdate) AddRoleIDs(ids ...uint64) *PermissionUpdate {
	pu.mutation.AddRoleIDs(ids...)
	return pu
}

// AddRoles adds the "roles" edges to the Role entity.
func (pu *PermissionUpdate) AddRoles(r ...*Role) *PermissionUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (pu *PermissionUpdate) ClearRoles() *PermissionUpdate {
	pu.mutation.ClearRoles()
	return pu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (pu *PermissionUpdate) RemoveRoleIDs(ids ...uint64) *PermissionUpdate {
	pu.mutation.RemoveRoleIDs(ids...)
	return pu
}

// RemoveRoles removes "roles" edges to Role entities.
func (pu *PermissionUpdate) RemoveRoles(r ...*Role) *PermissionUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PermissionUpdate) defaults() error {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		if permission.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized permission.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := permission.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermissionUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Permission.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Slug(); ok {
		if err := permission.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Permission.slug": %w`, err)}
		}
	}
	return nil
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: permission.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldUpdateTime,
		})
	}
	if value, ok := pu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldDeleteTime,
		})
	}
	if pu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: permission.FieldDeleteTime,
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldName,
		})
	}
	if value, ok := pu.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldSlug,
		})
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: permission.FieldStatus,
		})
	}
	if value, ok := pu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: permission.FieldStatus,
		})
	}
	if pu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !pu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetDeleteTime sets the "delete_time" field.
func (puo *PermissionUpdateOne) SetDeleteTime(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetDeleteTime(t)
	return puo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDeleteTime(t *time.Time) *PermissionUpdateOne {
	if t != nil {
		puo.SetDeleteTime(*t)
	}
	return puo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (puo *PermissionUpdateOne) ClearDeleteTime() *PermissionUpdateOne {
	puo.mutation.ClearDeleteTime()
	return puo
}

// SetName sets the "name" field.
func (puo *PermissionUpdateOne) SetName(s string) *PermissionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetSlug sets the "slug" field.
func (puo *PermissionUpdateOne) SetSlug(s string) *PermissionUpdateOne {
	puo.mutation.SetSlug(s)
	return puo
}

// SetStatus sets the "status" field.
func (puo *PermissionUpdateOne) SetStatus(i int32) *PermissionUpdateOne {
	puo.mutation.ResetStatus()
	puo.mutation.SetStatus(i)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableStatus(i *int32) *PermissionUpdateOne {
	if i != nil {
		puo.SetStatus(*i)
	}
	return puo
}

// AddStatus adds i to the "status" field.
func (puo *PermissionUpdateOne) AddStatus(i int32) *PermissionUpdateOne {
	puo.mutation.AddStatus(i)
	return puo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (puo *PermissionUpdateOne) AddRoleIDs(ids ...uint64) *PermissionUpdateOne {
	puo.mutation.AddRoleIDs(ids...)
	return puo
}

// AddRoles adds the "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) AddRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) ClearRoles() *PermissionUpdateOne {
	puo.mutation.ClearRoles()
	return puo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (puo *PermissionUpdateOne) RemoveRoleIDs(ids ...uint64) *PermissionUpdateOne {
	puo.mutation.RemoveRoleIDs(ids...)
	return puo
}

// RemoveRoles removes "roles" edges to Role entities.
func (puo *PermissionUpdateOne) RemoveRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRoleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	var (
		err  error
		node *Permission
	)
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PermissionUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		if permission.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized permission.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := permission.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermissionUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Permission.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Slug(); ok {
		if err := permission.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Permission.slug": %w`, err)}
		}
	}
	return nil
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: permission.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldUpdateTime,
		})
	}
	if value, ok := puo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldDeleteTime,
		})
	}
	if puo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: permission.FieldDeleteTime,
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldName,
		})
	}
	if value, ok := puo.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldSlug,
		})
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: permission.FieldStatus,
		})
	}
	if value, ok := puo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: permission.FieldStatus,
		})
	}
	if puo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !puo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
