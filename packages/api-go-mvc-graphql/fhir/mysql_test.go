package fhir_test

import (
	"testing"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
)

func Test_MySQL(t *testing.T) {
	db, err := fhir.MySQL()
	if err != nil {
		t.Fatal(err)
	}

	if db == nil {
		t.Fatal("db is nil")
	}

	if db.Error != nil {
		t.Fatal(db.Error)
	}

	dbTwice, err := fhir.MySQL()

	if err != nil {
		t.Fatal(err)
	}

	if dbTwice == nil {
		t.Fatal("dbTwice is nil")
	}

	if dbTwice.Error != nil {
		t.Fatal(dbTwice.Error)
	}

	if db != dbTwice {
		t.Fatal("MySQL() should return same instance")
	}
}
