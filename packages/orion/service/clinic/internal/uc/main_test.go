package uc_test

import (
	"os"
	"testing"

	"github.com/ispec-inc/starry/orion/app"
	"github.com/ispec-inc/starry/orion/app/config"
	"github.com/ispec-inc/starry/orion/app/test"
)

func TestMain(m *testing.M) {
	if err := config.Init(); err != nil {
		os.Exit(1)
	}

	_, err := app.MySQL()
	if err != nil {
		os.Exit(1)
	}

	test.InitMySQL()
	m.Run()
}
