package postgresql

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"main/ExtError"
	"reflect"
	"testing"
)

func TestIsTableExistPositive(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error init sqlmoke: %s", err)
	}
	defer db.Close()

	tName := "test_table"

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(1)

	mock.ExpectQuery(`^SELECT (.+) FROM information_schema.tables `).WithArgs(tName).WillReturnRows(rows)

	var pg DBPostgresql
	pg.DB = db
	result, extErr := pg.isTableExist(tName)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expect check exist table with name %s: %s", tName, err)
	}
	if extErr != nil {
		t.Fatalf("Unexpected error (expect nil): %s", extErr)
	}
	if result != true {
		t.Fatalf("Unexpected result (expect true): %s", extErr)
	}
}

func TestIsTableExistNegative(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error init sqlmoke: %s", err)
	}
	defer db.Close()

	tName := "test_table"

	mock.ExpectQuery(`^SELECT (.+) FROM information_schema.tables `).
		WithArgs(tName).WillReturnError(fmt.Errorf("db error"))

	var pg DBPostgresql
	pg.DB = db
	_, extErr := pg.isTableExist(tName)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expect check exist table with name %s: %s", tName, err)
	}

	if extErr == nil {
		t.Fatalf("Expected error (unexpected nil): %s", extErr)
	}

	expExtErr := ExtError.Resend("Error check exist table "+tName, 1, fmt.Errorf("db error"))

	if !reflect.DeepEqual(extErr, expExtErr) {
		t.Fatalf("Expected error: %s current error: %s", expExtErr, extErr)
	}
}

func TestIsIndexExistPositive(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error init sqlmoke: %s", err)
	}
	defer db.Close()

	iName := "test_index"

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(1)

	mock.ExpectQuery(`^SELECT (.+) FROM pg_class `).WithArgs(iName).WillReturnRows(rows)

	var pg DBPostgresql
	pg.DB = db
	result, extErr := pg.isIndexExist(iName)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expect check exist index with name %s: %s", iName, err)
	}
	if extErr != nil {
		t.Fatalf("Unexpected error (expect nil): %s", extErr)
	}
	if result != true {
		t.Fatalf("Unexpected result (expect true): %s", extErr)
	}
}

func TestIsIndexExistNegative(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error init sqlmoke: %s", err)
	}
	defer db.Close()

	iName := "test_index"

	mock.ExpectQuery(`^SELECT (.+) FROM pg_class `).
		WithArgs(iName).WillReturnError(fmt.Errorf("db error"))

	var pg DBPostgresql
	pg.DB = db
	_, extErr := pg.isIndexExist(iName)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expect check exist index with name %s: %s", iName, err)
	}

	if extErr == nil {
		t.Fatalf("Expected error (unexpected nil): %s", extErr)
	}

	expExtErr := ExtError.Resend("Error check exist index "+iName, 1, fmt.Errorf("db error"))

	if !reflect.DeepEqual(extErr, expExtErr) {
		t.Fatalf("Expected error: %s current error: %s", expExtErr, extErr)
	}
}

func TestCreatePositive(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error init sqlmoke: %s", err)
	}
	defer db.Close()

	create := "create query"
	mock.ExpectExec(create).WillReturnResult(sqlmock.NewResult(0, 0))

	var pg DBPostgresql
	pg.DB = db
	extErr := pg.create(create)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expect query '%s': %s", create, err)
	}
	if extErr != nil {
		t.Fatalf("Unexpected error (expect nil): %s", extErr)
	}
}

func TestCreateNegative(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error init sqlmoke: %s", err)
	}
	defer db.Close()

	create := "create query"
	mock.ExpectExec(create).WillReturnError(fmt.Errorf("db error"))

	var pg DBPostgresql
	pg.DB = db
	extErr := pg.create(create)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expect query '%s': %s", create, err)
	}

	if extErr == nil {
		t.Fatalf("Expected error (unexpected nil): %s", extErr)
	}

	expExtErr := ExtError.Resend("Error create "+create, 1, fmt.Errorf("db error"))

	if !reflect.DeepEqual(extErr, expExtErr) {
		t.Fatalf("Expected error: %s current error: %s", expExtErr, extErr)
	}
}
