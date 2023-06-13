package registry

type Registry struct {
	repo Repository
}

func New() (Registry, error) {
	repo, err := NewRepository()
	if err != nil {
		return Registry{}, err
	}

	return Registry{
		repo: repo,
	}, nil
}

func (r Registry) Repository() Repository {
	return r.repo
}
