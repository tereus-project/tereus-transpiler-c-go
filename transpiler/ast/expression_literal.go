package ast

import "strings"

type ASTExpressionLiteral struct {
	Value string
	Type  *ASTType
}

func NewASTExpressionLiteral(value string, typ_ *ASTType) *ASTExpressionLiteral {
	return &ASTExpressionLiteral{
		Value: value,
		Type:  typ_,
	}
}

func (e *ASTExpressionLiteral) String() string {
	if e.Type.Kind == ASTTypeKindChar || (e.Type.IsArray() && e.Type.ArrayType.Kind == ASTTypeKindChar) {
		return strings.ReplaceAll(e.Value, "\\0", "\\000")
	}

	return e.Value
}

func (e *ASTExpressionLiteral) GetType() *ASTType {
	return e.Type
}
