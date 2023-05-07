package enum

type (
	OrganizationType int32
)

var (
	OrganizationTypeProv int32 = 1
	OrganizationTypeOrg  int32 = 2
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
