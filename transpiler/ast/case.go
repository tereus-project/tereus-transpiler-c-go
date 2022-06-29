package ast

import (
	"fmt"
	"strings"

	"github.com/tereus-project/tereus-transpiler-c-go/transpiler/utils"
)

type ASTCase struct {
	Expression IASTExpression
	Block      *ASTBlock
}

func NewASTCase(expression IASTExpression, block *ASTBlock) *ASTCase {
	return &ASTCase{
		Expression: expression,
		Block:      block,
	}
}

func (i *ASTCase) String() string {
	return fmt.Sprintf("case %s:\n%s", i.Expression.String(), utils.Indent(i.StringInner()))
}

func (i *ASTCase) StringJoined(cases []*ASTCase) string {
	conditions := []string{}
	for _, c := range cases {
		conditions = append(conditions, c.Expression.String())
	}

	conditions = append(conditions, i.Expression.String())

	return fmt.Sprintf("case %s:\n%s", strings.Join(conditions, ","), utils.Indent(i.StringInner()))
}

func (i *ASTCase) StringInner() string {
	if i.Block.EndsWithReturn() || !i.Block.EndsWithBreak() {
		return i.Block.StringInner()
	}

	blockCopy := NewASTBlock(nil)
	blockCopy.Statements = i.Block.Statements[:len(i.Block.Statements)-1]
	return blockCopy.StringInner()
}
