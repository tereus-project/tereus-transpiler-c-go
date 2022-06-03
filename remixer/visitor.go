package remixer

import (
	"fmt"
	"go/format"
	"regexp"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	mapset "github.com/deckarep/golang-set"
	"github.com/tereus-project/tereus-remixer-c-go/libc"
	"github.com/tereus-project/tereus-remixer-c-go/parser"
	"github.com/tereus-project/tereus-remixer-c-go/remixer/ast"
	"github.com/tereus-project/tereus-remixer-c-go/remixer/scope"
	"github.com/tereus-project/tereus-remixer-c-go/remixer/utils"
)

type Visitor struct {
	Path string
	Code string

	Package string
	Imports mapset.Set
	Output  map[string]string

	Scope           *scope.Scope
	CurrentFunction *utils.Stack[string]
}

func NewVisitor(path string, code string) *Visitor {
	return &Visitor{
		Path: path,
		Code: code,

		Package: "main",
		Imports: mapset.NewSet(),
		Output:  make(map[string]string),

		Scope:           scope.NewScope(),
		CurrentFunction: utils.NewStack[string](),
	}
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
	formatted, err := format.Source([]byte(output))
	if err != nil {
		return output, fmt.Errorf("formatting error: %s", err)
	}

	formatted, err = imports.Process(v.Path, formatted, nil)
	if err != nil {
		return output, fmt.Errorf("imports formatting error: %s", err)
	}

	return string(formatted), nil
}

func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) (ast.IASTItem, error) {
	if child := ctx.FunctionDeclaration(); child != nil {
		return v.VisitFunctionDeclaration(child.(*parser.FunctionDeclarationContext))
	} else if child := ctx.StructDeclaration(); child != nil {
		return v.VisitStructDeclaration(child.(*parser.StructDeclarationContext))
	} else if child := ctx.EnumDeclaration(); child != nil {
		return v.VisitEnumDeclaration(child.(*parser.EnumDeclarationContext))
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
		v.Scope.Add(scope.NewScopeVariable(arg.Name, "", arg.Type))
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
				ast.NewASTVariableDeclarationItem(function.Args[0].Name, ast.NewASTExpressionLiteral("len(os.Args)", ast.NewASTType(ast.ASTTypeKindInt, "int"))),
			}

			v.Imports.Add("os")

			argumentsInitialization = append(argumentsInitialization, variable)
		}

		if len(function.Args) >= 2 {
			typ := ast.NewASTType(ast.ASTTypeKindArray, "array").
				SetArrayType(
					ast.NewASTType(ast.ASTTypeKindChar, "char"),
				)

			variable := ast.NewASTVariableDeclaration(typ)

			typ_ := ast.NewASTType(ast.ASTTypeKindPointer, "pointer").
				SetPointerType(
					ast.NewASTType(ast.ASTTypeKindPointer, "pointer").
						SetPointerType(
							ast.NewASTType(ast.ASTTypeKindChar, "char"),
						),
				)

			variable.Items = []*ast.ASTVariableDeclarationItem{
				// TODO: change this to a proper structure when it gets implemented
				ast.NewASTVariableDeclarationItem(function.Args[1].Name, ast.NewASTExpressionLiteral("os.Args", typ_)),
			}

			argumentsInitialization = append(argumentsInitialization, variable)
		}

		function.Args = nil
	}

	variableType := ast.NewASTType(ast.ASTTypeKindFunction, function.Name)
	variableType.FunctionType = function
	v.Scope.Parent.Add(scope.NewScopeVariable(function.Name, "", variableType))

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

		typ_ := ast.NewASTType(ast.ASTTypeKindFunction, "function")
		functionType := ast.NewASTFunction("os.Exit")
		functionType.Args = []*ast.ASTFunctionArgument{
			ast.NewASTFunctionArgument("code", ast.NewASTType(ast.ASTTypeKindInt, "int")),
		}
		functionType.ReturnType = ast.NewASTType(ast.ASTTypeKindVoid, "void")

		return ast.NewASTExpressionFunctionCall(ast.NewASTExpressionLiteral("os.Exit", typ_), []ast.IASTExpression{expression}), nil
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
	} else if child := ctx.StructDeclaration(); child != nil {
		structType, err := v.VisitStructDeclaration(child.(*parser.StructDeclarationContext))
		if err != nil {
			return nil, err
		}

		typ := ast.NewASTType(ast.ASTTypeKindStruct, structType.Name)
		typ.SetStructType(structType)

		return typ, nil
	} else if child := ctx.Star(); child != nil {
		pointerType, err := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
		if err != nil {
			return nil, err
		}

		typ := ast.NewASTType(ast.ASTTypeKindPointer, "pointer")
		typ.SetPointerType(pointerType)

		return typ, nil
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) (*ast.ASTStruct, error) {
	name := "_"

	if child := ctx.Identifier(); child != nil {
		name = child.GetText()
	}

	item := v.Scope.GetFirst(name)

	var astStruct *ast.ASTStruct
	if item != nil {
		scopeType := item.(*scope.ScopeType)
		if scopeType == nil {
			return nil, v.PositionedTranslationError(ctx.GetStart(), fmt.Sprintf("redefinition of '%s' as different kind of symbol", name))
		}

		scopeTypeUnderlyingType := scopeType.GetType()

		if scopeTypeUnderlyingType.Kind != ast.ASTTypeKindStruct {
			return nil, v.PositionedTranslationError(ctx.GetStart(), fmt.Sprintf("redefinition of '%s' as different kind of symbol", name))
		}

		astStruct = scopeTypeUnderlyingType.StructType
	} else {
		astStruct = ast.NewAstStructOpaque(name)

		typ := ast.NewASTType(ast.ASTTypeKindStruct, name)
		typ.SetStructType(astStruct)
		v.Scope.Add(scope.NewScopeType(name, name, typ))
	}

	if StructDeclarationBodyContext := ctx.StructDeclarationBody(); StructDeclarationBodyContext != nil {
		if !astStruct.IsOpaque {
			return nil, v.PositionedTranslationError(ctx.GetStart(), fmt.Sprintf("redefinition of '%s'", name))
		}

		properties, err := v.VisitStructDeclarationBody(StructDeclarationBodyContext.(*parser.StructDeclarationBodyContext))
		if err != nil {
			return nil, err
		}

		astStruct.SetProperties(properties)
	}

	return astStruct, nil
}

func (v *Visitor) VisitStructDeclarationBody(ctx *parser.StructDeclarationBodyContext) ([]*ast.ASTStructProperty, error) {
	structPropertiesContext := ctx.AllStructProperty()
	structProperties := make([]*ast.ASTStructProperty, len(structPropertiesContext))

	for i, child := range structPropertiesContext {
		structProperty, err := v.VisitStructProperty(child.(*parser.StructPropertyContext))
		if err != nil {
			return nil, err
		}

		structProperties[i] = structProperty
	}

	return structProperties, nil
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

func (v *Visitor) VisitEnumDeclaration(ctx *parser.EnumDeclarationContext) (*ast.ASTEnum, error) {
	var name string
	if child := ctx.Identifier(); child != nil {
		name = child.GetText()
	}

	properties, err := v.VisitEnumProperties(ctx.EnumProperties().(*parser.EnumPropertiesContext))
	if err != nil {
		return nil, err
	}

	return ast.NewASTEnum(name, properties), nil
}

func (v *Visitor) VisitEnumProperties(ctx *parser.EnumPropertiesContext) ([]*ast.ASTEnumProperty, error) {
	var err error

	var name string
	if child := ctx.Identifier(); child != nil {
		name = child.GetText()
	}

	var expression ast.IASTExpression
	if child := ctx.Expression(); child != nil {
		expression, err = v.VisitExpression(child)
		if err != nil {
			return nil, err
		}
	}

	v.Scope.Add(scope.NewScopeVariable(name, "", ast.NewASTType(ast.ASTTypeKindInt64, "int64")))

	properties := []*ast.ASTEnumProperty{ast.NewASTEnumProperty(name, expression)}

	if child := ctx.EnumProperties(); child != nil {
		others, err := v.VisitEnumProperties(child.(*parser.EnumPropertiesContext))
		if err != nil {
			return nil, err
		}

		properties = append(properties, others...)
	}

	return properties, nil
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
	variable := ast.NewASTVariableDeclarationItem(name, nil)

	if child := ctx.Expression(); child != nil {
		expression, err := v.VisitExpression(child)
		if err != nil {
			return nil, err
		}

		converted, err := ast.NewAstTypeConversion(expression, typ)
		if err != nil {
			return nil, err
		}

		variable.Expression = converted
	} else if child := ctx.ListInitialization(); child != nil {
		expressions, err := v.VisitListInitialization(child.(*parser.ListInitializationContext))
		if err != nil {
			return nil, err
		}

		if typ.IsStruct() {
			if len(expressions) > len(typ.StructType.Properties) {
				return nil, v.PositionedTranslationError(ctx.GetStart(), fmt.Sprintf("too many initializers for '%s'", typ.Name))
			}

			for i, expression := range expressions {
				converted, err := ast.NewAstTypeConversion(expression, typ.StructType.GetPropertyByIndex(i).Type)
				if err != nil {
					return nil, err
				}

				expressions[i] = converted
			}

			variable.Expression = ast.NewAstStructInitialization(typ, expressions)
		} else if len(expressions) == 1 {
			variable.Expression = expressions[0]
		} else {
			return nil, v.PositionedTranslationError(ctx.GetStart(), fmt.Sprintf("too many initializers for '%s'", typ.Name))
		}
	}

	v.Scope.Add(scope.NewScopeVariable(name, "", typ))

	variables := []*ast.ASTVariableDeclarationItem{variable}

	if child := ctx.VariableDeclarationList(); child != nil {
		others, err := v.VisitVariableDeclarationList(child.(*parser.VariableDeclarationListContext), typ)
		if err != nil {
			return nil, err
		}

		variables = append(variables, others...)
	}

	return variables, nil
}

func (v *Visitor) VisitListInitialization(ctx *parser.ListInitializationContext) ([]ast.IASTExpression, error) {
	expressionsContext := ctx.AllExpression()
	items := make([]ast.IASTExpression, len(expressionsContext))

	for i, child := range expressionsContext {
		expression, err := v.VisitExpression(child)
		if err != nil {
			return nil, err
		}

		items[i] = expression
	}

	return items, nil
}

func (v *Visitor) VisitNamedListInitialization(ctx *parser.NamedListInitializationContext) ([]string, []ast.IASTExpression, error) {
	namedListInitializationItems := ctx.AllNamedListInitializationItem()

	names := make([]string, len(namedListInitializationItems))
	expressions := make([]ast.IASTExpression, len(namedListInitializationItems))

	for i, child := range namedListInitializationItems {
		ctx := child.(*parser.NamedListInitializationItemContext)

		expression, err := v.VisitExpression(ctx.Expression())
		if err != nil {
			return nil, nil, err
		}

		names[i] = ctx.Identifier().GetText()
		expressions[i] = expression
	}

	return names, expressions, nil
}

func (v *Visitor) VisitExpression(ctx parser.IExpressionContext) (ast.IASTExpression, error) {
	return v.VisitExpressionWithConfigurableIsStatement(ctx, false)
}

func (v *Visitor) VisitExpressionWithConfigurableIsStatement(ctx parser.IExpressionContext, isStatement bool) (ast.IASTExpression, error) {
	switch child := ctx.(type) {
	case *parser.IdentifierExpressionContext:
		return v.VisitIdentifierExpression(child)
	case *parser.ConstantExpressionContext:
		typ_ := ast.NewASTType(ast.ASTTypeKindInt, "int")
		return ast.NewASTExpressionLiteral(child.GetText(), typ_), nil
	case *parser.ConstantStringExpressionContext:
		typ_ := ast.NewASTType(ast.ASTTypeKindArray, "array")
		typ_.ArrayType = ast.NewASTType(ast.ASTTypeKindChar, "char")
		return ast.NewASTExpressionLiteral(child.GetText(), typ_), nil
	case *parser.StructInitializationExpressionContext:
		typ, err := v.VisitTypeSpecifier(child.TypeSpecifier().(*parser.TypeSpecifierContext))
		if err != nil {
			return nil, err
		}

		if !typ.IsStruct() {
			return nil, v.PositionedTranslationError(child.GetStart(), fmt.Sprintf("'%s' is not a struct", typ.Name))
		}

		if child := child.ListInitialization(); child != nil {
			expressions, err := v.VisitListInitialization(child.(*parser.ListInitializationContext))
			if err != nil {
				return nil, err
			}

			if len(expressions) > len(typ.StructType.Properties) {
				return nil, v.PositionedTranslationError(ctx.GetStart(), fmt.Sprintf("too many initializers for '%s'", typ.Name))
			}

			for i, expression := range expressions {
				converted, err := ast.NewAstTypeConversion(expression, typ.StructType.GetPropertyByIndex(i).Type)
				if err != nil {
					return nil, err
				}

				expressions[i] = converted
			}

			return ast.NewAstStructInitialization(typ, expressions), nil
		}

		if child := child.NamedListInitialization(); child != nil {
			names, expressions, err := v.VisitNamedListInitialization(child.(*parser.NamedListInitializationContext))
			if err != nil {
				return nil, err
			}

			for i, expression := range expressions {
				converted, err := ast.NewAstTypeConversion(expression, typ.StructType.GetProperty(names[i]).Type)
				if err != nil {
					return nil, err
				}

				expressions[i] = converted
			}

			return ast.NewAstStructNamedInitialization(typ, names, expressions), nil
		}

		return nil, v.PositionedTranslationError(ctx.GetStart(), "Not implemented")
	case *parser.PropertyAccessExpressionContext:
		expression, err := v.VisitExpression(child.Expression())
		if err != nil {
			return nil, err
		}

		expressionType := expression.GetType()
		if expressionType.Kind != ast.ASTTypeKindStruct {
			return nil, v.PositionedTranslationError(child.Expression().GetStart(), "property access is only allowed on structs")
		}

		propertyName := child.Identifier().GetText()
		property := expressionType.StructType.GetProperty(propertyName)
		if property == nil {
			return nil, v.PositionedTranslationError(child.Identifier().GetSymbol(), fmt.Sprintf("struct has no property named '%s'", propertyName))
		}

		return ast.NewASTExpressionPropertyAccess(expression, property), nil
	case *parser.PointerPropertyAccessExpressionContext:
		expression, err := v.VisitExpression(child.Expression())
		if err != nil {
			return nil, err
		}

		expressionType := expression.GetType()

		if (expressionType.Kind != ast.ASTTypeKindPointer) || (expressionType.PointerType.Kind != ast.ASTTypeKindStruct) {
			return nil, v.PositionedTranslationError(child.Expression().GetStart(), "pointer property access is only allowed on struct pointers")
		}

		propertyName := child.Identifier().GetText()
		property := expressionType.PointerType.StructType.GetProperty(propertyName)
		if property == nil {
			return nil, v.PositionedTranslationError(child.Identifier().GetSymbol(), fmt.Sprintf("struct has no property named '%s'", propertyName))
		}

		return ast.NewASTExpressionPointerPropertyAccess(expression, property), nil
	case *parser.ParenthesizedExpressionContext:
		expression, err := v.VisitExpression(child.Expression())
		if err != nil {
			return nil, err
		}

		return ast.NewAstExpressionParenthesized(expression), nil
	case *parser.ArrayIndexExpressionContext:
		expression, err := v.VisitExpression(child.Expression(0))
		if err != nil {
			return nil, err
		}

		index, err := v.VisitExpression(child.Expression(1))
		if err != nil {
			return nil, err
		}

		return ast.NewASTExpressionIndex(expression, index), nil
	case *parser.CastExpressionContext:
		return v.VisitCastExpression(child)
	case *parser.UnaryExpressionPostContext:
		left, err := v.VisitExpression(child.Expression())
		if err != nil {
			return nil, err
		}

		return ast.NewASTExpressionUnaryPost(left, child.UnaryOperatorPost().GetText(), isStatement), nil
	case *parser.UnaryExpressionPreContext:
		right, err := v.VisitExpression(child.Expression())
		if err != nil {
			return nil, err
		}

		operator := child.UnaryOperatorPre().GetText()

		if operator == "*" {
			if right.GetType().Kind != ast.ASTTypeKindPointer {
				return nil, fmt.Errorf("invalid unary operator: %s", operator)
			}
		}

		return ast.NewASTExpressionUnaryPre(operator, right, isStatement), nil
	case *parser.SizeofExpressionContext:
		return v.VisitSizeofExpression(child)
	case *parser.AssignmentExpressionContext:
		left, err := v.VisitExpression(child.Expression(0))
		if err != nil {
			return nil, err
		}

		right, err := v.VisitExpression(child.Expression(1))
		if err != nil {
			return nil, err
		}

		converted, err := ast.NewAstTypeConversion(right, left.GetType())
		if err != nil {
			return nil, err
		}

		return ast.NewASTExpressionBinary(left, child.AssignementOperator().GetText(), converted, left.GetType()), nil
	case *parser.BinaryExpressionContext:
		left, err := v.VisitExpression(child.Expression(0))
		if err != nil {
			return nil, err
		}

		right, err := v.VisitExpression(child.Expression(1))
		if err != nil {
			return nil, err
		}

		converted, err := ast.NewAstTypeConversion(right, left.GetType())
		if err != nil {
			return nil, err
		}

		returnType := left.GetType()

		operator := child.BinaryOperator().(*parser.BinaryOperatorContext)

		if operator.LeftShift() != nil ||
			operator.RightShift() != nil ||
			operator.Less() != nil ||
			operator.Greater() != nil ||
			operator.LessEqual() != nil ||
			operator.GreaterEqual() != nil ||
			operator.Equal() != nil ||
			operator.NotEqual() != nil ||
			operator.AndAnd() != nil ||
			operator.OrOr() != nil {
			returnType = ast.NewASTType(ast.ASTTypeKindBool, "bool")
		}

		return ast.NewASTExpressionBinary(left, child.BinaryOperator().GetText(), converted, returnType), nil
	case *parser.FunctionCallExpressionContext:
		return v.VisitFunctionCallExpression(child)
	}

	return nil, v.PositionedTranslationError(ctx.GetStart(), "not implemented")
}

func (v *Visitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) (*ast.ASTExpressionLiteral, error) {
	if items := v.Scope.Get(ctx.GetText()); len(items) > 0 {
		return ast.NewASTExpressionLiteral(items[0].GetTranslatedName(), items[0].GetType()), nil
	}

	identifier := ""
	var typ *ast.ASTType = nil

	switch ctx.GetText() {
	case "printf":
		v.Imports.Add("fmt")
		identifier = "fmt.Printf"
		typ = ast.NewASTType(ast.ASTTypeKindFunction, "func").
			SetFunctionType(ast.NewASTFunction("printf").
				SetArgs([]*ast.ASTFunctionArgument{
					ast.NewASTFunctionArgument(
						"format",
						ast.NewASTType(ast.ASTTypeKindPointer, "pointer").
							SetPointerType(ast.NewASTType(ast.ASTTypeKindChar, "char")),
					),
				}).
				SetIsVariadic(true).
				SetReturnType(
					ast.NewASTType(ast.ASTTypeKindVoid, "void"),
				),
			)
	case "scanf":
		v.Imports.Add("fmt")
		identifier = "fmt.Scanf"
		typ = ast.NewASTType(ast.ASTTypeKindFunction, "func").
			SetFunctionType(ast.NewASTFunction("scanf").
				SetArgs([]*ast.ASTFunctionArgument{
					ast.NewASTFunctionArgument(
						"format",
						ast.NewASTType(ast.ASTTypeKindPointer, "pointer").
							SetPointerType(ast.NewASTType(ast.ASTTypeKindChar, "char")),
					),
				}).
				SetIsVariadic(true).
				SetReturnType(
					ast.NewASTType(ast.ASTTypeKindVoid, "void"),
				),
			)
	case "malloc":
		identifier = "libc.Malloc"
		typ = ast.NewASTType(ast.ASTTypeKindFunction, "func").
			SetFunctionType(ast.NewASTFunction("malloc").
				SetArgs([]*ast.ASTFunctionArgument{
					ast.NewASTFunctionArgument(
						"size",
						ast.NewASTType(ast.ASTTypeKindInt, "int"),
					),
				}).
				SetReturnType(
					ast.NewASTType(ast.ASTTypeKindPointer, "pointer").
						SetPointerType(ast.NewASTType(ast.ASTTypeKindVoid, "void")),
				),
			)
	case "free":
		identifier = "libc.Free"
		typ = ast.NewASTType(ast.ASTTypeKindFunction, "func").
			SetFunctionType(ast.NewASTFunction("free").
				SetArgs([]*ast.ASTFunctionArgument{
					ast.NewASTFunctionArgument(
						"ptr",
						ast.NewASTType(ast.ASTTypeKindPointer, "pointer").
							SetPointerType(ast.NewASTType(ast.ASTTypeKindVoid, "void")),
					),
				}).
				SetReturnType(
					ast.NewASTType(ast.ASTTypeKindVoid, "void"),
				),
			)
	case "assert":
		identifier = "libc.Assert"
		typ = ast.NewASTType(ast.ASTTypeKindFunction, "func").
			SetFunctionType(ast.NewASTFunction("assert").
				SetArgs([]*ast.ASTFunctionArgument{
					ast.NewASTFunctionArgument(
						"condition",
						ast.NewASTType(ast.ASTTypeKindBool, "bool"),
					),
				}).
				SetReturnType(
					ast.NewASTType(ast.ASTTypeKindVoid, "void"),
				),
			)
	}

	if identifier != "" && typ != nil {
		return ast.NewASTExpressionLiteral(identifier, typ), nil
	}

	return nil, v.PositionedTranslationError(ctx.GetStart(), fmt.Sprintf("identifier '%s' not found", ctx.GetText()))
}

func (v *Visitor) VisitCastExpression(ctx *parser.CastExpressionContext) (*ast.ASTExpressionCast, error) {
	expression, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	typ, err := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
	if err != nil {
		return nil, err
	}

	return ast.NewASTExpressionCast(expression, typ), nil
}

func (v *Visitor) VisitSizeofExpression(ctx *parser.SizeofExpressionContext) (ast.IASTExpression, error) {
	if child := ctx.Expression(); child != nil {
		expression, err := v.VisitExpression(child)
		if err != nil {
			return nil, err
		}

		return ast.NewASTExpressionCast(
			ast.NewASTExpressionFunctionCall(
				ast.NewASTExpressionLiteral("unsafe.Sizeof", ast.NewASTType(ast.ASTTypeKindAny, "any")),
				[]ast.IASTExpression{expression},
			),
			ast.NewASTType(ast.ASTTypeKindInt, "int")), nil
	}

	if child := ctx.TypeSpecifier(); child != nil {
		typ, err := v.VisitTypeSpecifier(child.(*parser.TypeSpecifierContext))
		if err != nil {
			return nil, err
		}

		var emptyExpression ast.IASTExpression

		switch typ.Kind {
		case ast.ASTTypeKindPointer:
			typ_ := ast.NewASTType(ast.ASTTypeKindPointer, "pointer").
				SetPointerType(
					ast.NewASTType(ast.ASTTypeKindVoid, "void"),
				)

			emptyExpression = ast.NewASTExpressionCast(ast.NewASTExpressionLiteral("nil", typ_), typ)
		// case ast.ASTTypeKindStruct:
		// 	emptyExpression =
		default:
			emptyExpression = ast.NewASTExpressionCast(ast.NewASTExpressionLiteral("0", ast.NewASTType(ast.ASTTypeKindInt, "int")), typ)
		}

		return ast.NewASTExpressionCast(
			ast.NewASTExpressionFunctionCall(
				ast.NewASTExpressionLiteral("unsafe.Sizeof", ast.NewASTType(ast.ASTTypeKindAny, "any")),
				[]ast.IASTExpression{emptyExpression},
			),
			ast.NewASTType(ast.ASTTypeKindInt, "int")), nil
	}

	return nil, v.PositionedTranslationError(ctx.GetStart(), "sizeof unsupported on this type")
}

func (v *Visitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) (*ast.ASTExpressionFunctionCall, error) {
	expression, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	expressionType := expression.GetType()

	if !expressionType.IsFunction() {
		return nil, v.PositionedTranslationError(ctx.GetStart(), "function call on non-function")
	}

	function := expressionType.FunctionType

	var args []ast.IASTExpression

	if child := ctx.FunctionCallArguments(); child != nil {
		args, err = v.VisitFunctionCallArguments(child.(*parser.FunctionCallArgumentsContext), function, 0)
		if err != nil {
			return nil, err
		}
	}

	if len(args) < len(function.Args) {
		return nil, v.PositionedTranslationError(ctx.GetStart(), "not enough arguments")
	}

	return ast.NewASTExpressionFunctionCall(expression, args), nil
}

func (v *Visitor) VisitFunctionCallArguments(ctx *parser.FunctionCallArgumentsContext, function *ast.ASTFunction, index int) ([]ast.IASTExpression, error) {
	if !function.IsVariadic && index >= len(function.Args) {
		return nil, v.PositionedTranslationError(ctx.GetStart(), "too many arguments")
	}

	expression, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	if index < len(function.Args) {
		expression, err = ast.NewAstTypeConversion(expression, function.Args[index].Type)
		if err != nil {
			return nil, err
		}
	}

	args := []ast.IASTExpression{expression}

	if child := ctx.FunctionCallArguments(); child != nil {
		others, err := v.VisitFunctionCallArguments(child.(*parser.FunctionCallArgumentsContext), function, index+1)
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
		return v.VisitExpressionWithConfigurableIsStatement(child, true)
	} else if child := ctx.FunctionReturn(); child != nil {
		return v.VisitFunctionReturn(child.(*parser.FunctionReturnContext))
	} else if child := ctx.Break(); child != nil {
		return ast.NewASTBreak(), nil
	} else if child := ctx.StructDeclaration(); child != nil {
		return v.VisitStructDeclaration(child.(*parser.StructDeclarationContext))
	} else if child := ctx.EnumDeclaration(); child != nil {
		return v.VisitEnumDeclaration(child.(*parser.EnumDeclarationContext))
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
	} else if child := ctx.GotoStatement(); child != nil {
		return v.VisitGotoStatement(child.(*parser.GotoStatementContext))
	} else if child := ctx.LabelStatement(); child != nil {
		return v.VisitLabelStatement(child.(*parser.LabelStatementContext))
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitGotoStatement(ctx *parser.GotoStatementContext) (*ast.ASTGoto, error) {
	return ast.NewASTGoto(ctx.Identifier().GetText()), nil
}

func (v *Visitor) VisitLabelStatement(ctx *parser.LabelStatementContext) (*ast.ASTLabel, error) {
	return ast.NewASTLabel(ctx.Identifier().GetText()), nil
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
		expression, err := v.VisitExpressionWithConfigurableIsStatement(child, true)
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

		if !libc.IsSupported(header) {
			return nil, v.PositionedTranslationError(ctx.IncludeDirective().GetSymbol(), "unsupported system header: "+header)
		}

		v.Imports.Add("github.com/tereus-project/tereus-remixer-c-go/libc")
	} else if matches := localIncludePreprocessorRegex.FindStringSubmatch(directive); len(matches) > 0 {
		return nil, v.TranslationError(ctx.BaseParserRuleContext, "unsupported include type")
	} else {
		return nil, v.TranslationError(ctx.BaseParserRuleContext, "unsupported include type")
	}

	return nil, nil
}
