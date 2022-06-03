package ast

import "fmt"

type ASTTypeConversion struct {
	Expression IASTExpression
	TargetType *ASTType
}

func NewAstTypeConversion(expression IASTExpression, targetType *ASTType) (*ASTTypeConversion, error) {
	if !expression.GetType().IsConvertibleTo(targetType) {
		return nil, fmt.Errorf("Can't implicitly convert '%s' to '%s'", expression.GetType().String(), targetType.String())
	}

	return &ASTTypeConversion{
		Expression: expression,
		TargetType: targetType,
	}, nil
}

func (t *ASTTypeConversion) GetType() *ASTType {
	return t.TargetType
}

func (t *ASTTypeConversion) String() string {
	if t.Expression.GetType().IsSameTo(t.TargetType) {
		return t.Expression.String()
	}

	if t.TargetType.IsInteger() || t.TargetType.IsFloat() {
		return fmt.Sprintf("%s(%s)", t.TargetType.String(), t.Expression.String())
	}

	return fmt.Sprintf("(%s)(%s)", t.TargetType.String(), t.Expression.String())
}
