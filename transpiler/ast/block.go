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
		lines := strings.Split(statement.String(), "\n")

		for j, line := range lines {
			lines[j] = "\t" + line
		}

		statements[i] = strings.Join(lines, "\n")
	}

	return "{\n" + strings.Join(statements, "\n") + "\n}"
}
