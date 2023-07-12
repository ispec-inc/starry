package app_test

import (
	"testing"

	"github.com/ispec-inc/starry/orion/app"
)

func Test_MySQL(t *testing.T) {
	t.Parallel()

	db, err := app.MySQL()
	if err != nil {
		t.Fatal(err)
	}

	if db == nil {
		t.Fatal("db is nil")
	}

	if db.Get().Error != nil {
		t.Fatal(db.Get().Error)
	}

	dbTwice, err := app.MySQL()

	if err != nil {
		t.Fatal(err)
	}

	if dbTwice == nil {
		t.Fatal("dbTwice is nil")
	}

	if dbTwice.Get().Error != nil {
		t.Fatal(dbTwice.Get().Error)
	}

	if db != dbTwice {
		t.Fatal("MySQL() should return same instance")
	}
}
