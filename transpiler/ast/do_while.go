package ast

import (
	"fmt"

	"github.com/tereus-project/tereus-transpiler-c-go/transpiler/utils"
)

type ASTDoWhile struct {
	Cond IASTExpression
	Then IASTItem
}

func NewASTDoWhile(cond IASTExpression, then IASTItem) *ASTDoWhile {
	return &ASTDoWhile{
		Cond: cond,
		Then: then,
	}
}

func (w *ASTDoWhile) String() string {
	statements := ""

	if block, ok := w.Then.(*ASTBlock); ok {
		statements = block.StringInner()
	} else {
		statements = w.Then.String()
	}

	condition := NewASTIf(NewASTExpressionUnaryPre("!", EnsureConditionalValidity(NewAstExpressionParenthesized(w.Cond)), false), NewASTBreak())

	return fmt.Sprintf("for {\n%s\n\n%s\n}", utils.Indent(statements), condition.String())
}
