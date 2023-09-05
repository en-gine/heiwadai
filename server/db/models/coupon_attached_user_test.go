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

func testCouponAttachedUsers(t *testing.T) {
	t.Parallel()

	query := CouponAttachedUsers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCouponAttachedUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
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

	count, err := CouponAttachedUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCouponAttachedUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CouponAttachedUsers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CouponAttachedUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCouponAttachedUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CouponAttachedUserSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CouponAttachedUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCouponAttachedUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CouponAttachedUserExists(ctx, tx, o.CouponID, o.UserID)
	if err != nil {
		t.Errorf("Unable to check if CouponAttachedUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CouponAttachedUserExists to return true, but got false.")
	}
}

func testCouponAttachedUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	couponAttachedUserFound, err := FindCouponAttachedUser(ctx, tx, o.CouponID, o.UserID)
	if err != nil {
		t.Error(err)
	}

	if couponAttachedUserFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCouponAttachedUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CouponAttachedUsers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCouponAttachedUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CouponAttachedUsers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCouponAttachedUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	couponAttachedUserOne := &CouponAttachedUser{}
	couponAttachedUserTwo := &CouponAttachedUser{}
	if err = randomize.Struct(seed, couponAttachedUserOne, couponAttachedUserDBTypes, false, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}
	if err = randomize.Struct(seed, couponAttachedUserTwo, couponAttachedUserDBTypes, false, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = couponAttachedUserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = couponAttachedUserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CouponAttachedUsers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCouponAttachedUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	couponAttachedUserOne := &CouponAttachedUser{}
	couponAttachedUserTwo := &CouponAttachedUser{}
	if err = randomize.Struct(seed, couponAttachedUserOne, couponAttachedUserDBTypes, false, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}
	if err = randomize.Struct(seed, couponAttachedUserTwo, couponAttachedUserDBTypes, false, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = couponAttachedUserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = couponAttachedUserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CouponAttachedUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func couponAttachedUserBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *CouponAttachedUser) error {
	*o = CouponAttachedUser{}
	return nil
}

func couponAttachedUserAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *CouponAttachedUser) error {
	*o = CouponAttachedUser{}
	return nil
}

func couponAttachedUserAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *CouponAttachedUser) error {
	*o = CouponAttachedUser{}
	return nil
}

func couponAttachedUserBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CouponAttachedUser) error {
	*o = CouponAttachedUser{}
	return nil
}

func couponAttachedUserAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CouponAttachedUser) error {
	*o = CouponAttachedUser{}
	return nil
}

func couponAttachedUserBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CouponAttachedUser) error {
	*o = CouponAttachedUser{}
	return nil
}

func couponAttachedUserAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CouponAttachedUser) error {
	*o = CouponAttachedUser{}
	return nil
}

func couponAttachedUserBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CouponAttachedUser) error {
	*o = CouponAttachedUser{}
	return nil
}

func couponAttachedUserAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CouponAttachedUser) error {
	*o = CouponAttachedUser{}
	return nil
}

func testCouponAttachedUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &CouponAttachedUser{}
	o := &CouponAttachedUser{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser object: %s", err)
	}

	AddCouponAttachedUserHook(boil.BeforeInsertHook, couponAttachedUserBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	couponAttachedUserBeforeInsertHooks = []CouponAttachedUserHook{}

	AddCouponAttachedUserHook(boil.AfterInsertHook, couponAttachedUserAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	couponAttachedUserAfterInsertHooks = []CouponAttachedUserHook{}

	AddCouponAttachedUserHook(boil.AfterSelectHook, couponAttachedUserAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	couponAttachedUserAfterSelectHooks = []CouponAttachedUserHook{}

	AddCouponAttachedUserHook(boil.BeforeUpdateHook, couponAttachedUserBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	couponAttachedUserBeforeUpdateHooks = []CouponAttachedUserHook{}

	AddCouponAttachedUserHook(boil.AfterUpdateHook, couponAttachedUserAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	couponAttachedUserAfterUpdateHooks = []CouponAttachedUserHook{}

	AddCouponAttachedUserHook(boil.BeforeDeleteHook, couponAttachedUserBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	couponAttachedUserBeforeDeleteHooks = []CouponAttachedUserHook{}

	AddCouponAttachedUserHook(boil.AfterDeleteHook, couponAttachedUserAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	couponAttachedUserAfterDeleteHooks = []CouponAttachedUserHook{}

	AddCouponAttachedUserHook(boil.BeforeUpsertHook, couponAttachedUserBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	couponAttachedUserBeforeUpsertHooks = []CouponAttachedUserHook{}

	AddCouponAttachedUserHook(boil.AfterUpsertHook, couponAttachedUserAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	couponAttachedUserAfterUpsertHooks = []CouponAttachedUserHook{}
}

func testCouponAttachedUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CouponAttachedUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCouponAttachedUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(couponAttachedUserColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CouponAttachedUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCouponAttachedUserToOneCouponUsingCoupon(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CouponAttachedUser
	var foreign Coupon

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, couponAttachedUserDBTypes, false, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, couponDBTypes, false, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CouponID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Coupon().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddCouponHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Coupon) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := CouponAttachedUserSlice{&local}
	if err = local.L.LoadCoupon(ctx, tx, false, (*[]*CouponAttachedUser)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Coupon == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Coupon = nil
	if err = local.L.LoadCoupon(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Coupon == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testCouponAttachedUserToOneUserDatumUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CouponAttachedUser
	var foreign UserDatum

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, couponAttachedUserDBTypes, false, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDatumDBTypes, false, userDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDatum struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.UserID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
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

	slice := CouponAttachedUserSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*CouponAttachedUser)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testCouponAttachedUserToOneSetOpCouponUsingCoupon(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CouponAttachedUser
	var b, c Coupon

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, couponAttachedUserDBTypes, false, strmangle.SetComplement(couponAttachedUserPrimaryKeyColumns, couponAttachedUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, couponDBTypes, false, strmangle.SetComplement(couponPrimaryKeyColumns, couponColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, couponDBTypes, false, strmangle.SetComplement(couponPrimaryKeyColumns, couponColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Coupon{&b, &c} {
		err = a.SetCoupon(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Coupon != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CouponAttachedUsers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CouponID != x.ID {
			t.Error("foreign key was wrong value", a.CouponID)
		}

		if exists, err := CouponAttachedUserExists(ctx, tx, a.CouponID, a.UserID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testCouponAttachedUserToOneSetOpUserDatumUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CouponAttachedUser
	var b, c UserDatum

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, couponAttachedUserDBTypes, false, strmangle.SetComplement(couponAttachedUserPrimaryKeyColumns, couponAttachedUserColumnsWithoutDefault)...); err != nil {
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
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserCouponAttachedUsers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.UserID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := CouponAttachedUserExists(ctx, tx, a.CouponID, a.UserID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testCouponAttachedUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
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

func testCouponAttachedUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CouponAttachedUserSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCouponAttachedUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CouponAttachedUsers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	couponAttachedUserDBTypes = map[string]string{`CouponID`: `uuid`, `UsedAt`: `timestamp with time zone`, `UserID`: `uuid`}
	_                         = bytes.MinRead
)

func testCouponAttachedUsersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(couponAttachedUserPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(couponAttachedUserAllColumns) == len(couponAttachedUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CouponAttachedUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCouponAttachedUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(couponAttachedUserAllColumns) == len(couponAttachedUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CouponAttachedUser{}
	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CouponAttachedUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, couponAttachedUserDBTypes, true, couponAttachedUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(couponAttachedUserAllColumns, couponAttachedUserPrimaryKeyColumns) {
		fields = couponAttachedUserAllColumns
	} else {
		fields = strmangle.SetComplement(
			couponAttachedUserAllColumns,
			couponAttachedUserPrimaryKeyColumns,
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

	slice := CouponAttachedUserSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCouponAttachedUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(couponAttachedUserAllColumns) == len(couponAttachedUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CouponAttachedUser{}
	if err = randomize.Struct(seed, &o, couponAttachedUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CouponAttachedUser: %s", err)
	}

	count, err := CouponAttachedUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, couponAttachedUserDBTypes, false, couponAttachedUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CouponAttachedUser struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CouponAttachedUser: %s", err)
	}

	count, err = CouponAttachedUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
