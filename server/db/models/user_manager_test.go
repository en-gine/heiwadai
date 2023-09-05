// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUserManagers(t *testing.T) {
	t.Parallel()

	query := UserManagers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserManagersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserManagersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserManagers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserManagersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserManagerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserManagersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserManagerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UserManager exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserManagerExists to return true, but got false.")
	}
}

func testUserManagersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userManagerFound, err := FindUserManager(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if userManagerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserManagersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserManagers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserManagersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserManagers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserManagersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userManagerOne := &UserManager{}
	userManagerTwo := &UserManager{}
	if err = randomize.Struct(seed, userManagerOne, userManagerDBTypes, false, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}
	if err = randomize.Struct(seed, userManagerTwo, userManagerDBTypes, false, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userManagerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userManagerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserManagers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserManagersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userManagerOne := &UserManager{}
	userManagerTwo := &UserManager{}
	if err = randomize.Struct(seed, userManagerOne, userManagerDBTypes, false, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}
	if err = randomize.Struct(seed, userManagerTwo, userManagerDBTypes, false, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userManagerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userManagerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userManagerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserManager) error {
	*o = UserManager{}
	return nil
}

func userManagerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserManager) error {
	*o = UserManager{}
	return nil
}

func userManagerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserManager) error {
	*o = UserManager{}
	return nil
}

func userManagerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserManager) error {
	*o = UserManager{}
	return nil
}

func userManagerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserManager) error {
	*o = UserManager{}
	return nil
}

func userManagerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserManager) error {
	*o = UserManager{}
	return nil
}

func userManagerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserManager) error {
	*o = UserManager{}
	return nil
}

func userManagerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserManager) error {
	*o = UserManager{}
	return nil
}

func userManagerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserManager) error {
	*o = UserManager{}
	return nil
}

func testUserManagersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserManager{}
	o := &UserManager{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userManagerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserManager object: %s", err)
	}

	AddUserManagerHook(boil.BeforeInsertHook, userManagerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userManagerBeforeInsertHooks = []UserManagerHook{}

	AddUserManagerHook(boil.AfterInsertHook, userManagerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userManagerAfterInsertHooks = []UserManagerHook{}

	AddUserManagerHook(boil.AfterSelectHook, userManagerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userManagerAfterSelectHooks = []UserManagerHook{}

	AddUserManagerHook(boil.BeforeUpdateHook, userManagerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userManagerBeforeUpdateHooks = []UserManagerHook{}

	AddUserManagerHook(boil.AfterUpdateHook, userManagerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userManagerAfterUpdateHooks = []UserManagerHook{}

	AddUserManagerHook(boil.BeforeDeleteHook, userManagerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userManagerBeforeDeleteHooks = []UserManagerHook{}

	AddUserManagerHook(boil.AfterDeleteHook, userManagerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userManagerAfterDeleteHooks = []UserManagerHook{}

	AddUserManagerHook(boil.BeforeUpsertHook, userManagerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userManagerBeforeUpsertHooks = []UserManagerHook{}

	AddUserManagerHook(boil.AfterUpsertHook, userManagerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userManagerAfterUpsertHooks = []UserManagerHook{}
}

func testUserManagersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserManagersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userManagerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserManagerOneToOneAdminUsingAdminAdmin(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var foreign Admin
	var local UserManager

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, adminDBTypes, true, adminColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Admin struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreign.AdminID = local.ID
	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.AdminAdmin().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.AdminID != foreign.AdminID {
		t.Errorf("want: %v, got %v", foreign.AdminID, check.AdminID)
	}

	ranAfterSelectHook := false
	AddAdminHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Admin) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := UserManagerSlice{&local}
	if err = local.L.LoadAdminAdmin(ctx, tx, false, (*[]*UserManager)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AdminAdmin == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AdminAdmin = nil
	if err = local.L.LoadAdminAdmin(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AdminAdmin == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testUserManagerOneToOneUserDatumUsingUserUserDatum(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var foreign UserDatum
	var local UserManager

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, userDatumDBTypes, true, userDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDatum struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreign.UserID = local.ID
	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.UserUserDatum().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UserID != foreign.UserID {
		t.Errorf("want: %v, got %v", foreign.UserID, check.UserID)
	}

	ranAfterSelectHook := false
	AddUserDatumHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *UserDatum) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := UserManagerSlice{&local}
	if err = local.L.LoadUserUserDatum(ctx, tx, false, (*[]*UserManager)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UserUserDatum == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.UserUserDatum = nil
	if err = local.L.LoadUserUserDatum(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UserUserDatum == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testUserManagerOneToOneSetOpAdminUsingAdminAdmin(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserManager
	var b, c Admin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userManagerDBTypes, false, strmangle.SetComplement(userManagerPrimaryKeyColumns, userManagerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, adminDBTypes, false, strmangle.SetComplement(adminPrimaryKeyColumns, adminColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, adminDBTypes, false, strmangle.SetComplement(adminPrimaryKeyColumns, adminColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Admin{&b, &c} {
		err = a.SetAdminAdmin(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AdminAdmin != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Admin != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.ID != x.AdminID {
			t.Error("foreign key was wrong value", a.ID)
		}

		if exists, err := AdminExists(ctx, tx, x.AdminID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'x' to exist")
		}

		if a.ID != x.AdminID {
			t.Error("foreign key was wrong value", a.ID, x.AdminID)
		}

		if _, err = x.Delete(ctx, tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testUserManagerOneToOneSetOpUserDatumUsingUserUserDatum(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserManager
	var b, c UserDatum

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userManagerDBTypes, false, strmangle.SetComplement(userManagerPrimaryKeyColumns, userManagerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDatumDBTypes, false, strmangle.SetComplement(userDatumPrimaryKeyColumns, userDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDatumDBTypes, false, strmangle.SetComplement(userDatumPrimaryKeyColumns, userDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UserDatum{&b, &c} {
		err = a.SetUserUserDatum(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.UserUserDatum != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.User != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.ID != x.UserID {
			t.Error("foreign key was wrong value", a.ID)
		}

		if exists, err := UserDatumExists(ctx, tx, x.UserID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'x' to exist")
		}

		if a.ID != x.UserID {
			t.Error("foreign key was wrong value", a.ID, x.UserID)
		}

		if _, err = x.Delete(ctx, tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testUserManagersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserManagersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserManagerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserManagersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserManagers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userManagerDBTypes = map[string]string{`ID`: `uuid`, `Email`: `character varying`, `IsAdmin`: `boolean`, `CreateAt`: `timestamp with time zone`, `UpdateAt`: `timestamp with time zone`}
	_                  = bytes.MinRead
)

func testUserManagersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userManagerAllColumns) == len(userManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserManagersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userManagerAllColumns) == len(userManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserManager{}
	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userManagerDBTypes, true, userManagerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userManagerAllColumns, userManagerPrimaryKeyColumns) {
		fields = userManagerAllColumns
	} else {
		fields = strmangle.SetComplement(
			userManagerAllColumns,
			userManagerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := UserManagerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserManagersUpsert(t *testing.T) {
	t.Parallel()

	if len(userManagerAllColumns) == len(userManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserManager{}
	if err = randomize.Struct(seed, &o, userManagerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserManager: %s", err)
	}

	count, err := UserManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userManagerDBTypes, false, userManagerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserManager struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserManager: %s", err)
	}

	count, err = UserManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}