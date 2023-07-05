package domain

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

type ID string

func NewID() (ID, error) {
	id := ulid.Make()

	return ID(fmt.Sprintf("%s", id)), nil
}
