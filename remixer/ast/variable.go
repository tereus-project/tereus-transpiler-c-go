package ast

import "fmt"

type ASTVariableDeclaration struct {
	Name       string
	Type       *ASTType
	Expression IASTExpression
}

func NewASTVariableDeclaration(name string, typ *ASTType) *ASTVariableDeclaration {
	return &ASTVariableDeclaration{
		Name: name,
		Type: typ,
	}
}

func (v *ASTVariableDeclaration) String() string {
	if v.Expression != nil {
		return fmt.Sprintf("%s := %s", v.Name, v.Expression.String())
	} else {
		return fmt.Sprintf("var %s %s", v.Name, v.Type.String())
	}
}
