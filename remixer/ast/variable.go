package ast

import (
	"fmt"
	"strings"
)

type ASTVariableDeclaration struct {
	Type  *ASTType
	Items []*ASTVariableDeclarationItem
}

type ASTVariableDeclarationItem struct {
	Name       string
	Expression IASTExpression
}

func NewASTVariableDeclaration(typ *ASTType) *ASTVariableDeclaration {
	return &ASTVariableDeclaration{
		Type: typ,
	}
}

func (v *ASTVariableDeclaration) String() string {
	names := make([]string, 0)
	expressions := make([]string, 0)

	for _, item := range v.Items {
		names = append(names, item.Name)

		if item.Expression != nil {
			expressions = append(expressions, item.Expression.String())
		} else {
			switch v.Type.Kind {
			case ASTTypeKindArray, ASTTypeKindPointer:
				expressions = append(expressions, "nil")
			case ASTTypeKindStruct:
				expressions = append(expressions, fmt.Sprintf("%s{}", v.Type.Name))
			// case ASTTypeKindString:
			// 	expressions = append(expressions, "nil")
			default:
				expressions = append(expressions, "0")
			}
		}
	}

	return fmt.Sprintf("%s := %s", strings.Join(names, ", "), strings.Join(expressions, ", "))
}

func NewASTVariableDeclarationItem(name string, expression IASTExpression) *ASTVariableDeclarationItem {
	return &ASTVariableDeclarationItem{
		Name:       name,
		Expression: expression,
	}
}
