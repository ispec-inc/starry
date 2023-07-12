package domain_test

import (
	"errors"
	"testing"

	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain"
)

func Test_String(t *testing.T) {
	t.Parallel()

	s := domain.String("test")

	if err := s.ValidateLength(1, 10); err != nil {
		t.Fatal(err)
	}

	err := s.ValidateLength(1, 3)

	if !errors.Is(err, domain.ErrStringInvalidLength) {
		t.Fatalf("expected %v to be %v", err, domain.ErrStringInvalidLength)
	}

}
