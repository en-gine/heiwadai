// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Checkin is an object representing the database table.
type Checkin struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	StoreID   string    `boil:"store_id" json:"store_id" toml:"store_id" yaml:"store_id"`
	UserID    string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CheckInAt time.Time `boil:"check_in_at" json:"check_in_at" toml:"check_in_at" yaml:"check_in_at"`
	Archive   bool      `boil:"archive" json:"archive" toml:"archive" yaml:"archive"`
	CreateAt  time.Time `boil:"create_at" json:"create_at" toml:"create_at" yaml:"create_at"`
	UpdateAt  time.Time `boil:"update_at" json:"update_at" toml:"update_at" yaml:"update_at"`

	R *checkinR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L checkinL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CheckinColumns = struct {
	ID        string
	StoreID   string
	UserID    string
	CheckInAt string
	Archive   string
	CreateAt  string
	UpdateAt  string
}{
	ID:        "id",
	StoreID:   "store_id",
	UserID:    "user_id",
	CheckInAt: "check_in_at",
	Archive:   "archive",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
}

var CheckinTableColumns = struct {
	ID        string
	StoreID   string
	UserID    string
	CheckInAt string
	Archive   string
	CreateAt  string
	UpdateAt  string
}{
	ID:        "checkin.id",
	StoreID:   "checkin.store_id",
	UserID:    "checkin.user_id",
	CheckInAt: "checkin.check_in_at",
	Archive:   "checkin.archive",
	CreateAt:  "checkin.create_at",
	UpdateAt:  "checkin.update_at",
}

// Generated where

var CheckinWhere = struct {
	ID        whereHelperstring
	StoreID   whereHelperstring
	UserID    whereHelperstring
	CheckInAt whereHelpertime_Time
	Archive   whereHelperbool
	CreateAt  whereHelpertime_Time
	UpdateAt  whereHelpertime_Time
}{
	ID:        whereHelperstring{field: "\"checkin\".\"id\""},
	StoreID:   whereHelperstring{field: "\"checkin\".\"store_id\""},
	UserID:    whereHelperstring{field: "\"checkin\".\"user_id\""},
	CheckInAt: whereHelpertime_Time{field: "\"checkin\".\"check_in_at\""},
	Archive:   whereHelperbool{field: "\"checkin\".\"archive\""},
	CreateAt:  whereHelpertime_Time{field: "\"checkin\".\"create_at\""},
	UpdateAt:  whereHelpertime_Time{field: "\"checkin\".\"update_at\""},
}

// CheckinRels is where relationship names are stored.
var CheckinRels = struct {
	Store string
	User  string
}{
	Store: "Store",
	User:  "User",
}

// checkinR is where relationships are stored.
type checkinR struct {
	Store *Store     `boil:"Store" json:"Store" toml:"Store" yaml:"Store"`
	User  *UserDatum `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*checkinR) NewStruct() *checkinR {
	return &checkinR{}
}

func (r *checkinR) GetStore() *Store {
	if r == nil {
		return nil
	}
	return r.Store
}

func (r *checkinR) GetUser() *UserDatum {
	if r == nil {
		return nil
	}
	return r.User
}

// checkinL is where Load methods for each relationship are stored.
type checkinL struct{}

var (
	checkinAllColumns            = []string{"id", "store_id", "user_id", "check_in_at", "archive", "create_at", "update_at"}
	checkinColumnsWithoutDefault = []string{"id", "store_id", "user_id", "check_in_at", "archive"}
	checkinColumnsWithDefault    = []string{"create_at", "update_at"}
	checkinPrimaryKeyColumns     = []string{"id"}
	checkinGeneratedColumns      = []string{}
)

type (
	// CheckinSlice is an alias for a slice of pointers to Checkin.
	// This should almost always be used instead of []Checkin.
	CheckinSlice []*Checkin
	// CheckinHook is the signature for custom Checkin hook methods
	CheckinHook func(context.Context, boil.ContextExecutor, *Checkin) error

	checkinQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	checkinType                 = reflect.TypeOf(&Checkin{})
	checkinMapping              = queries.MakeStructMapping(checkinType)
	checkinPrimaryKeyMapping, _ = queries.BindMapping(checkinType, checkinMapping, checkinPrimaryKeyColumns)
	checkinInsertCacheMut       sync.RWMutex
	checkinInsertCache          = make(map[string]insertCache)
	checkinUpdateCacheMut       sync.RWMutex
	checkinUpdateCache          = make(map[string]updateCache)
	checkinUpsertCacheMut       sync.RWMutex
	checkinUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var checkinAfterSelectHooks []CheckinHook

var checkinBeforeInsertHooks []CheckinHook
var checkinAfterInsertHooks []CheckinHook

var checkinBeforeUpdateHooks []CheckinHook
var checkinAfterUpdateHooks []CheckinHook

var checkinBeforeDeleteHooks []CheckinHook
var checkinAfterDeleteHooks []CheckinHook

var checkinBeforeUpsertHooks []CheckinHook
var checkinAfterUpsertHooks []CheckinHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Checkin) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checkinAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Checkin) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checkinBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Checkin) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checkinAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Checkin) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checkinBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Checkin) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checkinAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Checkin) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checkinBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Checkin) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checkinAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Checkin) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checkinBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Checkin) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checkinAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCheckinHook registers your hook function for all future operations.
func AddCheckinHook(hookPoint boil.HookPoint, checkinHook CheckinHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		checkinAfterSelectHooks = append(checkinAfterSelectHooks, checkinHook)
	case boil.BeforeInsertHook:
		checkinBeforeInsertHooks = append(checkinBeforeInsertHooks, checkinHook)
	case boil.AfterInsertHook:
		checkinAfterInsertHooks = append(checkinAfterInsertHooks, checkinHook)
	case boil.BeforeUpdateHook:
		checkinBeforeUpdateHooks = append(checkinBeforeUpdateHooks, checkinHook)
	case boil.AfterUpdateHook:
		checkinAfterUpdateHooks = append(checkinAfterUpdateHooks, checkinHook)
	case boil.BeforeDeleteHook:
		checkinBeforeDeleteHooks = append(checkinBeforeDeleteHooks, checkinHook)
	case boil.AfterDeleteHook:
		checkinAfterDeleteHooks = append(checkinAfterDeleteHooks, checkinHook)
	case boil.BeforeUpsertHook:
		checkinBeforeUpsertHooks = append(checkinBeforeUpsertHooks, checkinHook)
	case boil.AfterUpsertHook:
		checkinAfterUpsertHooks = append(checkinAfterUpsertHooks, checkinHook)
	}
}

// One returns a single checkin record from the query.
func (q checkinQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Checkin, error) {
	o := &Checkin{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for checkin")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Checkin records from the query.
func (q checkinQuery) All(ctx context.Context, exec boil.ContextExecutor) (CheckinSlice, error) {
	var o []*Checkin

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Checkin slice")
	}

	if len(checkinAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Checkin records in the query.
func (q checkinQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count checkin rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q checkinQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if checkin exists")
	}

	return count > 0, nil
}

// Store pointed to by the foreign key.
func (o *Checkin) Store(mods ...qm.QueryMod) storeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.StoreID),
	}

	queryMods = append(queryMods, mods...)

	return Stores(queryMods...)
}

// User pointed to by the foreign key.
func (o *Checkin) User(mods ...qm.QueryMod) userDatumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"user_id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return UserData(queryMods...)
}

// LoadStore allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (checkinL) LoadStore(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCheckin interface{}, mods queries.Applicator) error {
	var slice []*Checkin
	var object *Checkin

	if singular {
		var ok bool
		object, ok = maybeCheckin.(*Checkin)
		if !ok {
			object = new(Checkin)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCheckin)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCheckin))
			}
		}
	} else {
		s, ok := maybeCheckin.(*[]*Checkin)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCheckin)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCheckin))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &checkinR{}
		}
		args = append(args, object.StoreID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &checkinR{}
			}

			for _, a := range args {
				if a == obj.StoreID {
					continue Outer
				}
			}

			args = append(args, obj.StoreID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`store`),
		qm.WhereIn(`store.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Store")
	}

	var resultSlice []*Store
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Store")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for store")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for store")
	}

	if len(storeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Store = foreign
		if foreign.R == nil {
			foreign.R = &storeR{}
		}
		foreign.R.Checkins = append(foreign.R.Checkins, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StoreID == foreign.ID {
				local.R.Store = foreign
				if foreign.R == nil {
					foreign.R = &storeR{}
				}
				foreign.R.Checkins = append(foreign.R.Checkins, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (checkinL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCheckin interface{}, mods queries.Applicator) error {
	var slice []*Checkin
	var object *Checkin

	if singular {
		var ok bool
		object, ok = maybeCheckin.(*Checkin)
		if !ok {
			object = new(Checkin)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCheckin)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCheckin))
			}
		}
	} else {
		s, ok := maybeCheckin.(*[]*Checkin)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCheckin)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCheckin))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &checkinR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &checkinR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_data`),
		qm.WhereIn(`user_data.user_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserDatum")
	}

	var resultSlice []*UserDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserDatum")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_data")
	}

	if len(userDatumAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userDatumR{}
		}
		foreign.R.UserCheckins = append(foreign.R.UserCheckins, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userDatumR{}
				}
				foreign.R.UserCheckins = append(foreign.R.UserCheckins, local)
				break
			}
		}
	}

	return nil
}

// SetStore of the checkin to the related item.
// Sets o.R.Store to related.
// Adds o to related.R.Checkins.
func (o *Checkin) SetStore(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Store) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"checkin\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"store_id"}),
		strmangle.WhereClause("\"", "\"", 2, checkinPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StoreID = related.ID
	if o.R == nil {
		o.R = &checkinR{
			Store: related,
		}
	} else {
		o.R.Store = related
	}

	if related.R == nil {
		related.R = &storeR{
			Checkins: CheckinSlice{o},
		}
	} else {
		related.R.Checkins = append(related.R.Checkins, o)
	}

	return nil
}

// SetUser of the checkin to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserCheckins.
func (o *Checkin) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *UserDatum) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"checkin\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, checkinPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID
	if o.R == nil {
		o.R = &checkinR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userDatumR{
			UserCheckins: CheckinSlice{o},
		}
	} else {
		related.R.UserCheckins = append(related.R.UserCheckins, o)
	}

	return nil
}

// Checkins retrieves all the records using an executor.
func Checkins(mods ...qm.QueryMod) checkinQuery {
	mods = append(mods, qm.From("\"checkin\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"checkin\".*"})
	}

	return checkinQuery{q}
}

// FindCheckin retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCheckin(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Checkin, error) {
	checkinObj := &Checkin{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"checkin\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, checkinObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from checkin")
	}

	if err = checkinObj.doAfterSelectHooks(ctx, exec); err != nil {
		return checkinObj, err
	}

	return checkinObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Checkin) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no checkin provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreateAt.IsZero() {
			o.CreateAt = currTime
		}
		if o.UpdateAt.IsZero() {
			o.UpdateAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(checkinColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	checkinInsertCacheMut.RLock()
	cache, cached := checkinInsertCache[key]
	checkinInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			checkinAllColumns,
			checkinColumnsWithDefault,
			checkinColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(checkinType, checkinMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(checkinType, checkinMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"checkin\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"checkin\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into checkin")
	}

	if !cached {
		checkinInsertCacheMut.Lock()
		checkinInsertCache[key] = cache
		checkinInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Checkin.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Checkin) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdateAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	checkinUpdateCacheMut.RLock()
	cache, cached := checkinUpdateCache[key]
	checkinUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			checkinAllColumns,
			checkinPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update checkin, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"checkin\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, checkinPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(checkinType, checkinMapping, append(wl, checkinPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update checkin row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for checkin")
	}

	if !cached {
		checkinUpdateCacheMut.Lock()
		checkinUpdateCache[key] = cache
		checkinUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q checkinQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for checkin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for checkin")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CheckinSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), checkinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"checkin\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, checkinPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in checkin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all checkin")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Checkin) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no checkin provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreateAt.IsZero() {
			o.CreateAt = currTime
		}
		o.UpdateAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(checkinColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	checkinUpsertCacheMut.RLock()
	cache, cached := checkinUpsertCache[key]
	checkinUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			checkinAllColumns,
			checkinColumnsWithDefault,
			checkinColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			checkinAllColumns,
			checkinPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert checkin, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(checkinPrimaryKeyColumns))
			copy(conflict, checkinPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"checkin\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(checkinType, checkinMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(checkinType, checkinMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert checkin")
	}

	if !cached {
		checkinUpsertCacheMut.Lock()
		checkinUpsertCache[key] = cache
		checkinUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Checkin record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Checkin) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Checkin provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), checkinPrimaryKeyMapping)
	sql := "DELETE FROM \"checkin\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from checkin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for checkin")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q checkinQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no checkinQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from checkin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for checkin")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CheckinSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(checkinBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), checkinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"checkin\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, checkinPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from checkin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for checkin")
	}

	if len(checkinAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Checkin) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCheckin(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CheckinSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CheckinSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), checkinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"checkin\".* FROM \"checkin\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, checkinPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CheckinSlice")
	}

	*o = slice

	return nil
}

// CheckinExists checks if the Checkin row exists.
func CheckinExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"checkin\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if checkin exists")
	}

	return exists, nil
}

// Exists checks if the Checkin row exists.
func (o *Checkin) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CheckinExists(ctx, exec, o.ID)
}
