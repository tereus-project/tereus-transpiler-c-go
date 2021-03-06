package scope

import "github.com/tereus-project/tereus-transpiler-c-go/transpiler/ast"

type ScopeVariable struct {
	*ScopeItem
	Type *ast.ASTType
}

func NewScopeVariable(name string, translatedName string, typ *ast.ASTType) *ScopeVariable {
	return &ScopeVariable{
		ScopeItem: NewScopeItem(name, translatedName),
		Type:      typ,
	}
}

func (s *ScopeVariable) GetType() *ast.ASTType {
	return s.Type
}
