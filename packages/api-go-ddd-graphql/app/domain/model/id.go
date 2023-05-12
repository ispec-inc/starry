package model

import (
	"errors"
	"fmt"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/apperror"
	"github.com/oklog/ulid/v2"
)

type ID string

func NewID(name string) (ID, error) {
	if name == "" {
		return "", apperror.New(errors.New("id: name is empty"))
	}
	id := ulid.Make()

	return ID(fmt.Sprintf("%s/%s", name, id)), nil
}
