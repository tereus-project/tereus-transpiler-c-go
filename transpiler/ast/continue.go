package ast

type ASTContinue struct{}

func NewASTContinue() *ASTContinue {
	return &ASTContinue{}
}

func (b *ASTContinue) String() string {
	return "continue"
}
