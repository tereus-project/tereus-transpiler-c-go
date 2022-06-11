package ast

import (
	"fmt"
	"strings"
)

type AstStructInitialization struct {
	Struct *ASTType
	Fields []IASTExpression
}

func NewAstStructInitialization(typ *ASTType, fields []IASTExpression) *AstStructInitialization {
	return &AstStructInitialization{
		Struct: typ,
		Fields: fields,
	}
}

func (e *AstStructInitialization) String() string {
	fields := make([]string, len(e.Fields))

	for i, field := range e.Fields {
		fields[i] = field.String()
	}

	return fmt.Sprintf("%s{%s}", e.Struct.Name, strings.Join(fields, ", "))
}

func (e *AstStructInitialization) GetType() *ASTType {
	return e.Struct
}
