package ast

import "fmt"

type ASTTypedef struct {
	Name string
	Type *ASTType
}

func NewASTTypedef(name string, typ *ASTType) *ASTTypedef {
	return &ASTTypedef{
		Name: name,
		Type: typ,
	}
}

func (t *ASTTypedef) String() string {
	return fmt.Sprintf("type %s %s", t.Name, t.Type.String())
}
