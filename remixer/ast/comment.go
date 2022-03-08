package ast

type ASTComment struct {
	MultiLine bool
	Value     string
}

func NewASTComment(multiLine bool, value string) *ASTComment {
	return &ASTComment{
		MultiLine: multiLine,
		Value:     value,
	}
}

func (c *ASTComment) String() string {
	return c.Value
}
