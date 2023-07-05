package reader_test

import (
	"os"
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
)

func TestMain(m *testing.M) {
	_, err := app.MySQL()
	if err != nil {
		os.Exit(1)
	}
	m.Run()
}
