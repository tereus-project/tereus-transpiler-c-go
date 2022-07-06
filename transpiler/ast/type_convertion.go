package ast

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type ASTTypeConversion struct {
	Expression IASTExpression
	TargetType *ASTType
}

func NewAstTypeConversion(expression IASTExpression, targetType *ASTType) (*ASTTypeConversion, error) {
	if !expression.GetType().IsConvertibleTo(targetType) {
		logrus.WithFields(logrus.Fields{
			"expression":     expression,
			"expressionType": expression.GetType(),
			"targetType":     targetType,
		}).Error("Expression is not convertible to target type")
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
	originalType := t.Expression.GetType()

	if originalType.IsSameTo(t.TargetType) {
		return t.Expression.String()
	}

	if originalType.IsPointer() && t.TargetType.IsInteger() {
		return fmt.Sprintf("%s(unsafe.Pointer(%s))", t.TargetType.String(), t.Expression.String())
	}

	if t.TargetType.IsInteger() || t.TargetType.IsFloat() || t.TargetType.Kind == ASTTypeKindByte {
		return fmt.Sprintf("%s(%s)", t.TargetType.String(), t.Expression.String())
	}

	return fmt.Sprintf("(%s)(%s)", t.TargetType.String(), t.Expression.String())
}
