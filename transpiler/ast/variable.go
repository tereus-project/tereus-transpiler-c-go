package ast

import (
	"fmt"
	"strings"
)

type ASTVariableDeclaration struct {
	IsStatement          bool
	ForceFullDeclaration bool
	Items                []*ASTVariableDeclarationItem
}

type ASTVariableDeclarationItem struct {
	Name       string
	Type       *ASTType
	Expression IASTExpression
}

func NewASTVariableDeclaration(isStatement bool) *ASTVariableDeclaration {
	return &ASTVariableDeclaration{
		IsStatement: isStatement,
	}
}

func (v *ASTVariableDeclaration) SetForceFullDeclaration(forceFullDeclaration bool) *ASTVariableDeclaration {
	v.ForceFullDeclaration = forceFullDeclaration
	return v
}

func (v *ASTVariableDeclaration) String() string {
	if v.IsStatement {
		return v.StringStatement()
	}

	names := make([]string, 0)
	expressions := make([]string, 0)

	for _, item := range v.Items {
		names = append(names, item.Name)

		if item.Expression != nil {
			expressions = append(expressions, item.Expression.String())
		} else {
			expressions = append(expressions, v.getDefaultExpression(item))
		}
	}

	return fmt.Sprintf("%s := %s", strings.Join(names, ", "), strings.Join(expressions, ", "))
}

func (v *ASTVariableDeclaration) StringStatement() string {
	declarations := make([]string, len(v.Items))

	for i, item := range v.Items {
		var declaration string

		if item.Expression != nil {
			if v.ForceFullDeclaration {
				declaration = fmt.Sprintf("var %s %s = %s", item.Name, item.Type.String(), item.Expression.String())
			} else {
				declaration = fmt.Sprintf("%s := %s", item.Name, item.Expression.String())
			}
		} else {
			declaration = fmt.Sprintf("var %s %s", item.Name, item.Type.String())
		}

		declarations[i] = declaration
	}

	return strings.Join(declarations, "\n")
}

func (v *ASTVariableDeclaration) getDefaultExpression(item *ASTVariableDeclarationItem) string {
	switch item.Type.Kind {
	case ASTTypeKindArray:
		return fmt.Sprintf("%s{}", item.Type.String())
	case ASTTypeKindPointer:
		return fmt.Sprintf("(%s)(nil)", item.Type.String())
	case ASTTypeKindStruct:
		return fmt.Sprintf("%s{}", item.Type.Name)
	case ASTTypeKindFloat32:
		return "float32(0)"
	case ASTTypeKindFloat64:
		return "float64(0)"
	case ASTTypeKindInt:
		return "0"
	case ASTTypeKindInt16:
		return "int16(0)"
	case ASTTypeKindInt64:
		return "int64(0)"
	case ASTTypeKindChar:
		return "byte(0)"
	default:
		return "unknown"
	}
}

func NewASTVariableDeclarationItem(name string, typ *ASTType, expression IASTExpression) *ASTVariableDeclarationItem {
	return &ASTVariableDeclarationItem{
		Name:       name,
		Type:       typ,
		Expression: expression,
	}
}
