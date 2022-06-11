package ast

import "fmt"

type ASTExpressionCast struct {
	Expression IASTExpression
	Type       *ASTType
}

func NewASTExpressionCast(expression IASTExpression, typ *ASTType) *ASTExpressionCast {
	return &ASTExpressionCast{
		Expression: expression,
		Type:       typ,
	}
}

func (e *ASTExpressionCast) String() string {
	switch e.Type.Kind {
	case ASTTypeKindInt16, ASTTypeKindInt, ASTTypeKindInt64:
		return fmt.Sprintf("%s(%s)", e.Type.String(), e.Expression.String())
	default:
		return fmt.Sprintf("(%s)(%s)", e.Type.String(), e.Expression.String())
	}
}

func (e *ASTExpressionCast) GetType() *ASTType {
	return e.Type
}
