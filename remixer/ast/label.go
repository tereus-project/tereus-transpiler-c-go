package ast

import "fmt"

type ASTLabel struct {
	Name string
}

func NewASTLabel(name string) *ASTLabel {
	return &ASTLabel{
		Name: name,
	}
}

func (l *ASTLabel) String() string {
	return fmt.Sprintf("%s:", l.Name)
}
