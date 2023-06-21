package fhir_test

import (
	"testing"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
)

func TestRegistry(t *testing.T) {
	registry, err := fhir.NewRegistry()
	if err != nil {
		t.Fatal(err)
	}

	if registry.Q == nil {
		t.Fatal("Registry.Q is nil")
	}

	if registry.Cache == nil {
		t.Fatal("Registry.Cache is nil")
	}

	if registry.Validator == nil {
		t.Fatal("Registry.Validator is nil")
	}
}
