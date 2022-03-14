package ast

import "fmt"

type ASTExpressionUnaryPre struct {
	Operator string
	Operand  IASTExpression
}

func NewASTExpressionUnaryPre(operator string, operand IASTExpression) *ASTExpressionUnaryPre {
	return &ASTExpressionUnaryPre{
		Operator: operator,
		Operand:  operand,
	}
}

func (e *ASTExpressionUnaryPre) String() string {
	return fmt.Sprintf("%s%s", e.Operator, e.Operand.String())
}

func (e *ASTExpressionUnaryPre) GetType() *ASTType {
	return e.Operand.GetType()
}
