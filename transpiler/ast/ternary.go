package ast

import "fmt"

type ASTExpressionTernary struct {
	Cond        IASTExpression
	True        IASTExpression
	False       IASTExpression
	IsStatement bool
}

func NewASTExpressionTernary(cond IASTExpression, true_ IASTExpression, false_ IASTExpression, isStatement bool) *ASTExpressionTernary {
	return &ASTExpressionTernary{
		Cond:        cond,
		True:        true_,
		False:       false_,
		IsStatement: isStatement,
	}
}

func (e *ASTExpressionTernary) String() string {
	return fmt.Sprintf(
		"libc.Ternary(%s, func() %s { return %s }, func() %s { return %s })",
		e.Cond.String(),
		e.True.GetType().String(),
		e.True.String(),
		e.False.GetType().String(),
		e.False.String(),
	)
}

func (e *ASTExpressionTernary) GetType() *ASTType {
	return e.True.GetType()
}
