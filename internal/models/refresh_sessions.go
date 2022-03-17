// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// RefreshSession is an object representing the database table.
type RefreshSession struct {
	RefreshSecret string    `boil:"refresh_secret" json:"refresh_secret" toml:"refresh_secret" yaml:"refresh_secret"`
	UserID        int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	IP            string    `boil:"ip" json:"ip" toml:"ip" yaml:"ip"`
	UserAgent     string    `boil:"user_agent" json:"user_agent" toml:"user_agent" yaml:"user_agent"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *refreshSessionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L refreshSessionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RefreshSessionColumns = struct {
	RefreshSecret string
	UserID        string
	IP            string
	UserAgent     string
	CreatedAt     string
}{
	RefreshSecret: "refresh_secret",
	UserID:        "user_id",
	IP:            "ip",
	UserAgent:     "user_agent",
	CreatedAt:     "created_at",
}

var RefreshSessionTableColumns = struct {
	RefreshSecret string
	UserID        string
	IP            string
	UserAgent     string
	CreatedAt     string
}{
	RefreshSecret: "refresh_sessions.refresh_secret",
	UserID:        "refresh_sessions.user_id",
	IP:            "refresh_sessions.ip",
	UserAgent:     "refresh_sessions.user_agent",
	CreatedAt:     "refresh_sessions.created_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var RefreshSessionWhere = struct {
	RefreshSecret whereHelperstring
	UserID        whereHelperint
	IP            whereHelperstring
	UserAgent     whereHelperstring
	CreatedAt     whereHelpertime_Time
}{
	RefreshSecret: whereHelperstring{field: "\"refresh_sessions\".\"refresh_secret\""},
	UserID:        whereHelperint{field: "\"refresh_sessions\".\"user_id\""},
	IP:            whereHelperstring{field: "\"refresh_sessions\".\"ip\""},
	UserAgent:     whereHelperstring{field: "\"refresh_sessions\".\"user_agent\""},
	CreatedAt:     whereHelpertime_Time{field: "\"refresh_sessions\".\"created_at\""},
}

// RefreshSessionRels is where relationship names are stored.
var RefreshSessionRels = struct {
}{}

// refreshSessionR is where relationships are stored.
type refreshSessionR struct {
}

// NewStruct creates a new relationship struct
func (*refreshSessionR) NewStruct() *refreshSessionR {
	return &refreshSessionR{}
}

// refreshSessionL is where Load methods for each relationship are stored.
type refreshSessionL struct{}

var (
	refreshSessionAllColumns            = []string{"refresh_secret", "user_id", "ip", "user_agent", "created_at"}
	refreshSessionColumnsWithoutDefault = []string{"refresh_secret", "user_id", "ip", "user_agent"}
	refreshSessionColumnsWithDefault    = []string{"created_at"}
	refreshSessionPrimaryKeyColumns     = []string{"refresh_secret"}
	refreshSessionGeneratedColumns      = []string{}
)

type (
	// RefreshSessionSlice is an alias for a slice of pointers to RefreshSession.
	// This should almost always be used instead of []RefreshSession.
	RefreshSessionSlice []*RefreshSession
	// RefreshSessionHook is the signature for custom RefreshSession hook methods
	RefreshSessionHook func(context.Context, boil.ContextExecutor, *RefreshSession) error

	refreshSessionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	refreshSessionType                 = reflect.TypeOf(&RefreshSession{})
	refreshSessionMapping              = queries.MakeStructMapping(refreshSessionType)
	refreshSessionPrimaryKeyMapping, _ = queries.BindMapping(refreshSessionType, refreshSessionMapping, refreshSessionPrimaryKeyColumns)
	refreshSessionInsertCacheMut       sync.RWMutex
	refreshSessionInsertCache          = make(map[string]insertCache)
	refreshSessionUpdateCacheMut       sync.RWMutex
	refreshSessionUpdateCache          = make(map[string]updateCache)
	refreshSessionUpsertCacheMut       sync.RWMutex
	refreshSessionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var refreshSessionAfterSelectHooks []RefreshSessionHook

var refreshSessionBeforeInsertHooks []RefreshSessionHook
var refreshSessionAfterInsertHooks []RefreshSessionHook

var refreshSessionBeforeUpdateHooks []RefreshSessionHook
var refreshSessionAfterUpdateHooks []RefreshSessionHook

var refreshSessionBeforeDeleteHooks []RefreshSessionHook
var refreshSessionAfterDeleteHooks []RefreshSessionHook

var refreshSessionBeforeUpsertHooks []RefreshSessionHook
var refreshSessionAfterUpsertHooks []RefreshSessionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RefreshSession) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshSessionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RefreshSession) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshSessionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RefreshSession) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshSessionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RefreshSession) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshSessionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RefreshSession) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshSessionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RefreshSession) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshSessionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RefreshSession) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshSessionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RefreshSession) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshSessionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RefreshSession) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range refreshSessionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRefreshSessionHook registers your hook function for all future operations.
func AddRefreshSessionHook(hookPoint boil.HookPoint, refreshSessionHook RefreshSessionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		refreshSessionAfterSelectHooks = append(refreshSessionAfterSelectHooks, refreshSessionHook)
	case boil.BeforeInsertHook:
		refreshSessionBeforeInsertHooks = append(refreshSessionBeforeInsertHooks, refreshSessionHook)
	case boil.AfterInsertHook:
		refreshSessionAfterInsertHooks = append(refreshSessionAfterInsertHooks, refreshSessionHook)
	case boil.BeforeUpdateHook:
		refreshSessionBeforeUpdateHooks = append(refreshSessionBeforeUpdateHooks, refreshSessionHook)
	case boil.AfterUpdateHook:
		refreshSessionAfterUpdateHooks = append(refreshSessionAfterUpdateHooks, refreshSessionHook)
	case boil.BeforeDeleteHook:
		refreshSessionBeforeDeleteHooks = append(refreshSessionBeforeDeleteHooks, refreshSessionHook)
	case boil.AfterDeleteHook:
		refreshSessionAfterDeleteHooks = append(refreshSessionAfterDeleteHooks, refreshSessionHook)
	case boil.BeforeUpsertHook:
		refreshSessionBeforeUpsertHooks = append(refreshSessionBeforeUpsertHooks, refreshSessionHook)
	case boil.AfterUpsertHook:
		refreshSessionAfterUpsertHooks = append(refreshSessionAfterUpsertHooks, refreshSessionHook)
	}
}

// One returns a single refreshSession record from the query.
func (q refreshSessionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RefreshSession, error) {
	o := &RefreshSession{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for refresh_sessions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RefreshSession records from the query.
func (q refreshSessionQuery) All(ctx context.Context, exec boil.ContextExecutor) (RefreshSessionSlice, error) {
	var o []*RefreshSession

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RefreshSession slice")
	}

	if len(refreshSessionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RefreshSession records in the query.
func (q refreshSessionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count refresh_sessions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q refreshSessionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if refresh_sessions exists")
	}

	return count > 0, nil
}

// RefreshSessions retrieves all the records using an executor.
func RefreshSessions(mods ...qm.QueryMod) refreshSessionQuery {
	mods = append(mods, qm.From("\"refresh_sessions\""))
	return refreshSessionQuery{NewQuery(mods...)}
}

// FindRefreshSession retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRefreshSession(ctx context.Context, exec boil.ContextExecutor, refreshSecret string, selectCols ...string) (*RefreshSession, error) {
	refreshSessionObj := &RefreshSession{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"refresh_sessions\" where \"refresh_secret\"=$1", sel,
	)

	q := queries.Raw(query, refreshSecret)

	err := q.Bind(ctx, exec, refreshSessionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from refresh_sessions")
	}

	if err = refreshSessionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return refreshSessionObj, err
	}

	return refreshSessionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RefreshSession) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no refresh_sessions provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(refreshSessionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	refreshSessionInsertCacheMut.RLock()
	cache, cached := refreshSessionInsertCache[key]
	refreshSessionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			refreshSessionAllColumns,
			refreshSessionColumnsWithDefault,
			refreshSessionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(refreshSessionType, refreshSessionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(refreshSessionType, refreshSessionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"refresh_sessions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"refresh_sessions\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into refresh_sessions")
	}

	if !cached {
		refreshSessionInsertCacheMut.Lock()
		refreshSessionInsertCache[key] = cache
		refreshSessionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RefreshSession.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RefreshSession) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	refreshSessionUpdateCacheMut.RLock()
	cache, cached := refreshSessionUpdateCache[key]
	refreshSessionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			refreshSessionAllColumns,
			refreshSessionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update refresh_sessions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"refresh_sessions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, refreshSessionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(refreshSessionType, refreshSessionMapping, append(wl, refreshSessionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update refresh_sessions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for refresh_sessions")
	}

	if !cached {
		refreshSessionUpdateCacheMut.Lock()
		refreshSessionUpdateCache[key] = cache
		refreshSessionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q refreshSessionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for refresh_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for refresh_sessions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RefreshSessionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), refreshSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"refresh_sessions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, refreshSessionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in refreshSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all refreshSession")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RefreshSession) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no refresh_sessions provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(refreshSessionColumnsWithDefault, o)

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

	refreshSessionUpsertCacheMut.RLock()
	cache, cached := refreshSessionUpsertCache[key]
	refreshSessionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			refreshSessionAllColumns,
			refreshSessionColumnsWithDefault,
			refreshSessionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			refreshSessionAllColumns,
			refreshSessionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert refresh_sessions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(refreshSessionPrimaryKeyColumns))
			copy(conflict, refreshSessionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"refresh_sessions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(refreshSessionType, refreshSessionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(refreshSessionType, refreshSessionMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert refresh_sessions")
	}

	if !cached {
		refreshSessionUpsertCacheMut.Lock()
		refreshSessionUpsertCache[key] = cache
		refreshSessionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RefreshSession record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RefreshSession) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RefreshSession provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), refreshSessionPrimaryKeyMapping)
	sql := "DELETE FROM \"refresh_sessions\" WHERE \"refresh_secret\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from refresh_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for refresh_sessions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q refreshSessionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no refreshSessionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from refresh_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for refresh_sessions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RefreshSessionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(refreshSessionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), refreshSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"refresh_sessions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, refreshSessionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from refreshSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for refresh_sessions")
	}

	if len(refreshSessionAfterDeleteHooks) != 0 {
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
func (o *RefreshSession) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRefreshSession(ctx, exec, o.RefreshSecret)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RefreshSessionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RefreshSessionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), refreshSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"refresh_sessions\".* FROM \"refresh_sessions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, refreshSessionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RefreshSessionSlice")
	}

	*o = slice

	return nil
}

// RefreshSessionExists checks if the RefreshSession row exists.
func RefreshSessionExists(ctx context.Context, exec boil.ContextExecutor, refreshSecret string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"refresh_sessions\" where \"refresh_secret\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, refreshSecret)
	}
	row := exec.QueryRowContext(ctx, sql, refreshSecret)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if refresh_sessions exists")
	}

	return exists, nil
}
