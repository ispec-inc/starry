package organization

import (
	"fmt"

	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain"
)

const (
	maxOrganizationNameChars  int = 50
	maxOrganizationAliasChars int = 50
)

// Name 組織名の値オブジェクト
type Name struct {
	Name  domain.String
	Alias domain.String
}

func NewName(
	name string,
	alias string,
) (Name, error) {
	n := domain.String(name)
	if err := n.ValidateLength(1, maxOrganizationNameChars); err != nil {
		return Name{}, fmt.Errorf("name: %w", err)
	}

	a := domain.String(alias)
	if err := a.ValidateLength(1, maxOrganizationAliasChars); err != nil {
		return Name{}, fmt.Errorf("alias: %w", err)
	}

	return Name{
		Name:  n,
		Alias: a,
	}, nil
}
