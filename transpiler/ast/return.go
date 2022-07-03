package ast

import "fmt"

type ASTFunctionReturn struct {
	Expression IASTExpression
	IsInMain   bool
}

func NewASTFunctionReturn(isInMain bool) *ASTFunctionReturn {
	return &ASTFunctionReturn{
		IsInMain: isInMain,
	}
}

func (b *ASTFunctionReturn) SetExpression(expression IASTExpression) {
	b.Expression = expression
}

func (b *ASTFunctionReturn) String() string {
	if b.Expression != nil {
		if b.IsInMain {
			return fmt.Sprintf("os.Exit(%s)", b.Expression.String())
		}

		return fmt.Sprintf("return %s", b.Expression.String())
	}

	return "return"
}
