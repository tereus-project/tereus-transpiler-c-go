package ast

import "fmt"

type ASTExpressionUnaryPost struct {
	Operand  IASTExpression
	Operator string
}

func NewASTExpressionUnaryPost(operand IASTExpression, operator string) *ASTExpressionUnaryPost {
	return &ASTExpressionUnaryPost{
		Operand:  operand,
		Operator: operator,
	}
}

func (e *ASTExpressionUnaryPost) String() string {
	return fmt.Sprintf("%s%s", e.Operand.String(), e.Operator)
}

func (e *ASTExpressionUnaryPost) GetType() *ASTType {
	return e.Operand.GetType()
}
