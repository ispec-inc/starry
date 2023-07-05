package
type Controller struct {
	registry registry.Registry
}

func New() (*Controller, error) {
	rgst, err := registry.New()
	if err != nil {
		return nil, err
	}
	return &Controller{
		registry: rgst,
	}, nil
}

func (c Controller) Organization(ctx context.Context, args struct {
	ID graphql.ID
}) (resolver.Organization, error) {

	ipt := uc.GetOrganizationInput{
		ID: domain.ID(args.ID),
	}
	get := uc.NewGetOrganization(c.registry)

	opt, err := get.Do(ctx, ipt)
	if err != nil {
		return resolver.Organization{}, err
	}

	r := resolver.NewOrganization(opt.Organization)
	return r, nil
}
