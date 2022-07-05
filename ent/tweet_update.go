// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"goent/ent/predicate"
	"goent/ent/tweet"
	"goent/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TweetUpdate is the builder for updating Tweet entities.
type TweetUpdate struct {
	config
	hooks    []Hook
	mutation *TweetMutation
}

// Where appends a list predicates to the TweetUpdate builder.
func (tu *TweetUpdate) Where(ps ...predicate.Tweet) *TweetUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TweetUpdate) SetTitle(s string) *TweetUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetContent sets the "content" field.
func (tu *TweetUpdate) SetContent(s string) *TweetUpdate {
	tu.mutation.SetContent(s)
	return tu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableContent(s *string) *TweetUpdate {
	if s != nil {
		tu.SetContent(*s)
	}
	return tu
}

// ClearContent clears the value of the "content" field.
func (tu *TweetUpdate) ClearContent() *TweetUpdate {
	tu.mutation.ClearContent()
	return tu
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (tu *TweetUpdate) SetUsersID(id int) *TweetUpdate {
	tu.mutation.SetUsersID(id)
	return tu
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (tu *TweetUpdate) SetNillableUsersID(id *int) *TweetUpdate {
	if id != nil {
		tu = tu.SetUsersID(*id)
	}
	return tu
}

// SetUsers sets the "users" edge to the User entity.
func (tu *TweetUpdate) SetUsers(u *User) *TweetUpdate {
	return tu.SetUsersID(u.ID)
}

// Mutation returns the TweetMutation object of the builder.
func (tu *TweetUpdate) Mutation() *TweetMutation {
	return tu.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (tu *TweetUpdate) ClearUsers() *TweetUpdate {
	tu.mutation.ClearUsers()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TweetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TweetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TweetUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TweetUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TweetUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TweetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweet.Table,
			Columns: tweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tweet.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldTitle,
		})
	}
	if value, ok := tu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldContent,
		})
	}
	if tu.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tweet.FieldContent,
		})
	}
	if tu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.UsersTable,
			Columns: []string{tweet.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.UsersTable,
			Columns: []string{tweet.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TweetUpdateOne is the builder for updating a single Tweet entity.
type TweetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TweetMutation
}

// SetTitle sets the "title" field.
func (tuo *TweetUpdateOne) SetTitle(s string) *TweetUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetContent sets the "content" field.
func (tuo *TweetUpdateOne) SetContent(s string) *TweetUpdateOne {
	tuo.mutation.SetContent(s)
	return tuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableContent(s *string) *TweetUpdateOne {
	if s != nil {
		tuo.SetContent(*s)
	}
	return tuo
}

// ClearContent clears the value of the "content" field.
func (tuo *TweetUpdateOne) ClearContent() *TweetUpdateOne {
	tuo.mutation.ClearContent()
	return tuo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (tuo *TweetUpdateOne) SetUsersID(id int) *TweetUpdateOne {
	tuo.mutation.SetUsersID(id)
	return tuo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableUsersID(id *int) *TweetUpdateOne {
	if id != nil {
		tuo = tuo.SetUsersID(*id)
	}
	return tuo
}

// SetUsers sets the "users" edge to the User entity.
func (tuo *TweetUpdateOne) SetUsers(u *User) *TweetUpdateOne {
	return tuo.SetUsersID(u.ID)
}

// Mutation returns the TweetMutation object of the builder.
func (tuo *TweetUpdateOne) Mutation() *TweetMutation {
	return tuo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (tuo *TweetUpdateOne) ClearUsers() *TweetUpdateOne {
	tuo.mutation.ClearUsers()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TweetUpdateOne) Select(field string, fields ...string) *TweetUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tweet entity.
func (tuo *TweetUpdateOne) Save(ctx context.Context) (*Tweet, error) {
	var (
		err  error
		node *Tweet
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TweetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TweetUpdateOne) SaveX(ctx context.Context) *Tweet {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TweetUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TweetUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TweetUpdateOne) sqlSave(ctx context.Context) (_node *Tweet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweet.Table,
			Columns: tweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tweet.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tweet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.FieldID)
		for _, f := range fields {
			if !tweet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldTitle,
		})
	}
	if value, ok := tuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldContent,
		})
	}
	if tuo.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tweet.FieldContent,
		})
	}
	if tuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.UsersTable,
			Columns: []string{tweet.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.UsersTable,
			Columns: []string{tweet.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tweet{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}