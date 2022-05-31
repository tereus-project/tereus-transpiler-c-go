package ast

import "fmt"

type ASTTypeKind int64

const (
	ASTTypeKindAny ASTTypeKind = iota

	ASTTypeKindVoid
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
	StructType   *ASTStruct
}

func NewASTType(kind ASTTypeKind, name string) *ASTType {
	return &ASTType{
		Kind: kind,
		Name: name,
	}
}

func (t *ASTType) SetPointerType(pointerType *ASTType) *ASTType {
	t.PointerType = pointerType
	return t
}

func (t *ASTType) SetArrayType(arrayType *ASTType) *ASTType {
	t.ArrayType = arrayType
	return t
}

func (t *ASTType) SetFunctionType(functionType *ASTFunction) *ASTType {
	t.FunctionType = functionType
	return t
}

func (t *ASTType) SetStructType(structType *ASTStruct) *ASTType {
	t.StructType = structType
	return t
}

func (t *ASTType) String() string {
	switch t.Kind {
	case ASTTypeKindArray:
		return fmt.Sprintf("[]%s", t.ArrayType.String())
	case ASTTypeKindPointer:
		return fmt.Sprintf("*%s", t.PointerType.String())
	case ASTTypeKindFunction:
		return t.FunctionType.TypeString(false)
	case ASTTypeKindStruct:
		return t.StructType.TypeString()
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
