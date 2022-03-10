package ast

import "fmt"

type ASTExpressionCast struct {
	Expression IASTExpression
	Type       *ASTType
}

func NewASTExpressionCast(expression IASTExpression, typ *ASTType) *ASTExpressionCast {
	return &ASTExpressionCast{
		Expression: expression,
		Type:       typ,
	}
}

func (e *ASTExpressionCast) String() string {
	return fmt.Sprintf("(%s)(%s)", e.Type.String(), e.Expression.String())
}
