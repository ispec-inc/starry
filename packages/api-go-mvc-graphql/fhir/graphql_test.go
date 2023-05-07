package fhir_test

import (
	"testing"

	"github.com/graph-gophers/graphql-go"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
)

func Test_GraphQLIDValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		id   *graphql.ID
		want graphql.ID
	}{
		{
			name: "success",
			id:   fhir.GraphQLID(graphql.ID("test")),
			want: "test",
		},
		{
			name: "success",
			id:   nil,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fhir.GraphQLIDValue(tt.id); got != tt.want {
				t.Errorf("GraphQLIDValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
