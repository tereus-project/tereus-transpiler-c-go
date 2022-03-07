package remixer

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	mapset "github.com/deckarep/golang-set"
	"github.com/tereus-project/tereus-remixer-c-go/parser"
	"github.com/tereus-project/tereus-remixer-c-go/remixer/ast"
	"github.com/tereus-project/tereus-remixer-c-go/remixer/utils"
)

type Visitor struct {
	Path string
	Code string

	Package string
	Imports mapset.Set
	Output  map[string]string

	CurrentFunction *utils.Stack
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
		Imports: mapset.NewSet(),
		Output:  make(map[string]string),

		CurrentFunction: utils.NewStack(),
	}

	return visitor, nil
}

func (v *Visitor) PositionedTranslationError(start antlr.Token, message string) error {
	return fmt.Errorf("%s:%d:%d: %s", v.Path, start.GetLine(), start.GetColumn(), message)
}

func (v *Visitor) TranslationError(token *antlr.BaseParserRuleContext, message string) error {
	return v.PositionedTranslationError(token.GetStart(), message)
}

func (v *Visitor) NotImplementedError(token *antlr.BaseParserRuleContext) error {
	return v.TranslationError(token, "not implemented")
}

func (v *Visitor) VisitTranslation(ctx *parser.TranslationContext) (string, error) {
	output := "package " + v.Package + "\n\n"

	code := ""

	for _, declaration := range ctx.AllDeclaration() {
		declaration, e := v.VisitDeclaration(declaration.(*parser.DeclarationContext))
		if e != nil {
			return "", e
		}

		code += declaration.String()
	}

	if v.Imports.Cardinality() > 0 {
		output += "import (\n"

		for import_ := range v.Imports.Iter() {
			output += "\t\"" + import_.(string) + "\"\n"
		}

		output += ")\n\n"
	}

	output += code

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

	v.CurrentFunction.Push(name)

	argumentsInitialization := make([]ast.IASTItem, 0)

	if name == "main" {
		if len(function.Args) > 2 {
			return nil, v.TranslationError(ctx.BaseParserRuleContext, "main function can only have 2 arguments")
		}

		if len(function.Args) >= 1 {
			variable := ast.NewASTVariableDeclaration(ast.NewASTType(ast.ASTTypeKindInt, "int"))
			variable.Items = []*ast.ASTVariableDeclarationItem{
				// TODO: change this to a proper structure when it gets implemented
				ast.NewASTVariableDeclarationItem(function.Args[0].Name, ast.NewASTExpressionLiteral("len(os.Args)")),
			}

			v.Imports.Add("os")

			argumentsInitialization = append(argumentsInitialization, variable)
		}

		if len(function.Args) >= 2 {
			typ := ast.NewASTType(ast.ASTTypeKindArray, "array")
			typ.ArrayType = ast.NewASTType(ast.ASTTypeKindRune, "string")

			variable := ast.NewASTVariableDeclaration(typ)
			variable.Items = []*ast.ASTVariableDeclarationItem{
				// TODO: change this to a proper structure when it gets implemented
				ast.NewASTVariableDeclarationItem(function.Args[1].Name, ast.NewASTExpressionLiteral("os.Args")),
			}

			argumentsInitialization = append(argumentsInitialization, variable)
		}

		function.Args = nil
	}

	function.Body, e = v.VisitBlock(ctx.Block().(*parser.BlockContext))
	if e != nil {
		return nil, e
	}

	function.Body.Statements = append(argumentsInitialization, function.Body.Statements...)

	v.CurrentFunction.Pop()

	return function, nil
}

func (v *Visitor) VisitFunctionArguments(ctx *parser.FunctionArgumentsContext) ([]*ast.ASTFunctionArgument, error) {
	name := ctx.Identifier().GetText()

	typ, e := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
	if e != nil {
		return nil, e
	}

	args := []*ast.ASTFunctionArgument{ast.NewASTFunctionArgument(name, typ)}

	if child := ctx.FunctionArguments(); child != nil {
		others, e := v.VisitFunctionArguments(child.(*parser.FunctionArgumentsContext))
		if e != nil {
			return nil, e
		}

		args = append(args, others...)
	}

	return args, nil
}

func (v *Visitor) VisitFunctionReturn(ctx *parser.FunctionReturnContext) (ast.IASTItem, error) {
	expression, e := v.VisitExpression(ctx.Expression())
	if e != nil {
		return nil, e
	}

	if v.CurrentFunction.Top() == "main" {
		v.Imports.Add("os")
		return ast.NewASTExpressionFunctionCall(ast.NewASTExpressionLiteral("os.Exit"), []ast.IASTExpression{expression}), nil
	}

	return ast.NewASTFunctionReturn(expression), nil
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
	} else if child := ctx.Star(); child != nil {
		pointerType, e := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
		if e != nil {
			return nil, e
		}

		typ := ast.NewASTType(ast.ASTTypeKindPointer, "pointer")
		typ.PointerType = pointerType

		return typ, nil
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) (*ast.ASTVariableDeclaration, error) {
	typ, e := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
	if e != nil {
		return nil, e
	}

	variable := ast.NewASTVariableDeclaration(typ)

	items, e := v.VisitVariableDeclarationList(ctx.VariableDeclarationList().(*parser.VariableDeclarationListContext))
	if e != nil {
		return nil, e
	}

	variable.Items = items

	return variable, nil
}

func (v *Visitor) VisitVariableDeclarationList(ctx *parser.VariableDeclarationListContext) ([]*ast.ASTVariableDeclarationItem, error) {
	name := ctx.Identifier().GetText()

	item := ast.NewASTVariableDeclarationItem(name, nil)

	if child := ctx.Expression(); child != nil {
		expression, e := v.VisitExpression(child)
		if e != nil {
			return nil, e
		}

		item.Expression = expression
	}

	items := []*ast.ASTVariableDeclarationItem{item}

	if child := ctx.VariableDeclarationList(); child != nil {
		others, e := v.VisitVariableDeclarationList(child.(*parser.VariableDeclarationListContext))
		if e != nil {
			return nil, e
		}

		items = append(items, others...)
	}

	return items, nil
}

func (v *Visitor) VisitExpression(ctx parser.IExpressionContext) (ast.IASTExpression, error) {
	switch child := ctx.(type) {
	case *parser.IdentifierExpressionContext:
		return ast.NewASTExpressionLiteral(child.GetText()), nil
	case *parser.ConstantExpressionContext:
		return ast.NewASTExpressionLiteral(child.GetText()), nil
	case *parser.ParenthesizedExpressionContext:
		expression, e := v.VisitExpression(child.Expression())
		if e != nil {
			return nil, e
		}

		return ast.NewAstParenthesizedExpression(expression), nil
	case *parser.UnaryExpressionPostContext:
		left, e := v.VisitExpression(child.Expression())
		if e != nil {
			return nil, e
		}

		return ast.NewASTExpressionUnaryPost(left, child.UnaryOperatorPost().GetText()), nil
	case *parser.UnaryExpressionPreContext:
		right, e := v.VisitExpression(child.Expression())
		if e != nil {
			return nil, e
		}

		return ast.NewASTExpressionUnaryPre(child.UnaryOperatorPre().GetText(), right), nil
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
	case *parser.BinaryExpressionContext:
		left, e := v.VisitExpression(child.Expression(0))
		if e != nil {
			return nil, e
		}

		right, e := v.VisitExpression(child.Expression(1))
		if e != nil {
			return nil, e
		}

		return ast.NewASTExpressionBinary(left, child.BinaryOperator().GetText(), right), nil
	case *parser.FunctionCallExpressionContext:
		expression, e := v.VisitExpression(child.Expression())
		if e != nil {
			return nil, e
		}

		var args []ast.IASTExpression

		if child := child.FunctionCallArguments(); child != nil {
			args, e = v.VisitFunctionCallArguments(child.(*parser.FunctionCallArgumentsContext))
			if e != nil {
				return nil, e
			}
		}

		return ast.NewASTExpressionFunctionCall(expression, args), nil
	}

	return nil, v.PositionedTranslationError(ctx.GetStart(), "not implemented")
}

func (v *Visitor) VisitFunctionCallArguments(ctx *parser.FunctionCallArgumentsContext) ([]ast.IASTExpression, error) {
	expression, e := v.VisitExpression(ctx.Expression())
	if e != nil {
		return nil, e
	}

	args := []ast.IASTExpression{expression}

	if child := ctx.FunctionCallArguments(); child != nil {
		others, e := v.VisitFunctionCallArguments(child.(*parser.FunctionCallArgumentsContext))
		if e != nil {
			return nil, e
		}

		args = append(args, others...)
	}

	return args, nil
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
		variableDeclaration, e := v.VisitVariableDeclaration(child.(*parser.VariableDeclarationContext))
		if e != nil {
			return nil, e
		}

		return variableDeclaration, nil
	} else if child := ctx.Expression(); child != nil {
		return v.VisitExpression(child)
	} else if child := ctx.FunctionReturn(); child != nil {
		return v.VisitFunctionReturn(child.(*parser.FunctionReturnContext))
	} else if child := ctx.IfStatement(); child != nil {
		return v.VisitIfStatement(child.(*parser.IfStatementContext))
	} else if child := ctx.ForStatement(); child != nil {
		return v.VisitForStatement(child.(*parser.ForStatementContext))
	} else if child := ctx.WhileStatement(); child != nil {
		return v.VisitWhileStatement(child.(*parser.WhileStatementContext))
	} else if child := ctx.Block(); child != nil {
		return v.VisitBlock(child.(*parser.BlockContext))
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) (*ast.ASTIf, error) {
	condition, e := v.VisitExpression(ctx.Expression())
	if e != nil {
		return nil, e
	}

	then, e := v.VisitStatement(ctx.Statement(0).(*parser.StatementContext))
	if e != nil {
		return nil, e
	}

	if_ := ast.NewASTIf(condition, then)

	if ctx.Else() != nil {
		if_.Else, e = v.VisitStatement(ctx.Statement(1).(*parser.StatementContext))
		if e != nil {
			return nil, e
		}
	}

	return if_, nil
}

func (v *Visitor) VisitForStatement(ctx *parser.ForStatementContext) (*ast.ASTFor, error) {
	for_ := ast.NewASTFor()

	if child := ctx.GetInit(); child != nil {
		expression, e := v.VisitExpression(child)
		if e != nil {
			return nil, e
		}

		for_.Init = expression
	}

	if child := ctx.VariableDeclaration(); child != nil {
		variableDeclaration, e := v.VisitVariableDeclaration(child.(*parser.VariableDeclarationContext))
		if e != nil {
			return nil, e
		}

		for_.Init = variableDeclaration
	}

	if child := ctx.GetCondition(); child != nil {
		expression, e := v.VisitExpression(child)
		if e != nil {
			return nil, e
		}

		for_.Cond = expression
	}

	if child := ctx.GetPost(); child != nil {
		expression, e := v.VisitExpression(child)
		if e != nil {
			return nil, e
		}

		for_.Post = expression
	}

	body, e := v.VisitStatement(ctx.Statement().(*parser.StatementContext))
	if e != nil {
		return nil, e
	}

	for_.Statement = body

	return for_, e
}

func (v *Visitor) VisitWhileStatement(ctx *parser.WhileStatementContext) (*ast.ASTWhile, error) {
	cond, e := v.VisitExpression(ctx.Expression())
	if e != nil {
		return nil, e
	}

	statement, e := v.VisitStatement(ctx.Statement().(*parser.StatementContext))
	if e != nil {
		return nil, e
	}

	return ast.NewASTWhile(cond, statement), nil
}
