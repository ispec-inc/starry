package registry

// Registry インターフェースに対して実装を提供するレジストリ
type Registry struct {
	repo Repository
}

// New Registryのコンストラクタ
func New() (Registry, error) {
	repo, err := NewRepository()
	if err != nil {
		return Registry{}, err
	}

	return Registry{
		repo: repo,
	}, nil
}

// Repository リポジトリのレジストリを返す
func (r Registry) Repository() Repository {
	return r.repo
}
