package organization

// Type 組織の種別
type Type uint

const (
	TypeProv = "PROV"
	TypeOrg  = "ORG"
)

var (
	typeProv Type = 1 // TypeProv 事業者
	typeOrg  Type = 2 // TypeOrg 組織

	typeToString = map[Type]string{
		typeProv: TypeProv,
		typeOrg:  TypeOrg,
	}
)

// NewType 組織の種別を文字列から生成する
func NewType(v uint) (Type, error) {
	t := Type(v)
	if _, ok := typeToString[t]; !ok {
		return 0, ErrTypeStringIsInvalid
	}

	return t, nil
}

// String 文字列を返す
func (t Type) String() string {
	return typeToString[t]
}
