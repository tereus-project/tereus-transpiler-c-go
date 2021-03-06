package scope

import "github.com/tereus-project/tereus-transpiler-c-go/transpiler/ast"

type ScopeType struct {
	*ScopeItem
	Type *ast.ASTType
}

func NewScopeType(name string, translatedName string, typ *ast.ASTType) *ScopeType {
	return &ScopeType{
		ScopeItem: NewScopeItem(name, translatedName),
		Type:      typ,
	}
}

func (s *ScopeType) GetType() *ast.ASTType {
	return s.Type
}
