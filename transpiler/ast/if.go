package ast

import "fmt"

type ASTIf struct {
	Condition IASTExpression
	Then      IASTItem
	Else      IASTItem
}

func NewASTIf(condition IASTExpression, then IASTItem) *ASTIf {
	return &ASTIf{
		Condition: condition,
		Then:      then,
	}
}

func (i *ASTIf) String() string {
	then := ""

	if _, ok := i.Then.(*ASTBlock); ok {
		then = i.Then.String()
	} else {
		then = NewASTBlock([]IASTItem{i.Then}).String()
	}

	if i.Else != nil {
		else_ := ""

		if block, ok := i.Else.(*ASTBlock); ok {
			else_ = block.String()
		} else if block, ok := i.Else.(*ASTIf); ok {
			else_ = block.String()
		} else {
			else_ = NewASTBlock([]IASTItem{i.Else}).String()
		}

		return fmt.Sprintf("if %s %s else %s", i.Condition.String(), then, else_)
	}

	return fmt.Sprintf("if %s %s", i.Condition.String(), then)
}
