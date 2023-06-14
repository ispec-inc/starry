package enum

type (
	OrganizationType int32
)

const (
	OrganizationTypeProv int32 = iota + 1
	OrganizationTypeOrg
)

func OrganizationTypeString(v int32) string {
	switch v {
	case OrganizationTypeProv:
		return "PROV"
	case OrganizationTypeOrg:
		return "ORG"
	default:
		return ""
	}
}

func NewOrganizationType(v string) int32 {
	switch v {
	case "PROV":
		return OrganizationTypeProv
	case "ORG":
		return OrganizationTypeOrg
	default:
		return 0
	}
}
