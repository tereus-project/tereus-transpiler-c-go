package ast

import "fmt"

type ASTFunctionReturn struct {
	Expression IASTExpression
}

func NewASTFunctionReturn(expression IASTExpression) *ASTFunctionReturn {
	return &ASTFunctionReturn{
		Expression: expression,
	}
}

func (b *ASTFunctionReturn) String() string {
	return fmt.Sprintf("return %s", b.Expression.String())
}
