package ast

import "fmt"

type ASTExpressionIndex struct {
	Left  IASTExpression
	Index IASTExpression
}

func NewASTExpressionIndex(left IASTExpression, index IASTExpression) *ASTExpressionIndex {
	return &ASTExpressionIndex{
		Left:  left,
		Index: index,
	}
}

func (e *ASTExpressionIndex) String() string {
	switch e.Left.GetType().Kind {
	case ASTTypeKindPointer:
		sizeof := fmt.Sprintf("unsafe.Sizeof(%s)", getDefaultExpression(e.Left.GetType()))
		return fmt.Sprintf("*(%s)(unsafe.Add(unsafe.Pointer(%s), uintptr(%s) * %s))", e.Left.GetType().String(), e.Left.String(), e.Index.String(), sizeof)
	default:
		return fmt.Sprintf("%s[%s]", e.Left.String(), e.Index.String())
	}
}

func (e *ASTExpressionIndex) GetType() *ASTType {
	typ := e.Left.GetType()

	switch typ.Kind {
	case ASTTypeKindArray:
		return typ.ArrayType
	case ASTTypeKindPointer:
		return typ.PointerType
	}

	return nil
}
