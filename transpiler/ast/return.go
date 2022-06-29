package ast

import "fmt"

type ASTFunctionReturn struct {
	Expression IASTExpression
	IsInMain   bool
}

func NewASTFunctionReturn(expression IASTExpression, isInMain bool) *ASTFunctionReturn {
	return &ASTFunctionReturn{
		Expression: expression,
		IsInMain:   isInMain,
	}
}

func (b *ASTFunctionReturn) String() string {
	if b.IsInMain {
		return fmt.Sprintf("os.Exit(%s)", b.Expression.String())
	}

	return fmt.Sprintf("return %s", b.Expression.String())
}
