package api

import "github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic"

func NewPubSub() error {
	if err := clinic.NewEvent(); err != nil {
		return err
	}

	return clinic.NewCommand()
}
