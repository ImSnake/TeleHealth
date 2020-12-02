// Code generated by SQLBoiler 4.3.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testPatients(t *testing.T) {
	t.Parallel()

	query := Patients()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPatientsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
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

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatientsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Patients().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatientsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatientSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatientsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PatientExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Patient exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PatientExists to return true, but got false.")
	}
}

func testPatientsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	patientFound, err := FindPatient(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if patientFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPatientsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Patients().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPatientsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Patients().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPatientsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	patientOne := &Patient{}
	patientTwo := &Patient{}
	if err = randomize.Struct(seed, patientOne, patientDBTypes, false, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}
	if err = randomize.Struct(seed, patientTwo, patientDBTypes, false, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = patientOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = patientTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Patients().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPatientsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	patientOne := &Patient{}
	patientTwo := &Patient{}
	if err = randomize.Struct(seed, patientOne, patientDBTypes, false, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}
	if err = randomize.Struct(seed, patientTwo, patientDBTypes, false, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = patientOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = patientTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func patientBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func testPatientsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Patient{}
	o := &Patient{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, patientDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Patient object: %s", err)
	}

	AddPatientHook(boil.BeforeInsertHook, patientBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	patientBeforeInsertHooks = []PatientHook{}

	AddPatientHook(boil.AfterInsertHook, patientAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	patientAfterInsertHooks = []PatientHook{}

	AddPatientHook(boil.AfterSelectHook, patientAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	patientAfterSelectHooks = []PatientHook{}

	AddPatientHook(boil.BeforeUpdateHook, patientBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	patientBeforeUpdateHooks = []PatientHook{}

	AddPatientHook(boil.AfterUpdateHook, patientAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	patientAfterUpdateHooks = []PatientHook{}

	AddPatientHook(boil.BeforeDeleteHook, patientBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	patientBeforeDeleteHooks = []PatientHook{}

	AddPatientHook(boil.AfterDeleteHook, patientAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	patientAfterDeleteHooks = []PatientHook{}

	AddPatientHook(boil.BeforeUpsertHook, patientBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	patientBeforeUpsertHooks = []PatientHook{}

	AddPatientHook(boil.AfterUpsertHook, patientAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	patientAfterUpsertHooks = []PatientHook{}
}

func testPatientsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPatientsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(patientColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPatientsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
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

func testPatientsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatientSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPatientsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Patients().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	patientDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `Email`: `varchar`, `PWD`: `varchar`, `PhotoURL`: `text`, `Contacts`: `varchar`, `Balance`: `int`, `Doctor`: `tinyint`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_              = bytes.MinRead
)

func testPatientsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(patientPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(patientAllColumns) == len(patientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, patientDBTypes, true, patientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPatientsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(patientAllColumns) == len(patientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, patientDBTypes, true, patientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(patientAllColumns, patientPrimaryKeyColumns) {
		fields = patientAllColumns
	} else {
		fields = strmangle.SetComplement(
			patientAllColumns,
			patientPrimaryKeyColumns,
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

	slice := PatientSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPatientsUpsert(t *testing.T) {
	t.Parallel()

	if len(patientAllColumns) == len(patientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLPatientUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Patient{}
	if err = randomize.Struct(seed, &o, patientDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Patient: %s", err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, patientDBTypes, false, patientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Patient: %s", err)
	}

	count, err = Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
