package ast

import "fmt"

type ASTTypeKind int64

const (
	ASTTypeKindVoid ASTTypeKind = iota
	ASTTypeKindInt
	ASTTypeKindInt16
	ASTTypeKindInt64
	ASTTypeKindRune
	ASTTypeKindFloat32
	ASTTypeKindFloat64

	ASTTypeKindArray
	ASTTypeKindPointer

	ASTTypeKindStruct
)

type ASTType struct {
	Kind ASTTypeKind
	Name string

	ArrayType   *ASTType
	PointerType *ASTType
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
	default:
		return t.Name
	}
}

func (t *ASTType) IsVoid() bool {
	return t.Kind == ASTTypeKindVoid
}
