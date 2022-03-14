package ast

import "fmt"

type ASTExpressionParenthesized struct {
	Expression IASTExpression
}

func NewAstExpressionParenthesized(expression IASTExpression) *ASTExpressionParenthesized {
	return &ASTExpressionParenthesized{
		Expression: expression,
	}
}

func (e *ASTExpressionParenthesized) String() string {
	return fmt.Sprintf("(%s)", e.Expression.String())
}

func (e *ASTExpressionParenthesized) GetType() *ASTType {
	return e.Expression.GetType()
}
