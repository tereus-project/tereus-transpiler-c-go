package ast

import "fmt"

type ASTExpressionUnaryPost struct {
	Operand  IASTExpression
	Operator string

	IsStatement bool
}

func NewASTExpressionUnaryPost(operand IASTExpression, operator string, isStatement bool) *ASTExpressionUnaryPost {
	return &ASTExpressionUnaryPost{
		Operand:  operand,
		Operator: operator,

		IsStatement: isStatement,
	}
}

func (e *ASTExpressionUnaryPost) String() string {
	if !e.IsStatement {
		if e.Operator == "++" {
			switch e.Operand.GetType().Kind {
			case ASTTypeKindPointer:
				return fmt.Sprintf("libc.PostIncPtr(&%s)", e.Operand.String())
			default:
				return fmt.Sprintf("libc.PostInc(&%s)", e.Operand.String())
			}
		} else if e.Operator == "--" {
			switch e.Operand.GetType().Kind {
			case ASTTypeKindPointer:
				return fmt.Sprintf("libc.PostDecPtr(&%s)", e.Operand.String())
			default:
				return fmt.Sprintf("libc.PostDec(&%s)", e.Operand.String())
			}
		}
	}

	return fmt.Sprintf("%s%s", e.Operand.String(), e.Operator)
}

func (e *ASTExpressionUnaryPost) GetType() *ASTType {
	return e.Operand.GetType()
}
