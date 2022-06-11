package ast

import (
	"fmt"
	"strings"
)

type AstStructNamedInitialization struct {
	Struct      *ASTType
	Names       []string
	Expressions []IASTExpression
}

func NewAstStructNamedInitialization(typ *ASTType, names []string, expressions []IASTExpression) *AstStructNamedInitialization {
	return &AstStructNamedInitialization{
		Struct:      typ,
		Names:       names,
		Expressions: expressions,
	}
}

func (e *AstStructNamedInitialization) String() string {
	structType := e.Struct.StructType
	fields := make([]string, len(e.Expressions))

	for i, field := range e.Expressions {
		fields[i] = fmt.Sprintf("\t%s: %s", structType.GetProperty(e.Names[i]).GetTranslatedName(), field.String())
	}

	return fmt.Sprintf("%s{\n%s,\n}", e.Struct.Name, strings.Join(fields, ",\n"))
}

func (s *AstStructNamedInitialization) GetType() *ASTType {
	return s.Struct
}
