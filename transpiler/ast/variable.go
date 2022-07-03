package ast

import (
	"fmt"
	"strings"
)

type ASTVariableDeclaration struct {
	Items []*ASTVariableDeclarationItem
}

type ASTVariableDeclarationItem struct {
	Name       string
	Type       *ASTType
	Expression IASTExpression
}

func NewASTVariableDeclaration() *ASTVariableDeclaration {
	return &ASTVariableDeclaration{}
}

func (v *ASTVariableDeclaration) String() string {
	names := make([]string, 0)
	expressions := make([]string, 0)

	for _, item := range v.Items {
		names = append(names, item.Name)

		if item.Expression != nil {
			expressions = append(expressions, item.Expression.String())
		} else {
			switch item.Type.Kind {
			case ASTTypeKindArray:
				expressions = append(expressions, fmt.Sprintf("%s{}", item.Type.String()))
			case ASTTypeKindPointer:
				expressions = append(expressions, fmt.Sprintf("(%s)(nil)", item.Type.String()))
			case ASTTypeKindStruct:
				expressions = append(expressions, fmt.Sprintf("%s{}", item.Type.Name))
			case ASTTypeKindFloat32:
				expressions = append(expressions, "float32(0)")
			case ASTTypeKindFloat64:
				expressions = append(expressions, "float64(0)")
			case ASTTypeKindInt:
				expressions = append(expressions, "0")
			case ASTTypeKindInt16:
				expressions = append(expressions, "int16(0)")
			case ASTTypeKindInt64:
				expressions = append(expressions, "int64(0)")
			case ASTTypeKindChar:
				expressions = append(expressions, "byte(0)")
			default:
				expressions = append(expressions, "unknown")
			}
		}
	}

	return fmt.Sprintf("%s := %s", strings.Join(names, ", "), strings.Join(expressions, ", "))
}

func NewASTVariableDeclarationItem(name string, typ *ASTType, expression IASTExpression) *ASTVariableDeclarationItem {
	return &ASTVariableDeclarationItem{
		Name:       name,
		Type:       typ,
		Expression: expression,
	}
}
