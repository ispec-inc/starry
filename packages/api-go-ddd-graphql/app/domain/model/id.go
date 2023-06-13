package model

import (
	"errors"
	"fmt"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/oklog/ulid/v2"
)

type ID string

func NewID(name string) (ID, error) {
	if name == "" {
		return "", app.NewError(errors.New("id: name is empty"))
	}
	id := ulid.Make()

	return ID(fmt.Sprintf("%s/%s", name, id)), nil
}
