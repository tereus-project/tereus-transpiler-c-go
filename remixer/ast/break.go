package ast

type ASTBreak struct{}

func NewASTBreak() *ASTBreak {
	return &ASTBreak{}
}

func (b *ASTBreak) String() string {
	return "break"
}
