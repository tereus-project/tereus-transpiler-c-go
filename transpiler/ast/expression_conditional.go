package ast

func EnsureConditionalValidity(expression IASTExpression) IASTExpression {
	if expression.GetType().IsPointer() {
		return NewASTExpressionBinary(expression, "!=", NewASTExpressionLiteral("nil", expression.GetType()), NewASTType(ASTTypeKindBool, "bool"))
	}

	if expression.GetType().IsInteger() {
		return NewASTExpressionBinary(expression, ">", NewASTExpressionLiteral("0", expression.GetType()), NewASTType(ASTTypeKindBool, "bool"))
	}

	if expression.GetType().IsFloat() {
		return NewASTExpressionBinary(expression, ">", NewASTExpressionLiteral("0.0", expression.GetType()), NewASTType(ASTTypeKindBool, "bool"))
	}

	return expression
}
