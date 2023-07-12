package domain_test

import (
	"testing"

	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain"
)

func Test_NewID(t *testing.T) {
	t.Parallel()

	id := domain.NewID()
	t.Log(id)
}
