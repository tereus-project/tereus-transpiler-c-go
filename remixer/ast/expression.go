package ast

import (
	"fmt"
	"strings"
)

type IASTExpression interface {
	IASTItem
}

type ASTExpressionBinary struct {
	Left     IASTExpression
	Operator string
	Right    IASTExpression
}

type ASTExpressionUnaryPost struct {
	Operand  IASTExpression
	Operator string
}

type ASTExpressionUnaryPre struct {
	Operator string
	Operand  IASTExpression
}

type ASTExpressionLiteral struct {
	Value string
}

type ASTExpressionFunctionCall struct {
	Left IASTExpression
	Args []IASTExpression
}

func NewASTExpressionBinary(left IASTExpression, operator string, right IASTExpression) *ASTExpressionBinary {
	return &ASTExpressionBinary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (e *ASTExpressionBinary) String() string {
	return fmt.Sprintf("%s %s %s", e.Left.String(), e.Operator, e.Right.String())
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

func NewASTExpressionUnaryPre(operator string, operand IASTExpression) *ASTExpressionUnaryPre {
	return &ASTExpressionUnaryPre{
		Operator: operator,
		Operand:  operand,
	}
}

func (e *ASTExpressionUnaryPre) String() string {
	return fmt.Sprintf("%s%s", e.Operator, e.Operand.String())
}

func NewASTExpressionLiteral(value string) *ASTExpressionLiteral {
	return &ASTExpressionLiteral{
		Value: value,
	}
}

func (e *ASTExpressionLiteral) String() string {
	return e.Value
}

func NewASTExpressionFunctionCall(left IASTExpression, args []IASTExpression) *ASTExpressionFunctionCall {
	return &ASTExpressionFunctionCall{
		Left: left,
		Args: args,
	}
}

func (e *ASTExpressionFunctionCall) String() string {
	args := make([]string, len(e.Args))

	for i, arg := range e.Args {
		args[i] = arg.String()
	}

	return fmt.Sprintf("%s(%s)", e.Left.String(), strings.Join(args, ", "))
}
