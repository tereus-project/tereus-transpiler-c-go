package ast

import (
	"fmt"
	"strings"

	"github.com/tereus-project/tereus-transpiler-c-go/transpiler/utils"
)

type ASTBlock struct {
	Statements []IASTItem
}

func NewASTBlock(statements []IASTItem) *ASTBlock {
	return &ASTBlock{
		Statements: statements,
	}
}

func (b *ASTBlock) String() string {
	inner := b.StringInner()
	return fmt.Sprintf("{\n%s\n}", utils.Indent(inner))
}

func (b *ASTBlock) StringInner() string {
	statements := make([]string, len(b.Statements))

	for i, statement := range b.Statements {
		statements[i] = statement.String()
	}

	return strings.Join(statements, "\n")
}

func (b *ASTBlock) IsEmpty() bool {
	return len(b.Statements) == 0
}

func blockEndsWith[T any](block *ASTBlock) bool {
	if block.IsEmpty() {
		return false
	}

	_, ok := block.Statements[len(block.Statements)-1].(T)
	return ok
}

func (b *ASTBlock) EndsWithBreak() bool {
	return blockEndsWith[*ASTBreak](b)
}

func (b *ASTBlock) EndsWithReturn() bool {
	return blockEndsWith[*ASTFunctionReturn](b)
}
