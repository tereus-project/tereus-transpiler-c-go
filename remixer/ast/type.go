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

	ASTTypeKindBool

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

func (t *ASTType) IsPointer() bool {
	return t.Kind == ASTTypeKindPointer
}

func (t *ASTType) IsStruct() bool {
	return t.Kind == ASTTypeKindStruct
}

func (t *ASTType) IsInteger() bool {
	return t.Kind == ASTTypeKindInt ||
		t.Kind == ASTTypeKindInt16 ||
		t.Kind == ASTTypeKindInt64 ||
		t.Kind == ASTTypeKindChar
}

func (t *ASTType) IsFloat() bool {
	return t.Kind == ASTTypeKindFloat32 ||
		t.Kind == ASTTypeKindFloat64
}

func (t *ASTType) IsConvertibleTo(targetType *ASTType) bool {
	if t.IsSameTo(targetType) {
		return true
	}

	if targetType.IsInteger() || targetType.IsFloat() {
		return t.IsInteger() || t.IsFloat()
	}

	if targetType.IsPointer() {
		if !t.IsPointer() {
			return false
		}

		return t.PointerType.IsConvertibleTo(targetType.PointerType)
	}

	return false
}

func (t *ASTType) IsSameTo(other *ASTType) bool {
	if t.Kind == ASTTypeKindPointer && other.Kind == ASTTypeKindArray {
		return t.PointerType.IsSameTo(other.ArrayType)
	}

	if t.Kind == ASTTypeKindArray && other.Kind == ASTTypeKindPointer {
		return t.ArrayType.IsSameTo(other.PointerType)
	}

	if t.Kind != other.Kind {
		return false
	}

	if t.Kind == ASTTypeKindPointer {
		return t.PointerType.IsSameTo(other.PointerType)
	}

	if t.Kind == ASTTypeKindArray {
		return t.ArrayType.IsSameTo(other.ArrayType)
	}

	if t.Kind == ASTTypeKindStruct {
		return t.StructType.IsSameTo(other.StructType)
	}

	return t.Name == other.Name
}
