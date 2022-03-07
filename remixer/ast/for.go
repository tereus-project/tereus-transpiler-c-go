package ast

import "fmt"

type ASTFor struct {
	Init IASTItem
	Cond IASTExpression
	Post IASTExpression

	Statement IASTItem
}

func NewASTFor() *ASTFor {
	return &ASTFor{}
}

func (f *ASTFor) String() string {
	init := ""
	if f.Init != nil {
		init = f.Init.String()
	}

	cond := ""
	if f.Cond != nil {
		cond = f.Cond.String()
	}

	post := ""
	if f.Post != nil {
		post = f.Post.String()
	}

	return fmt.Sprintf("for %s; %s; %s %s", init, cond, post, f.Statement.String())
}
