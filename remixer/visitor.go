package remixer

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	mapset "github.com/deckarep/golang-set"
	"github.com/tereus-project/tereus-remixer-c-go/libc"
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

	Scope           *Scope
	CurrentFunction *utils.Stack
}

func NewVisitor(path string) (*Visitor, error) {
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	visitor := &Visitor{
		Path: path,
		Code: string(code),

		Package: "main",
		Imports: mapset.NewSet(),
		Output:  make(map[string]string),

		Scope:           NewScope(),
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
		declaration, err := v.VisitDeclaration(declaration.(*parser.DeclarationContext))
		if err != nil {
			return "", err
		}

		if declaration != nil {
			code += declaration.String() + "\n\n"
		}
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
	} else if child := ctx.StructDeclaration(); child != nil {
		return v.VisitStructDeclaration(child.(*parser.StructDeclarationContext))
	} else if child := ctx.IncludePreprocessor(); child != nil {
		return v.VisitIncludePreprocessor(child.(*parser.IncludePreprocessorContext))
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) (*ast.ASTFunction, error) {
	var err error

	name := ctx.Identifier().GetText()
	function := ast.NewASTFunction(name)

	if name != "main" {
		function.ReturnType, err = v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
		if err != nil {
			return nil, err
		}
	} else {
		function.ReturnType = ast.NewASTType(ast.ASTTypeKindVoid, "void")
	}

	if child := ctx.FunctionArguments(); child != nil {
		function.Args, err = v.VisitFunctionArguments(child.(*parser.FunctionArgumentsContext))
		if err != nil {
			return nil, err
		}
	}

	v.CurrentFunction.Push(name)
	v.Scope.Push()

	for _, arg := range function.Args {
		v.Scope.Add(NewScopeVariable(arg.Name, "", arg.Type))
	}

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
			typ.ArrayType = ast.NewASTType(ast.ASTTypeKindChar, "string")

			variable := ast.NewASTVariableDeclaration(typ)
			variable.Items = []*ast.ASTVariableDeclarationItem{
				// TODO: change this to a proper structure when it gets implemented
				ast.NewASTVariableDeclarationItem(function.Args[1].Name, ast.NewASTExpressionLiteral("os.Args")),
			}

			argumentsInitialization = append(argumentsInitialization, variable)
		}

		function.Args = nil
	}

	variableType := ast.NewASTType(ast.ASTTypeKindFunction, function.Name)
	variableType.FunctionType = function
	v.Scope.Parent.Add(NewScopeVariable(function.Name, "", variableType))

	function.Body, err = v.VisitBlock(ctx.Block().(*parser.BlockContext))
	if err != nil {
		return nil, err
	}

	function.Body.Statements = append(argumentsInitialization, function.Body.Statements...)

	v.CurrentFunction.Pop()
	v.Scope.Pop()

	return function, nil
}

func (v *Visitor) VisitFunctionArguments(ctx *parser.FunctionArgumentsContext) ([]*ast.ASTFunctionArgument, error) {
	name := ctx.Identifier().GetText()

	typ, err := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
	if err != nil {
		return nil, err
	}

	args := []*ast.ASTFunctionArgument{ast.NewASTFunctionArgument(name, typ)}

	if child := ctx.FunctionArguments(); child != nil {
		others, err := v.VisitFunctionArguments(child.(*parser.FunctionArgumentsContext))
		if err != nil {
			return nil, err
		}

		args = append(args, others...)
	}

	return args, nil
}

func (v *Visitor) VisitFunctionReturn(ctx *parser.FunctionReturnContext) (ast.IASTItem, error) {
	expression, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
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
		return ast.NewASTType(ast.ASTTypeKindChar, "int8"), nil
	} else if child := ctx.Float(); child != nil {
		return ast.NewASTType(ast.ASTTypeKindFloat32, "float32"), nil
	} else if child := ctx.Double(); child != nil {
		return ast.NewASTType(ast.ASTTypeKindFloat64, "float64"), nil
	} else if child := ctx.Star(); child != nil {
		pointerType, err := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
		if err != nil {
			return nil, err
		}

		typ := ast.NewASTType(ast.ASTTypeKindPointer, "pointer")
		typ.PointerType = pointerType

		return typ, nil
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) (ast.IASTItem, error) {
	name := "_"

	if child := ctx.Identifier(); child != nil {
		name = child.GetText()
	}

	properties := make([]*ast.ASTStructProperty, 0)

	for _, child := range ctx.AllStructProperty() {
		member, err := v.VisitStructProperty(child.(*parser.StructPropertyContext))
		if err != nil {
			return nil, err
		}

		properties = append(properties, member)
	}

	return ast.NewASTStruct(name, properties), nil
}

func (v *Visitor) VisitStructProperty(ctx *parser.StructPropertyContext) (*ast.ASTStructProperty, error) {
	name := "_"

	if child := ctx.Identifier(); child != nil {
		name = child.GetText()
	}

	typ, err := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
	if err != nil {
		return nil, err
	}

	return ast.NewASTStructProperty(name, typ), nil
}

func (v *Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) (*ast.ASTVariableDeclaration, error) {
	typ, err := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
	if err != nil {
		return nil, err
	}

	variable := ast.NewASTVariableDeclaration(typ)

	items, err := v.VisitVariableDeclarationList(ctx.VariableDeclarationList().(*parser.VariableDeclarationListContext), typ)
	if err != nil {
		return nil, err
	}

	variable.Items = items

	return variable, nil
}

func (v *Visitor) VisitVariableDeclarationList(ctx *parser.VariableDeclarationListContext, typ *ast.ASTType) ([]*ast.ASTVariableDeclarationItem, error) {
	name := ctx.Identifier().GetText()

	item := ast.NewASTVariableDeclarationItem(name, nil)

	if child := ctx.Expression(); child != nil {
		expression, err := v.VisitExpression(child)
		if err != nil {
			return nil, err
		}

		item.Expression = expression
	}

	v.Scope.Add(NewScopeVariable(name, "", typ))

	items := []*ast.ASTVariableDeclarationItem{item}

	if child := ctx.VariableDeclarationList(); child != nil {
		others, err := v.VisitVariableDeclarationList(child.(*parser.VariableDeclarationListContext), typ)
		if err != nil {
			return nil, err
		}

		items = append(items, others...)
	}

	return items, nil
}

func (v *Visitor) VisitExpression(ctx parser.IExpressionContext) (ast.IASTExpression, error) {
	switch child := ctx.(type) {
	case *parser.IdentifierExpressionContext:
		return v.VisitIdentifierExpression(child)
	case *parser.ConstantExpressionContext:
		return ast.NewASTExpressionLiteral(child.GetText()), nil
	case *parser.ConstantStringExpressionContext:
		return ast.NewASTExpressionLiteral(child.GetText()), nil
	case *parser.ParenthesizedExpressionContext:
		expression, err := v.VisitExpression(child.Expression())
		if err != nil {
			return nil, err
		}

		return ast.NewAstParenthesizedExpression(expression), nil
	case *parser.UnaryExpressionPostContext:
		left, err := v.VisitExpression(child.Expression())
		if err != nil {
			return nil, err
		}

		return ast.NewASTExpressionUnaryPost(left, child.UnaryOperatorPost().GetText()), nil
	case *parser.UnaryExpressionPreContext:
		right, err := v.VisitExpression(child.Expression())
		if err != nil {
			return nil, err
		}

		operator := child.UnaryOperatorPre().GetText()

		if operator == "++" || operator == "--" {
			return nil, v.TranslationError(child.BaseParserRuleContext, fmt.Sprintf("unary operator '%s' not supported by Go", operator))
		}

		return ast.NewASTExpressionUnaryPre(operator, right), nil
	case *parser.AssignmentExpressionContext:
		left, err := v.VisitExpression(child.Expression(0))
		if err != nil {
			return nil, err
		}

		right, err := v.VisitExpression(child.Expression(1))
		if err != nil {
			return nil, err
		}

		return ast.NewASTExpressionBinary(left, child.AssignementOperator().GetText(), right), nil
	case *parser.BinaryExpressionContext:
		left, err := v.VisitExpression(child.Expression(0))
		if err != nil {
			return nil, err
		}

		right, err := v.VisitExpression(child.Expression(1))
		if err != nil {
			return nil, err
		}

		return ast.NewASTExpressionBinary(left, child.BinaryOperator().GetText(), right), nil
	case *parser.FunctionCallExpressionContext:
		return v.VisitFunctionCallExpression(child)
	}

	return nil, v.PositionedTranslationError(ctx.GetStart(), "not implemented")
}

func (v *Visitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) (*ast.ASTExpressionLiteral, error) {
	if items := v.Scope.Get(ctx.GetText()); len(items) > 0 {
		return ast.NewASTExpressionLiteral(items[0].GetTranslatedName()), nil
	}

	identifier := ""

	switch ctx.GetText() {
	case "printf":
		v.Imports.Add("fmt")
		identifier = "fmt.Printf"
	case "scanf":
		v.Imports.Add("fmt")
		identifier = "fmt.Scanf"
	case "malloc":
		identifier = "libc.Malloc"
	case "free":
		identifier = "libc.Free"
	}

	if identifier != "" {
		return ast.NewASTExpressionLiteral(identifier), nil
	}

	return nil, v.PositionedTranslationError(ctx.GetStart(), fmt.Sprintf("identifier '%s' not found", ctx.GetText()))
}

func (v *Visitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) (*ast.ASTExpressionFunctionCall, error) {
	expression, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	var args []ast.IASTExpression

	if child := ctx.FunctionCallArguments(); child != nil {
		args, err = v.VisitFunctionCallArguments(child.(*parser.FunctionCallArgumentsContext))
		if err != nil {
			return nil, err
		}
	}

	return ast.NewASTExpressionFunctionCall(expression, args), nil
}

func (v *Visitor) VisitFunctionCallArguments(ctx *parser.FunctionCallArgumentsContext) ([]ast.IASTExpression, error) {
	expression, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	args := []ast.IASTExpression{expression}

	if child := ctx.FunctionCallArguments(); child != nil {
		others, err := v.VisitFunctionCallArguments(child.(*parser.FunctionCallArgumentsContext))
		if err != nil {
			return nil, err
		}

		args = append(args, others...)
	}

	return args, nil
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) (*ast.ASTBlock, error) {
	statements := make([]ast.IASTItem, 0)

	v.Scope.Push()

	for _, statement := range ctx.AllStatement() {
		statement, err := v.VisitStatement(statement.(*parser.StatementContext))
		if err != nil {
			return nil, err
		}

		statements = append(statements, statement)
	}

	v.Scope.Pop()

	return ast.NewASTBlock(statements), nil
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) (ast.IASTItem, error) {
	if child := ctx.VariableDeclaration(); child != nil {
		variableDeclaration, err := v.VisitVariableDeclaration(child.(*parser.VariableDeclarationContext))
		if err != nil {
			return nil, err
		}

		return variableDeclaration, nil
	} else if child := ctx.Expression(); child != nil {
		return v.VisitExpression(child)
	} else if child := ctx.FunctionReturn(); child != nil {
		return v.VisitFunctionReturn(child.(*parser.FunctionReturnContext))
	} else if child := ctx.Break(); child != nil {
		return ast.NewASTBreak(), nil
	} else if child := ctx.StructDeclaration(); child != nil {
		return v.VisitStructDeclaration(child.(*parser.StructDeclarationContext))
	} else if child := ctx.IfStatement(); child != nil {
		return v.VisitIfStatement(child.(*parser.IfStatementContext))
	} else if child := ctx.ForStatement(); child != nil {
		return v.VisitForStatement(child.(*parser.ForStatementContext))
	} else if child := ctx.WhileStatement(); child != nil {
		return v.VisitWhileStatement(child.(*parser.WhileStatementContext))
	} else if child := ctx.Block(); child != nil {
		return v.VisitBlock(child.(*parser.BlockContext))
	} else if child := ctx.BlockComment(); child != nil {
		return ast.NewASTComment(true, child.GetText()), nil
	} else if child := ctx.LineComment(); child != nil {
		return ast.NewASTComment(false, child.GetText()), nil
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) (*ast.ASTIf, error) {
	v.Scope.Push()

	condition, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	then, err := v.VisitStatement(ctx.Statement(0).(*parser.StatementContext))
	if err != nil {
		return nil, err
	}

	v.Scope.Pop()

	if_ := ast.NewASTIf(condition, then)

	if ctx.Else() != nil {
		v.Scope.Push()

		if_.Else, err = v.VisitStatement(ctx.Statement(1).(*parser.StatementContext))
		if err != nil {
			return nil, err
		}

		v.Scope.Pop()
	}

	return if_, nil
}

func (v *Visitor) VisitForStatement(ctx *parser.ForStatementContext) (*ast.ASTFor, error) {
	v.Scope.Push()

	for_ := ast.NewASTFor()

	if child := ctx.GetInit(); child != nil {
		expression, err := v.VisitExpression(child)
		if err != nil {
			return nil, err
		}

		for_.Init = expression
	}

	if child := ctx.VariableDeclaration(); child != nil {
		variableDeclaration, err := v.VisitVariableDeclaration(child.(*parser.VariableDeclarationContext))
		if err != nil {
			return nil, err
		}

		for_.Init = variableDeclaration
	}

	if child := ctx.GetCondition(); child != nil {
		expression, err := v.VisitExpression(child)
		if err != nil {
			return nil, err
		}

		for_.Cond = expression
	}

	if child := ctx.GetPost(); child != nil {
		expression, err := v.VisitExpression(child)
		if err != nil {
			return nil, err
		}

		for_.Post = expression
	}

	body, err := v.VisitStatement(ctx.Statement().(*parser.StatementContext))
	if err != nil {
		return nil, err
	}

	for_.Statement = body

	v.Scope.Pop()

	return for_, nil
}

func (v *Visitor) VisitWhileStatement(ctx *parser.WhileStatementContext) (*ast.ASTWhile, error) {
	v.Scope.Push()

	cond, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	statement, err := v.VisitStatement(ctx.Statement().(*parser.StatementContext))
	if err != nil {
		return nil, err
	}

	v.Scope.Pop()

	return ast.NewASTWhile(cond, statement), nil
}

var (
	systemIncludePreprocessorRegex = regexp.MustCompile(`^#[ \t]*include[ \t]*<(.+)>$`)
	localIncludePreprocessorRegex  = regexp.MustCompile(`^#[ \t]*include[ \t]*"(.+)"$`)
)

func (v *Visitor) VisitIncludePreprocessor(ctx *parser.IncludePreprocessorContext) (ast.IASTItem, error) {
	directive := strings.TrimSpace(ctx.IncludeDirective().GetText())

	if matches := systemIncludePreprocessorRegex.FindStringSubmatch(directive); len(matches) > 0 {
		header := matches[1]

		if libc.IsSupported(header) {
			v.Imports.Add("github.com/tereus-project/tereus-remixer-c-go/libc")
		}
	} else if matches := localIncludePreprocessorRegex.FindStringSubmatch(directive); len(matches) > 0 {
		return nil, v.TranslationError(ctx.BaseParserRuleContext, "unsupported include type")
	} else {
		return nil, v.TranslationError(ctx.BaseParserRuleContext, "unsupported include type")
	}

	return nil, nil
}
