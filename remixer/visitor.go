package remixer

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tereus-project/tereus-remixer-c-go/parser"
	"github.com/tereus-project/tereus-remixer-c-go/remixer/ast"
)

type Visitor struct {
	Path string
	Code string

	Package string
	Imports []string
	Output  map[string]string
}

func NewVisitor(path string) (*Visitor, error) {
	code, e := os.ReadFile(path)
	if e != nil {
		return nil, e
	}

	visitor := &Visitor{
		Path: path,
		Code: string(code),

		Package: "main",
		Imports: make([]string, 0),
		Output:  make(map[string]string, 0),
	}

	return visitor, nil
}

func (v *Visitor) NotImplementedError(token *antlr.BaseParserRuleContext) error {
	start := token.GetStart()
	return fmt.Errorf("%s:%d:%d: %s", v.Path, start.GetLine(), start.GetColumn(), "not implemented")
}

func (v *Visitor) VisitTranslation(ctx *parser.TranslationContext) (string, error) {
	output := "package " + v.Package + "\n\n"

	for _, declaration := range ctx.AllDeclaration() {
		declaration, e := v.VisitDeclaration(declaration.(*parser.DeclarationContext))
		if e != nil {
			return "", e
		}

		output += declaration.String()
	}

	return output, nil
}

func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) (ast.IASTItem, error) {
	if child := ctx.FunctionDeclaration(); child != nil {
		return v.VisitFunctionDeclaration(child.(*parser.FunctionDeclarationContext))
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) (*ast.ASTFunction, error) {
	var e error

	name := ctx.Identifier().GetText()
	function := ast.NewASTFunction(name)

	if name != "main" {
		function.ReturnType, e = v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
		if e != nil {
			return nil, e
		}
	} else {
		function.ReturnType = ast.NewASTType(ast.ASTTypeKindVoid, "void")
	}

	if child := ctx.FunctionArguments(); child != nil {
		function.Args, e = v.VisitFunctionArguments(child.(*parser.FunctionArgumentsContext))
		if e != nil {
			return nil, e
		}
	}

	function.Body, e = v.VisitBlock(ctx.Block().(*parser.BlockContext))
	if e != nil {
		return nil, e
	}

	return function, nil
}

func (v *Visitor) VisitFunctionArguments(ctx *parser.FunctionArgumentsContext) ([]*ast.ASTFunctionArgument, error) {
	name := ctx.Identifier().GetText()

	typ, e := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
	if e != nil {
		return nil, e
	}

	args := []*ast.ASTFunctionArgument{ast.NewASTFunctionArgument(name, typ)}

	if ctx.FunctionArguments() != nil {
		others, e := v.VisitFunctionArguments(ctx.FunctionArguments().(*parser.FunctionArgumentsContext))
		if e != nil {
			return nil, e
		}

		args = append(args, others...)
	}

	return args, nil
}

func (v *Visitor) VisitTypeSpecifier(ctx *parser.TypeSpecifierContext) (*ast.ASTType, error) {
	if child := ctx.Void(); child != nil {
		return ast.NewASTType(ast.ASTTypeKindVoid, "void"), nil
	} else if child := ctx.Int(); child != nil {
		return ast.NewASTType(ast.ASTTypeKindInt, "int"), nil
	} else if child := ctx.Short(); child != nil {
		return ast.NewASTType(ast.ASTTypeKindInt, "int16"), nil
	} else if child := ctx.Long(); child != nil {
		return ast.NewASTType(ast.ASTTypeKindInt, "int64"), nil
	} else if child := ctx.Char(); child != nil {
		return ast.NewASTType(ast.ASTTypeKindRune, "rune"), nil
	} else if child := ctx.Float(); child != nil {
		return ast.NewASTType(ast.ASTTypeKindFloat32, "float32"), nil
	} else if child := ctx.Double(); child != nil {
		return ast.NewASTType(ast.ASTTypeKindFloat64, "float64"), nil
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) (*ast.ASTVariableDeclaration, error) {
	name := ctx.Identifier().GetText()

	typ, e := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
	if e != nil {
		return nil, e
	}

	variable := ast.NewASTVariableDeclaration(name, typ)

	if child := ctx.Expression(); child != nil {
		expression, e := v.VisitExpression(child)
		if e != nil {
			return nil, e
		}

		variable.Expression = expression
	}

	return variable, nil
}

func (v *Visitor) VisitExpression(ctx parser.IExpressionContext) (ast.IASTExpression, error) {
	switch child := ctx.(type) {
	case *parser.IdentifierExpressionContext:
		return ast.NewASTExpressionLiteral(child.GetText()), nil
	case *parser.ConstantExpressionContext:
		return ast.NewASTExpressionLiteral(child.GetText()), nil
	case *parser.ParenthesizedExpressionContext:
		return v.VisitExpression(child.Expression())
	case *parser.AssignmentExpressionContext:
		left, e := v.VisitExpression(child.Expression(0))
		if e != nil {
			return nil, e
		}

		right, e := v.VisitExpression(child.Expression(1))
		if e != nil {
			return nil, e
		}

		return ast.NewASTExpressionBinary(left, child.AssignementOperator().GetText(), right), nil
	}

	return nil, v.NotImplementedError(ctx.(*parser.ExpressionContext).BaseParserRuleContext)
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) (*ast.ASTBlock, error) {
	statements := make([]ast.IASTItem, 0)

	for _, statement := range ctx.AllStatement() {
		statement, e := v.VisitStatement(statement.(*parser.StatementContext))
		if e != nil {
			return nil, e
		}

		statements = append(statements, statement)
	}

	return ast.NewASTBlock(statements), nil
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) (ast.IASTItem, error) {
	if child := ctx.VariableDeclaration(); child != nil {
		return v.VisitVariableDeclaration(child.(*parser.VariableDeclarationContext))
	} else if child := ctx.Expression(); child != nil {
		return v.VisitExpression(child)
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}
