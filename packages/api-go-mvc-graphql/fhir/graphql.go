package fhir

import "github.com/graph-gophers/graphql-go"

func GraphQLIDValue(id *graphql.ID) graphql.ID {
	if id == nil {
		return ""
	}

	return *id
}

func GraphQLID(id graphql.ID) *graphql.ID {
	return &id
}
