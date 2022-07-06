package ast

import (
	"fmt"
)

type ASTExpressionEqualityBinary struct {
	Left     IASTExpression
	Operator string
	Right    IASTExpression
}

func NewASTExpressionEqualityBinary(left IASTExpression, operator string, right IASTExpression) *ASTExpressionEqualityBinary {
	return &ASTExpressionEqualityBinary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (e *ASTExpressionEqualityBinary) String() string {
	return fmt.Sprintf("%s %s %s", e.Left, e.Operator, e.Right)
}

func (e *ASTExpressionEqualityBinary) GetType() *ASTType {
	return NewASTType(ASTTypeKindBool, "bool")
}
