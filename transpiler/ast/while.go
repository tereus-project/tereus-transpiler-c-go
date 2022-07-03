package ast

import "fmt"

type ASTWhile struct {
	Cond IASTExpression
	Then IASTItem
}

func NewASTWhile(cond IASTExpression, then IASTItem) *ASTWhile {
	return &ASTWhile{
		Cond: cond,
		Then: then,
	}
}

func (w *ASTWhile) String() string {
	then := ""

	if _, ok := w.Then.(*ASTBlock); ok {
		then = w.Then.String()
	} else {
		then = NewASTBlock([]IASTItem{w.Then}).String()
	}

	return fmt.Sprintf("for %s %s", EnsureConditionalValidity(w.Cond).String(), then)
}
