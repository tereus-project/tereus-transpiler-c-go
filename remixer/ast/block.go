package ast

import "strings"

type ASTBlock struct {
	Statements []IASTItem
}

func NewASTBlock(statements []IASTItem) *ASTBlock {
	return &ASTBlock{
		Statements: statements,
	}
}

func (b *ASTBlock) String() string {
	statements := make([]string, len(b.Statements))

	for i, statement := range b.Statements {
		statements[i] = "\t" + statement.String()
	}

	return "{\n" + strings.Join(statements, "\n") + "\n}"
}
