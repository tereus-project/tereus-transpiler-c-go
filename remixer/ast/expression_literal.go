package ast

type ASTExpressionLiteral struct {
	Value string
}

func NewASTExpressionLiteral(value string) *ASTExpressionLiteral {
	return &ASTExpressionLiteral{
		Value: value,
	}
}

func (e *ASTExpressionLiteral) String() string {
	return e.Value
}
