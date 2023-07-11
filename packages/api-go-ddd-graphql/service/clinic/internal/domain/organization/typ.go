package organization

// Type 組織の種別
type Type int

var (
	TypeProv int32 = 1 // TypeProv 事業者
	TypeOrg  int32 = 2 // TypeOrg 組織
)

// TypeString 組織の種別を文字列に変換する
func TypeString(v int32) string {
	switch v {
	case TypeProv:
		return "PROV"
	case TypeOrg:
		return "ORG"
	default:
		return ""
	}
}

// NewType 組織の種別を文字列から生成する
func NewType(v string) int32 {
	switch v {
	case "PROV":
		return TypeProv
	case "ORG":
		return TypeOrg
	default:
		return 0
	}
}
