package ast

import (
	"fmt"
)

type ASTExpressionConditionalBinary struct {
	Left     IASTExpression
	Operator string
	Right    IASTExpression
}

func NewASTExpressionConditionalBinary(left IASTExpression, operator string, right IASTExpression) *ASTExpressionConditionalBinary {
	return &ASTExpressionConditionalBinary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (e *ASTExpressionConditionalBinary) String() string {
	return fmt.Sprintf("%s %s %s", EnsureConditionalValidity(e.Left).String(), e.Operator, EnsureConditionalValidity(e.Right).String())
}

func (e *ASTExpressionConditionalBinary) GetType() *ASTType {
	return NewASTType(ASTTypeKindBool, "bool")
}
