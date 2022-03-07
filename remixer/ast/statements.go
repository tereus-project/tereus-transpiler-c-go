package ast

import "strings"

type ASTMultipleStatements struct {
	Statements []IASTItem
}

func NewASTMultipleStatements(statements []IASTItem) *ASTMultipleStatements {
	return &ASTMultipleStatements{
		Statements: statements,
	}
}

func (m *ASTMultipleStatements) String() string {
	statements := make([]string, len(m.Statements))

	for i, statement := range m.Statements {
		statements[i] = statement.String()
	}

	return strings.Join(statements, "\n")
}
