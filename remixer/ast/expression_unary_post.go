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
	if e.Operator == "++" || e.Operator == "--" {
		switch e.Operand.GetType().Kind {
		case ASTTypeKindPointer:
			return fmt.Sprintf("libc.PostIncPtr(&%s)", e.Operand.String())
		default:
			return fmt.Sprintf("libc.PostInc(&%s)", e.Operand.String())
		}
	} else {
		return fmt.Sprintf("%s%s", e.Operand.String(), e.Operator)
	}
}

func (e *ASTExpressionUnaryPost) GetType() *ASTType {
	return e.Operand.GetType()
}
