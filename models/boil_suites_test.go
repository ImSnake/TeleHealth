// Code generated by SQLBoiler 4.3.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Analyzes", testAnalyzes)
	t.Run("Doctors", testDoctors)
	t.Run("Patients", testPatients)
	t.Run("Receptions", testReceptions)
	t.Run("Reviews", testReviews)
	t.Run("SpecializationKits", testSpecializationKits)
	t.Run("Specializations", testSpecializations)
	t.Run("Timetables", testTimetables)
}

func TestDelete(t *testing.T) {
	t.Run("Analyzes", testAnalyzesDelete)
	t.Run("Doctors", testDoctorsDelete)
	t.Run("Patients", testPatientsDelete)
	t.Run("Receptions", testReceptionsDelete)
	t.Run("Reviews", testReviewsDelete)
	t.Run("SpecializationKits", testSpecializationKitsDelete)
	t.Run("Specializations", testSpecializationsDelete)
	t.Run("Timetables", testTimetablesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Analyzes", testAnalyzesQueryDeleteAll)
	t.Run("Doctors", testDoctorsQueryDeleteAll)
	t.Run("Patients", testPatientsQueryDeleteAll)
	t.Run("Receptions", testReceptionsQueryDeleteAll)
	t.Run("Reviews", testReviewsQueryDeleteAll)
	t.Run("SpecializationKits", testSpecializationKitsQueryDeleteAll)
	t.Run("Specializations", testSpecializationsQueryDeleteAll)
	t.Run("Timetables", testTimetablesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Analyzes", testAnalyzesSliceDeleteAll)
	t.Run("Doctors", testDoctorsSliceDeleteAll)
	t.Run("Patients", testPatientsSliceDeleteAll)
	t.Run("Receptions", testReceptionsSliceDeleteAll)
	t.Run("Reviews", testReviewsSliceDeleteAll)
	t.Run("SpecializationKits", testSpecializationKitsSliceDeleteAll)
	t.Run("Specializations", testSpecializationsSliceDeleteAll)
	t.Run("Timetables", testTimetablesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Analyzes", testAnalyzesExists)
	t.Run("Doctors", testDoctorsExists)
	t.Run("Patients", testPatientsExists)
	t.Run("Receptions", testReceptionsExists)
	t.Run("Reviews", testReviewsExists)
	t.Run("SpecializationKits", testSpecializationKitsExists)
	t.Run("Specializations", testSpecializationsExists)
	t.Run("Timetables", testTimetablesExists)
}

func TestFind(t *testing.T) {
	t.Run("Analyzes", testAnalyzesFind)
	t.Run("Doctors", testDoctorsFind)
	t.Run("Patients", testPatientsFind)
	t.Run("Receptions", testReceptionsFind)
	t.Run("Reviews", testReviewsFind)
	t.Run("SpecializationKits", testSpecializationKitsFind)
	t.Run("Specializations", testSpecializationsFind)
	t.Run("Timetables", testTimetablesFind)
}

func TestBind(t *testing.T) {
	t.Run("Analyzes", testAnalyzesBind)
	t.Run("Doctors", testDoctorsBind)
	t.Run("Patients", testPatientsBind)
	t.Run("Receptions", testReceptionsBind)
	t.Run("Reviews", testReviewsBind)
	t.Run("SpecializationKits", testSpecializationKitsBind)
	t.Run("Specializations", testSpecializationsBind)
	t.Run("Timetables", testTimetablesBind)
}

func TestOne(t *testing.T) {
	t.Run("Analyzes", testAnalyzesOne)
	t.Run("Doctors", testDoctorsOne)
	t.Run("Patients", testPatientsOne)
	t.Run("Receptions", testReceptionsOne)
	t.Run("Reviews", testReviewsOne)
	t.Run("SpecializationKits", testSpecializationKitsOne)
	t.Run("Specializations", testSpecializationsOne)
	t.Run("Timetables", testTimetablesOne)
}

func TestAll(t *testing.T) {
	t.Run("Analyzes", testAnalyzesAll)
	t.Run("Doctors", testDoctorsAll)
	t.Run("Patients", testPatientsAll)
	t.Run("Receptions", testReceptionsAll)
	t.Run("Reviews", testReviewsAll)
	t.Run("SpecializationKits", testSpecializationKitsAll)
	t.Run("Specializations", testSpecializationsAll)
	t.Run("Timetables", testTimetablesAll)
}

func TestCount(t *testing.T) {
	t.Run("Analyzes", testAnalyzesCount)
	t.Run("Doctors", testDoctorsCount)
	t.Run("Patients", testPatientsCount)
	t.Run("Receptions", testReceptionsCount)
	t.Run("Reviews", testReviewsCount)
	t.Run("SpecializationKits", testSpecializationKitsCount)
	t.Run("Specializations", testSpecializationsCount)
	t.Run("Timetables", testTimetablesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Analyzes", testAnalyzesHooks)
	t.Run("Doctors", testDoctorsHooks)
	t.Run("Patients", testPatientsHooks)
	t.Run("Receptions", testReceptionsHooks)
	t.Run("Reviews", testReviewsHooks)
	t.Run("SpecializationKits", testSpecializationKitsHooks)
	t.Run("Specializations", testSpecializationsHooks)
	t.Run("Timetables", testTimetablesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Analyzes", testAnalyzesInsert)
	t.Run("Analyzes", testAnalyzesInsertWhitelist)
	t.Run("Doctors", testDoctorsInsert)
	t.Run("Doctors", testDoctorsInsertWhitelist)
	t.Run("Patients", testPatientsInsert)
	t.Run("Patients", testPatientsInsertWhitelist)
	t.Run("Receptions", testReceptionsInsert)
	t.Run("Receptions", testReceptionsInsertWhitelist)
	t.Run("Reviews", testReviewsInsert)
	t.Run("Reviews", testReviewsInsertWhitelist)
	t.Run("SpecializationKits", testSpecializationKitsInsert)
	t.Run("SpecializationKits", testSpecializationKitsInsertWhitelist)
	t.Run("Specializations", testSpecializationsInsert)
	t.Run("Specializations", testSpecializationsInsertWhitelist)
	t.Run("Timetables", testTimetablesInsert)
	t.Run("Timetables", testTimetablesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Analyzes", testAnalyzesReload)
	t.Run("Doctors", testDoctorsReload)
	t.Run("Patients", testPatientsReload)
	t.Run("Receptions", testReceptionsReload)
	t.Run("Reviews", testReviewsReload)
	t.Run("SpecializationKits", testSpecializationKitsReload)
	t.Run("Specializations", testSpecializationsReload)
	t.Run("Timetables", testTimetablesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Analyzes", testAnalyzesReloadAll)
	t.Run("Doctors", testDoctorsReloadAll)
	t.Run("Patients", testPatientsReloadAll)
	t.Run("Receptions", testReceptionsReloadAll)
	t.Run("Reviews", testReviewsReloadAll)
	t.Run("SpecializationKits", testSpecializationKitsReloadAll)
	t.Run("Specializations", testSpecializationsReloadAll)
	t.Run("Timetables", testTimetablesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Analyzes", testAnalyzesSelect)
	t.Run("Doctors", testDoctorsSelect)
	t.Run("Patients", testPatientsSelect)
	t.Run("Receptions", testReceptionsSelect)
	t.Run("Reviews", testReviewsSelect)
	t.Run("SpecializationKits", testSpecializationKitsSelect)
	t.Run("Specializations", testSpecializationsSelect)
	t.Run("Timetables", testTimetablesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Analyzes", testAnalyzesUpdate)
	t.Run("Doctors", testDoctorsUpdate)
	t.Run("Patients", testPatientsUpdate)
	t.Run("Receptions", testReceptionsUpdate)
	t.Run("Reviews", testReviewsUpdate)
	t.Run("SpecializationKits", testSpecializationKitsUpdate)
	t.Run("Specializations", testSpecializationsUpdate)
	t.Run("Timetables", testTimetablesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Analyzes", testAnalyzesSliceUpdateAll)
	t.Run("Doctors", testDoctorsSliceUpdateAll)
	t.Run("Patients", testPatientsSliceUpdateAll)
	t.Run("Receptions", testReceptionsSliceUpdateAll)
	t.Run("Reviews", testReviewsSliceUpdateAll)
	t.Run("SpecializationKits", testSpecializationKitsSliceUpdateAll)
	t.Run("Specializations", testSpecializationsSliceUpdateAll)
	t.Run("Timetables", testTimetablesSliceUpdateAll)
}
