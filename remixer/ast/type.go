package ast

import "fmt"

type ASTTypeKind int64

const (
	ASTTypeKindVoid ASTTypeKind = iota
	ASTTypeKindInt
	ASTTypeKindInt16
	ASTTypeKindInt64
	ASTTypeKindChar
	ASTTypeKindFloat32
	ASTTypeKindFloat64

	ASTTypeKindArray
	ASTTypeKindPointer

	ASTTypeKindStruct

	ASTTypeKindFunction
)

type ASTType struct {
	Kind ASTTypeKind
	Name string

	ArrayType    *ASTType
	PointerType  *ASTType
	FunctionType *ASTFunction
}

func NewASTType(kind ASTTypeKind, name string) *ASTType {
	return &ASTType{
		Kind: kind,
		Name: name,
	}
}

func (t *ASTType) String() string {
	switch t.Kind {
	case ASTTypeKindArray:
		return fmt.Sprintf("[]%s", t.ArrayType.String())
	case ASTTypeKindPointer:
		return fmt.Sprintf("*%s", t.PointerType.String())
	case ASTTypeKindFunction:
		return t.FunctionType.TypeString(false)
	default:
		return t.Name
	}
}

func (t *ASTType) IsVoid() bool {
	return t.Kind == ASTTypeKindVoid
}

func (t *ASTType) IsFunction() bool {
	return t.Kind == ASTTypeKindFunction
}
