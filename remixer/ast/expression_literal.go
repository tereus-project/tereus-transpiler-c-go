package ast

type ASTExpressionLiteral struct {
	Value string
	Type  *ASTType
}

func NewASTExpressionLiteral(value string, typ_ *ASTType) *ASTExpressionLiteral {
	return &ASTExpressionLiteral{
		Value: value,
		Type:  typ_,
	}
}

func (e *ASTExpressionLiteral) String() string {
	return e.Value
}

func (e *ASTExpressionLiteral) GetType() *ASTType {
	return e.Type
}
