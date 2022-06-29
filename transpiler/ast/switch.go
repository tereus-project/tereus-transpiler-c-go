package ast

import (
	"fmt"

	"github.com/tereus-project/tereus-transpiler-c-go/transpiler/utils"
)

type ASTSwitch struct {
	Expression  IASTExpression
	Cases       []*ASTCase
	DefaultCase *ASTDefault
}

func NewASTSwitch(expression IASTExpression) *ASTSwitch {
	return &ASTSwitch{
		Expression: expression,
	}
}

func (s *ASTSwitch) AddCase(case_ *ASTCase) {
	s.Cases = append(s.Cases, case_)
}

func (s *ASTSwitch) SetDefaultCase(defaultCase *ASTDefault) {
	s.DefaultCase = defaultCase
}

func (i *ASTSwitch) String() string {
	var output string

	output += fmt.Sprintf("switch %s {\n", i.Expression.String())

	stackedCases := []*ASTCase{}

	for _, c := range i.Cases {
		if c.Block.IsEmpty() {
			stackedCases = append(stackedCases, c)
			continue
		}

		if len(stackedCases) > 0 {
			output += utils.Indent(c.StringJoined(stackedCases))
			stackedCases = nil
		} else {
			output += utils.Indent(c.String())
		}

		output += "\n"
	}

	if i.DefaultCase != nil {
		output += utils.Indent(i.DefaultCase.String())
	}

	output += "\n}"

	return output
}
