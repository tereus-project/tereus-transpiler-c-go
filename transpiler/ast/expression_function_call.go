package ast

import (
	"fmt"
	"strings"
)

type ASTExpressionFunctionCall struct {
	Left IASTExpression
	Args []IASTExpression
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

func (e *ASTExpressionFunctionCall) GetType() *ASTType {
	return e.Left.GetType().FunctionType.ReturnType
}
