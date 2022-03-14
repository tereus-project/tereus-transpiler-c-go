package ast

type IASTExpression interface {
	IASTItem

	GetType() *ASTType
}
