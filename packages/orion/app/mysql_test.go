package app_test

import (
	"testing"

	"github.com/ispec-inc/starry/orion/app"
	"github.com/ispec-inc/starry/orion/app/config"
)

func Test_MySQL(t *testing.T) {
	t.Parallel()

	if err := config.Init(); err != nil {
		t.Fatal(err)
	}

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

}
