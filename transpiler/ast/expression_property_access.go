package ast

import (
	"fmt"
)

type ASTExpressionPropertyAccess struct {
	Left     IASTExpression
	Property *ASTStructProperty
}

func NewASTExpressionPropertyAccess(left IASTExpression, property *ASTStructProperty) *ASTExpressionPropertyAccess {
	return &ASTExpressionPropertyAccess{
		Left:     left,
		Property: property,
	}
}

func (e *ASTExpressionPropertyAccess) String() string {
	return fmt.Sprintf("%s.%s", e.Left.String(), e.Property.GetTranslatedName())
}

func (e *ASTExpressionPropertyAccess) GetType() *ASTType {
	return e.Property.Type
}
