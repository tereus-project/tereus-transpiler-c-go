package ast

import "fmt"

type ASTExpressionUnaryPre struct {
	Operator string
	Operand  IASTExpression

	IsStatement bool
}

func NewASTExpressionUnaryPre(operator string, operand IASTExpression, isStatement bool) *ASTExpressionUnaryPre {
	return &ASTExpressionUnaryPre{
		Operator: operator,
		Operand:  operand,

		IsStatement: isStatement,
	}
}

func (e *ASTExpressionUnaryPre) String() string {
	if e.Operator == "++" || e.Operator == "--" {
		if !e.IsStatement {
			return fmt.Sprintf("%s%s", e.Operand.String(), e.Operator)
		}

		if e.Operator == "++" {
			switch e.Operand.GetType().Kind {
			case ASTTypeKindPointer:
				return fmt.Sprintf("libc.PreIncPtr(&%s)", e.Operand.String())
			default:
				return fmt.Sprintf("libc.PreInc(&%s)", e.Operand.String())
			}
		} else if e.Operator == "--" {
			switch e.Operand.GetType().Kind {
			case ASTTypeKindPointer:
				return fmt.Sprintf("libc.PreDecPtr(&%s)", e.Operand.String())
			default:
				return fmt.Sprintf("libc.PreDec(&%s)", e.Operand.String())
			}
		}
	}

	return fmt.Sprintf("%s%s", e.Operator, e.Operand.String())
}

func (e *ASTExpressionUnaryPre) GetType() *ASTType {
	return e.Operand.GetType()
}
