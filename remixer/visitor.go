package remixer

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tereus-project/tereus-remixer-c-go/parser"
)

type Visitor struct {
	Path   string
	Code   string
	Output map[string]string
}

func NewVisitor(path string) (*Visitor, error) {
	code, e := os.ReadFile(path)
	if e != nil {
		return nil, e
	}

	visitor := &Visitor{
		Path:   path,
		Code:   string(code),
		Output: make(map[string]string, 0),
	}

	return visitor, nil
}

func (v *Visitor) NotImplementedError(token *antlr.BaseParserRuleContext) error {
	start := token.GetStart()
	return fmt.Errorf("%s:%d:%d: %s", v.Path, start.GetLine(), start.GetColumn(), "not implemented")
}

func (v *Visitor) VisitTranslation(ctx *parser.TranslationContext) (string, error) {
	output := ""

	for _, declaration := range ctx.AllDeclaration() {
		declaration, e := v.VisitDeclaration(declaration.(*parser.DeclarationContext))
		if e != nil {
			return "", e
		}

		output += declaration
	}

	return output, nil
}

func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) (string, error) {
	if child := ctx.FunctionDeclaration(); child != nil {
		return v.VisitFunctionDeclaration(child.(*parser.FunctionDeclarationContext))
	}

	return "", v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) (string, error) {
	var e error

	name := ctx.Identifier().GetText()
	typ := ""

	if name != "main" {
		typ, e = v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
		if e != nil {
			return "", e
		}
	}

	var args string

	if child := ctx.FunctionArguments(); child != nil {
		args, e = v.VisitFunctionArguments(child.(*parser.FunctionArgumentsContext))
		if e != nil {
			return "", e
		}
	}

	body, e := v.VisitBlock(ctx.Block().(*parser.BlockContext))
	if e != nil {
		return "", e
	}

	if typ != "" {
		typ = " " + typ
	}

	return fmt.Sprintf("func %s(%s)%s %s", name, args, typ, body), nil
}

func (v *Visitor) VisitFunctionArguments(ctx *parser.FunctionArgumentsContext) (string, error) {
	name := ctx.Identifier().GetText()

	typ, e := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
	if e != nil {
		return "", e
	}

	if ctx.FunctionArguments() != nil {
		args, e := v.VisitFunctionArguments(ctx.FunctionArguments().(*parser.FunctionArgumentsContext))
		if e != nil {
			return "", e
		}

		return fmt.Sprintf("%s %s, %s", name, typ, args), nil
	}

	return fmt.Sprintf("%s %s", name, typ), nil
}

func (v *Visitor) VisitTypeSpecifier(ctx *parser.TypeSpecifierContext) (string, error) {
	if child := ctx.Void(); child != nil {
		return "", nil
	} else if child := ctx.Int(); child != nil {
		return "int", nil
	} else if child := ctx.Short(); child != nil {
		return "int16", nil
	} else if child := ctx.Long(); child != nil {
		return "int64", nil
	} else if child := ctx.Char(); child != nil {
		return "rune", nil
	} else if child := ctx.Float(); child != nil {
		return "float32", nil
	} else if child := ctx.Double(); child != nil {
		return "float64", nil
	}

	return "", v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) (string, error) {
	typ, e := v.VisitTypeSpecifier(ctx.TypeSpecifier().(*parser.TypeSpecifierContext))
	if e != nil {
		return "", e
	}

	name := ctx.Identifier().GetText()

	if child := ctx.Expression(); child != nil {
		expression, e := v.VisitExpression(child.(*parser.ExpressionContext))
		if e != nil {
			return "", e
		}

		return fmt.Sprintf("%s := %s", name, expression), nil
	}

	return fmt.Sprintf("var %s %s", name, typ), nil
}

func (v *Visitor) VisitExpression(ctx *parser.ExpressionContext) (string, error) {
	if child := ctx.Constant(); child != nil {
		return child.GetText(), nil
	}

	return "", v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) (string, error) {
	output := ""

	for _, statement := range ctx.AllStatement() {
		statement, e := v.VisitStatement(statement.(*parser.StatementContext))
		if e != nil {
			return "", e
		}

		output += statement
	}

	lines := strings.Split(output, "\n")

	for i, line := range lines {
		lines[i] = "    " + line
	}

	return "{\n" + strings.Join(lines, "\n") + "\n}", nil
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) (string, error) {
	if child := ctx.VariableDeclaration(); child != nil {
		return v.VisitVariableDeclaration(child.(*parser.VariableDeclarationContext))
	}

	return "", v.NotImplementedError(ctx.BaseParserRuleContext)
}
