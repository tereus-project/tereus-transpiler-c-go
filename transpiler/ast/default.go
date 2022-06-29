package ast

import (
	"fmt"

	"github.com/tereus-project/tereus-transpiler-c-go/transpiler/utils"
)

type ASTDefault struct {
	Block *ASTBlock
}

func NewASTDefault(block *ASTBlock) *ASTDefault {
	return &ASTDefault{
		Block: block,
	}
}

func (i *ASTDefault) String() string {
	return fmt.Sprintf("default:\n%s", utils.Indent(i.Block.StringInner()))
}
