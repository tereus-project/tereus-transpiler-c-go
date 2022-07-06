package ast

import "fmt"

type ASTExpressionBinary struct {
	Left       IASTExpression
	Operator   string
	Right      IASTExpression
	ReturnType *ASTType
}

func NewASTExpressionBinary(left IASTExpression, operator string, right IASTExpression, returnType *ASTType) *ASTExpressionBinary {
	return &ASTExpressionBinary{
		Left:       left,
		Operator:   operator,
		Right:      right,
		ReturnType: returnType,
	}
}

func (e *ASTExpressionBinary) String() string {
	leftType := e.Left.GetType()
	if leftType.IsPointer() {
		if e.Operator == "+" || e.Operator == "-" {
			sizeof := fmt.Sprintf("unsafe.Sizeof(%s)", getDefaultExpression(leftType))
			return fmt.Sprintf("unsafe.Add(unsafe.Pointer(uintptr(%s)), uintptr(%s) * %s)", e.Left.String(), e.Right.String(), sizeof)
		}
	}

	return fmt.Sprintf("%s %s %s", e.Left.String(), e.Operator, e.Right.String())
}

func (e *ASTExpressionBinary) GetType() *ASTType {
	return e.ReturnType
}
