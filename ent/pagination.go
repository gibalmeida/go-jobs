// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/gibalmeida/go-jobs/ent/department"
	"github.com/gibalmeida/go-jobs/ent/job"
	"github.com/gibalmeida/go-jobs/ent/user"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func getCollectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	field := fc.Field

walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Name == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return getCollectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

// DepartmentEdge is the edge representation of Department.
type DepartmentEdge struct {
	Node   *Department `json:"node"`
	Cursor Cursor      `json:"cursor"`
}

// DepartmentConnection is the connection containing edges to Department.
type DepartmentConnection struct {
	Edges      []*DepartmentEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

// DepartmentPaginateOption enables pagination customization.
type DepartmentPaginateOption func(*departmentPager) error

// WithDepartmentOrder configures pagination ordering.
func WithDepartmentOrder(order *DepartmentOrder) DepartmentPaginateOption {
	if order == nil {
		order = DefaultDepartmentOrder
	}
	o := *order
	return func(pager *departmentPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultDepartmentOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithDepartmentFilter configures pagination filter.
func WithDepartmentFilter(filter func(*DepartmentQuery) (*DepartmentQuery, error)) DepartmentPaginateOption {
	return func(pager *departmentPager) error {
		if filter == nil {
			return errors.New("DepartmentQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type departmentPager struct {
	order  *DepartmentOrder
	filter func(*DepartmentQuery) (*DepartmentQuery, error)
}

func newDepartmentPager(opts []DepartmentPaginateOption) (*departmentPager, error) {
	pager := &departmentPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultDepartmentOrder
	}
	return pager, nil
}

func (p *departmentPager) applyFilter(query *DepartmentQuery) (*DepartmentQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *departmentPager) toCursor(d *Department) Cursor {
	return p.order.Field.toCursor(d)
}

func (p *departmentPager) applyCursors(query *DepartmentQuery, after, before *Cursor) *DepartmentQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultDepartmentOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *departmentPager) applyOrder(query *DepartmentQuery, reverse bool) *DepartmentQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultDepartmentOrder.Field {
		query = query.Order(direction.orderFunc(DefaultDepartmentOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Department.
func (d *DepartmentQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...DepartmentPaginateOption,
) (*DepartmentConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newDepartmentPager(opts)
	if err != nil {
		return nil, err
	}

	if d, err = pager.applyFilter(d); err != nil {
		return nil, err
	}

	conn := &DepartmentConnection{Edges: []*DepartmentEdge{}}
	if !hasCollectedField(ctx, edgesField) ||
		first != nil && *first == 0 ||
		last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := d.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) &&
		hasCollectedField(ctx, totalCountField) {
		count, err := d.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	d = pager.applyCursors(d, after, before)
	d = pager.applyOrder(d, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		d = d.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		d = d.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := d.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Department
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Department {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Department {
			return nodes[i]
		}
	}

	conn.Edges = make([]*DepartmentEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &DepartmentEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

var (
	// DepartmentOrderFieldName orders Department by name.
	DepartmentOrderFieldName = &DepartmentOrderField{
		field: department.FieldName,
		toCursor: func(d *Department) Cursor {
			return Cursor{
				ID:    d.ID,
				Value: d.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f DepartmentOrderField) String() string {
	var str string
	switch f.field {
	case department.FieldName:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f DepartmentOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *DepartmentOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("DepartmentOrderField %T must be a string", v)
	}
	switch str {
	case "NAME":
		*f = *DepartmentOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid DepartmentOrderField", str)
	}
	return nil
}

// DepartmentOrderField defines the ordering field of Department.
type DepartmentOrderField struct {
	field    string
	toCursor func(*Department) Cursor
}

// DepartmentOrder defines the ordering of Department.
type DepartmentOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *DepartmentOrderField `json:"field"`
}

// DefaultDepartmentOrder is the default ordering of Department.
var DefaultDepartmentOrder = &DepartmentOrder{
	Direction: OrderDirectionAsc,
	Field: &DepartmentOrderField{
		field: department.FieldID,
		toCursor: func(d *Department) Cursor {
			return Cursor{ID: d.ID}
		},
	},
}

// ToEdge converts Department into DepartmentEdge.
func (d *Department) ToEdge(order *DepartmentOrder) *DepartmentEdge {
	if order == nil {
		order = DefaultDepartmentOrder
	}
	return &DepartmentEdge{
		Node:   d,
		Cursor: order.Field.toCursor(d),
	}
}

// JobEdge is the edge representation of Job.
type JobEdge struct {
	Node   *Job   `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// JobConnection is the connection containing edges to Job.
type JobConnection struct {
	Edges      []*JobEdge `json:"edges"`
	PageInfo   PageInfo   `json:"pageInfo"`
	TotalCount int        `json:"totalCount"`
}

// JobPaginateOption enables pagination customization.
type JobPaginateOption func(*jobPager) error

// WithJobOrder configures pagination ordering.
func WithJobOrder(order *JobOrder) JobPaginateOption {
	if order == nil {
		order = DefaultJobOrder
	}
	o := *order
	return func(pager *jobPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultJobOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithJobFilter configures pagination filter.
func WithJobFilter(filter func(*JobQuery) (*JobQuery, error)) JobPaginateOption {
	return func(pager *jobPager) error {
		if filter == nil {
			return errors.New("JobQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type jobPager struct {
	order  *JobOrder
	filter func(*JobQuery) (*JobQuery, error)
}

func newJobPager(opts []JobPaginateOption) (*jobPager, error) {
	pager := &jobPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultJobOrder
	}
	return pager, nil
}

func (p *jobPager) applyFilter(query *JobQuery) (*JobQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *jobPager) toCursor(j *Job) Cursor {
	return p.order.Field.toCursor(j)
}

func (p *jobPager) applyCursors(query *JobQuery, after, before *Cursor) *JobQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultJobOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *jobPager) applyOrder(query *JobQuery, reverse bool) *JobQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultJobOrder.Field {
		query = query.Order(direction.orderFunc(DefaultJobOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Job.
func (j *JobQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...JobPaginateOption,
) (*JobConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newJobPager(opts)
	if err != nil {
		return nil, err
	}

	if j, err = pager.applyFilter(j); err != nil {
		return nil, err
	}

	conn := &JobConnection{Edges: []*JobEdge{}}
	if !hasCollectedField(ctx, edgesField) ||
		first != nil && *first == 0 ||
		last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := j.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) &&
		hasCollectedField(ctx, totalCountField) {
		count, err := j.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	j = pager.applyCursors(j, after, before)
	j = pager.applyOrder(j, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		j = j.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		j = j.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := j.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Job
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Job {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Job {
			return nodes[i]
		}
	}

	conn.Edges = make([]*JobEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &JobEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

var (
	// JobOrderFieldName orders Job by name.
	JobOrderFieldName = &JobOrderField{
		field: job.FieldName,
		toCursor: func(j *Job) Cursor {
			return Cursor{
				ID:    j.ID,
				Value: j.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f JobOrderField) String() string {
	var str string
	switch f.field {
	case job.FieldName:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f JobOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *JobOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("JobOrderField %T must be a string", v)
	}
	switch str {
	case "NAME":
		*f = *JobOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid JobOrderField", str)
	}
	return nil
}

// JobOrderField defines the ordering field of Job.
type JobOrderField struct {
	field    string
	toCursor func(*Job) Cursor
}

// JobOrder defines the ordering of Job.
type JobOrder struct {
	Direction OrderDirection `json:"direction"`
	Field     *JobOrderField `json:"field"`
}

// DefaultJobOrder is the default ordering of Job.
var DefaultJobOrder = &JobOrder{
	Direction: OrderDirectionAsc,
	Field: &JobOrderField{
		field: job.FieldID,
		toCursor: func(j *Job) Cursor {
			return Cursor{ID: j.ID}
		},
	},
}

// ToEdge converts Job into JobEdge.
func (j *Job) ToEdge(order *JobOrder) *JobEdge {
	if order == nil {
		order = DefaultJobOrder
	}
	return &JobEdge{
		Node:   j,
		Cursor: order.Field.toCursor(j),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}

	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}

	conn := &UserConnection{Edges: []*UserEdge{}}
	if !hasCollectedField(ctx, edgesField) ||
		first != nil && *first == 0 ||
		last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := u.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) &&
		hasCollectedField(ctx, totalCountField) {
		count, err := u.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		u = u.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := u.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}

	conn.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

var (
	// UserOrderFieldName orders User by name.
	UserOrderFieldName = &UserOrderField{
		field: user.FieldName,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Name,
			}
		},
	}
	// UserOrderFieldEmail orders User by email.
	UserOrderFieldEmail = &UserOrderField{
		field: user.FieldEmail,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Email,
			}
		},
	}
	// UserOrderFieldRole orders User by role.
	UserOrderFieldRole = &UserOrderField{
		field: user.FieldRole,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Role,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserOrderField) String() string {
	var str string
	switch f.field {
	case user.FieldName:
		str = "NAME"
	case user.FieldEmail:
		str = "EMAIL"
	case user.FieldRole:
		str = "ROLE"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserOrderField %T must be a string", v)
	}
	switch str {
	case "NAME":
		*f = *UserOrderFieldName
	case "EMAIL":
		*f = *UserOrderFieldEmail
	case "ROLE":
		*f = *UserOrderFieldRole
	default:
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
