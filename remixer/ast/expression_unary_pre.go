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
	if e.Operator == "++" || e.Operator == "--" {
		switch e.Operand.GetType().Kind {
		case ASTTypeKindPointer:
			return fmt.Sprintf("libc.PreIncPtr(&%s)", e.Operand.String())
		default:
			return fmt.Sprintf("libc.PreInc(&%s)", e.Operand.String())
		}
	} else {
		return fmt.Sprintf("%s%s", e.Operand.String(), e.Operator)
	}
}

func (e *ASTExpressionUnaryPre) GetType() *ASTType {
	return e.Operand.GetType()
}
