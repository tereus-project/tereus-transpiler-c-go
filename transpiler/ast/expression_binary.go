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
	return fmt.Sprintf("%s %s %s", e.Left.String(), e.Operator, e.Right.String())
}

func (e *ASTExpressionBinary) GetType() *ASTType {
	return e.ReturnType
}
