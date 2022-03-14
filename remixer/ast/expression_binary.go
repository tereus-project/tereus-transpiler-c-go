package ast

import "fmt"

type ASTExpressionBinary struct {
	Left     IASTExpression
	Operator string
	Right    IASTExpression
}

func NewASTExpressionBinary(left IASTExpression, operator string, right IASTExpression) *ASTExpressionBinary {
	return &ASTExpressionBinary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (e *ASTExpressionBinary) String() string {
	return fmt.Sprintf("%s %s %s", e.Left.String(), e.Operator, e.Right.String())
}

func (e *ASTExpressionBinary) GetType() *ASTType {
	return e.Left.GetType()
}
