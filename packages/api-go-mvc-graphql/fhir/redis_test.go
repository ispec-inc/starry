package fhir_test

import (
	"testing"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
)

func Test_Redis(t *testing.T) {
	r := fhir.Redis()

	r2 := fhir.Redis()

	if r != r2 {
		t.Fatal("Redis() should return same instance")
	}
}
