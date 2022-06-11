package ast

import (
	"fmt"
	"strings"
)

type ASTEnum struct {
	Name       string
	Properties []*ASTEnumProperty
}

type ASTEnumProperty struct {
	Name       string
	Expression IASTExpression
}

func NewASTEnum(name string, properties []*ASTEnumProperty) *ASTEnum {
	return &ASTEnum{
		Name:       name,
		Properties: properties,
	}
}

func (s *ASTEnum) String() string {
	properties := make([]string, len(s.Properties))

	lastExpression := ""
	untilLastExpression := 0

	for i, property := range s.Properties {
		untilLastExpression += 1
		properties[i] = "\t" + property.String()

		if property.Expression == nil {
			if lastExpression == "" {
				if i == 0 {
					properties[i] += " = iota"
				}
			} else {
				properties[i] += fmt.Sprintf(" = %s + %d", lastExpression, untilLastExpression)
			}
		} else {
			lastExpression = property.Expression.String()
			untilLastExpression = 0
		}
	}

	return fmt.Sprintf("type %s int64\nconst (\n%s\n)", s.Name, strings.Join(properties, "\n"))
}

func NewASTEnumProperty(name string, expression IASTExpression) *ASTEnumProperty {
	return &ASTEnumProperty{
		Name:       name,
		Expression: expression,
	}
}

func (m *ASTEnumProperty) String() string {
	if expression := m.Expression; expression != nil {
		return fmt.Sprintf("%s = %s", m.Name, expression.String())
	} else {
		return m.Name
	}
}
