package ast

import "fmt"

type ASTParenthesizedExpression struct {
	Expression IASTExpression
}

func NewAstParenthesizedExpression(expression IASTExpression) *ASTParenthesizedExpression {
	return &ASTParenthesizedExpression{
		Expression: expression,
	}
}

func (e *ASTParenthesizedExpression) String() string {
	return fmt.Sprintf("(%s)", e.Expression.String())
}
