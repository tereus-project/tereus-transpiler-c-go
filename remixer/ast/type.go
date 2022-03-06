package ast

type ASTTypeKind int64

const (
	ASTTypeKindVoid ASTTypeKind = iota
	ASTTypeKindInt
	ASTTypeKindInt16
	ASTTypeKindInt64
	ASTTypeKindRune
	ASTTypeKindFloat32
	ASTTypeKindFloat64

	ASTTypeKindStruct
)

type ASTType struct {
	Kind ASTTypeKind
	Name string
}

func NewASTType(kind ASTTypeKind, name string) *ASTType {
	return &ASTType{
		Kind: kind,
		Name: name,
	}
}

func (t *ASTType) String() string {
	return t.Name
}

func (t *ASTType) IsVoid() bool {
	return t.Kind == ASTTypeKindVoid
}
