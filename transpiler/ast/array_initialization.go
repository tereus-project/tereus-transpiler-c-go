package ast

import (
	"fmt"
	"strings"
)

type AstArrayInitialization struct {
	Type   *ASTType
	Fields []IASTExpression
}

func NewAstArrayInitialization(typ *ASTType, fields []IASTExpression) *AstArrayInitialization {
	return &AstArrayInitialization{
		Type:   typ,
		Fields: fields,
	}
}

func (e *AstArrayInitialization) String() string {
	fields := make([]string, len(e.Fields))

	for i, field := range e.Fields {
		fields[i] = field.String()
	}

	return fmt.Sprintf("%s{%s}", e.Type.String(), strings.Join(fields, ", "))
}

func (e *AstArrayInitialization) GetType() *ASTType {
	return e.Type
}
