package controller

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/registry"
)

type Controller struct {
	registry registry.Registry
}

func New(registry registry.Registry) *Controller {
	return &Controller{
		registry: registry,
	}
}
