package ast

import "fmt"

type ASTGoto struct {
	Label string
}

func NewASTGoto(label string) *ASTGoto {
	return &ASTGoto{
		Label: label,
	}
}

func (g *ASTGoto) String() string {
	return fmt.Sprintf("goto %s", g.Label)
}
