package organization

// OrganizationType 組織の種別
type OrganizationType int

var (
	OrganizationTypeProv int32 = 1 // OrganizationTypeProv 事業者
	OrganizationTypeOrg  int32 = 2 // OrganizationTypeOrg 組織
)

// OrganizationTypeString 組織の種別を文字列に変換する
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

// NewType 組織の種別を文字列から生成する
func NewType(v string) int32 {
	switch v {
	case "PROV":
		return OrganizationTypeProv
	case "ORG":
		return OrganizationTypeOrg
	default:
		return 0
	}
}
