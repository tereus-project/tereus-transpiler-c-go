package ast

import (
	"fmt"
)

type ASTExpressionPointerPropertyAccess struct {
	Left     IASTExpression
	Property *ASTStructProperty
}

func NewASTExpressionPointerPropertyAccess(left IASTExpression, property *ASTStructProperty) *ASTExpressionPointerPropertyAccess {
	return &ASTExpressionPointerPropertyAccess{
		Left:     left,
		Property: property,
	}
}

func (e *ASTExpressionPointerPropertyAccess) String() string {
	return fmt.Sprintf("%s.%s", e.Left.String(), e.Property.GetTranslatedName())
}

func (e *ASTExpressionPointerPropertyAccess) GetType() *ASTType {
	return e.Property.Type
}
