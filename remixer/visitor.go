package remixer

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tereus-project/tereus-remixer-c-go/parser"
	"github.com/tereus-project/tereus-remixer-c-go/remixer/language"
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

func (v *Visitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) {}

func (v *Visitor) VisitGenericSelection(ctx *parser.GenericSelectionContext) {}

func (v *Visitor) VisitGenericAssocList(ctx *parser.GenericAssocListContext) {}

func (v *Visitor) VisitGenericAssociation(ctx *parser.GenericAssociationContext) {}

func (v *Visitor) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) {}

func (v *Visitor) VisitArgumentExpressionList(ctx *parser.ArgumentExpressionListContext) {}

func (v *Visitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) {}

func (v *Visitor) VisitUnaryOperator(ctx *parser.UnaryOperatorContext) {}

func (v *Visitor) VisitCastExpression(ctx *parser.CastExpressionContext) {}

func (v *Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) {}

func (v *Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) {}

func (v *Visitor) VisitShiftExpression(ctx *parser.ShiftExpressionContext) {}

func (v *Visitor) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) {}

func (v *Visitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) {}

func (v *Visitor) VisitAndExpression(ctx *parser.AndExpressionContext) {}

func (v *Visitor) VisitExclusiveOrExpression(ctx *parser.ExclusiveOrExpressionContext) {}

func (v *Visitor) VisitInclusiveOrExpression(ctx *parser.InclusiveOrExpressionContext) {}

func (v *Visitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) {}

func (v *Visitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) {}

func (v *Visitor) VisitConditionalExpression(ctx *parser.ConditionalExpressionContext) {}

func (v *Visitor) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) {}

func (v *Visitor) VisitAssignmentOperator(ctx *parser.AssignmentOperatorContext) {}

func (v *Visitor) VisitExpression(ctx *parser.ExpressionContext) {}

func (v *Visitor) VisitConstantExpression(ctx *parser.ConstantExpressionContext) {}

func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) (string, error) {
	return "", v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitDeclarationSpecifiers(ctx *parser.DeclarationSpecifiersContext) (*language.Specifiers, error) {
	specifiers := language.NewSpecifiers()

	for _, specifier := range ctx.AllDeclarationSpecifier() {
		specifier, e := v.VisitDeclarationSpecifier(specifier.(*parser.DeclarationSpecifierContext))
		if e != nil {
			return nil, e
		}

		switch specifier.(type) {
		case *language.SpecifierType:
			specifiers.Type = specifier.(*language.SpecifierType)
		default:
			return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
		}
	}

	return specifiers, nil
}

func (v *Visitor) VisitDeclarationSpecifiers2(ctx *parser.DeclarationSpecifiers2Context) {}

func (v *Visitor) VisitDeclarationSpecifier(ctx *parser.DeclarationSpecifierContext) (language.ISpecifier, error) {
	ctx.GetStart()
	if child := ctx.TypeSpecifier(); child != nil {
		return v.VisitTypeSpecifier(child.(*parser.TypeSpecifierContext))
	}

	return "", v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitInitDeclaratorList(ctx *parser.InitDeclaratorListContext) {}

func (v *Visitor) VisitInitDeclarator(ctx *parser.InitDeclaratorContext) {}

func (v *Visitor) VisitStorageClassSpecifier(ctx *parser.StorageClassSpecifierContext) {}

func (v *Visitor) VisitTypeSpecifier(ctx *parser.TypeSpecifierContext) (*language.SpecifierType, error) {
	if child := ctx.Void(); child != nil {
		return language.NewSpecifierType("void"), nil
	} else if child := ctx.Char(); child != nil {
		return language.NewSpecifierType("char"), nil
	} else if child := ctx.Short(); child != nil {
		return language.NewSpecifierType("short"), nil
	} else if child := ctx.Int(); child != nil {
		return language.NewSpecifierType("int"), nil
	} else if child := ctx.Long(); child != nil {
		return language.NewSpecifierType("long"), nil
	} else if child := ctx.Float(); child != nil {
		return language.NewSpecifierType("float"), nil
	} else if child := ctx.Double(); child != nil {
		return language.NewSpecifierType("double"), nil
	} else if child := ctx.Signed(); child != nil {
		return language.NewSpecifierType("signed"), nil
	} else if child := ctx.Unsigned(); child != nil {
		return language.NewSpecifierType("unsigned"), nil
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitStructOrUnionSpecifier(ctx *parser.StructOrUnionSpecifierContext) {}

func (v *Visitor) VisitStructOrUnion(ctx *parser.StructOrUnionContext) {}

func (v *Visitor) VisitStructDeclarationList(ctx *parser.StructDeclarationListContext) {}

func (v *Visitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) {}

func (v *Visitor) VisitSpecifierQualifierList(ctx *parser.SpecifierQualifierListContext) {}

func (v *Visitor) VisitStructDeclaratorList(ctx *parser.StructDeclaratorListContext) {}

func (v *Visitor) VisitStructDeclarator(ctx *parser.StructDeclaratorContext) {}

func (v *Visitor) VisitEnumSpecifier(ctx *parser.EnumSpecifierContext) {}

func (v *Visitor) VisitEnumeratorList(ctx *parser.EnumeratorListContext) {}

func (v *Visitor) VisitEnumerator(ctx *parser.EnumeratorContext) {}

func (v *Visitor) VisitEnumerationConstant(ctx *parser.EnumerationConstantContext) {}

func (v *Visitor) VisitAtomicTypeSpecifier(ctx *parser.AtomicTypeSpecifierContext) {}

func (v *Visitor) VisitTypeQualifier(ctx *parser.TypeQualifierContext) {}

func (v *Visitor) VisitFunctionSpecifier(ctx *parser.FunctionSpecifierContext) {}

func (v *Visitor) VisitAlignmentSpecifier(ctx *parser.AlignmentSpecifierContext) {}

func (v *Visitor) VisitDeclarator(ctx *parser.DeclaratorContext) (language.IDeclarator, error) {
	if child := ctx.Pointer(); child != nil {
		return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
	}

	directDeclarator, e := v.VisitDirectDeclarator(ctx.DirectDeclarator().(*parser.DirectDeclaratorContext))
	if e != nil {
		return nil, e
	}

	return directDeclarator, nil
}

func (v *Visitor) VisitDirectDeclarator(ctx *parser.DirectDeclaratorContext) (language.IDeclarator, error) {
	if child := ctx.Identifier(); child != nil {
		return language.NewDeclarator(child.GetText()), nil
	} else if child := ctx.DirectDeclarator(); child != nil {
		directDeclarator, e := v.VisitDirectDeclarator(child.(*parser.DirectDeclaratorContext))
		if e != nil {
			return nil, e
		}

		// Functions
		if child := ctx.LeftParen(); child != nil {
			if child := ctx.IdentifierList(); child != nil {
				identifiers := v.VisitIdentifierList(child.(*parser.IdentifierListContext))
				return language.NewDeclaratorFunction(directDeclarator, identifiers), nil
			}

			return language.NewDeclaratorFunction(directDeclarator, make([]language.IDeclarator, 0)), nil
		}

		return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
	}

	return nil, v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitGccDeclaratorExtension(ctx *parser.GccDeclaratorExtensionContext) {}

func (v *Visitor) VisitGccAttributeSpecifier(ctx *parser.GccAttributeSpecifierContext) {}

func (v *Visitor) VisitGccAttributeList(ctx *parser.GccAttributeListContext) {}

func (v *Visitor) VisitGccAttribute(ctx *parser.GccAttributeContext) {}

func (v *Visitor) VisitNestedParenthesesBlock(ctx *parser.NestedParenthesesBlockContext) {}

func (v *Visitor) VisitPointer(ctx *parser.PointerContext) {}

func (v *Visitor) VisitTypeQualifierList(ctx *parser.TypeQualifierListContext) {}

func (v *Visitor) VisitParameterTypeList(ctx *parser.ParameterTypeListContext) {}

func (v *Visitor) VisitParameterList(ctx *parser.ParameterListContext) {}

func (v *Visitor) VisitParameterDeclaration(ctx *parser.ParameterDeclarationContext) {}

func (v *Visitor) VisitIdentifierList(ctx *parser.IdentifierListContext) []language.IDeclarator {
	identifiers := make([]language.IDeclarator, 0)

	for _, child := range ctx.AllIdentifier() {
		identifiers = append(identifiers, language.NewDeclarator(child.GetText()))
	}

	return identifiers
}

func (v *Visitor) VisitTypeName(ctx *parser.TypeNameContext) {}

func (v *Visitor) VisitAbstractDeclarator(ctx *parser.AbstractDeclaratorContext) {}

func (v *Visitor) VisitDirectAbstractDeclarator(ctx *parser.DirectAbstractDeclaratorContext) {}

func (v *Visitor) VisitTypedefName(ctx *parser.TypedefNameContext) {}

func (v *Visitor) VisitInitializer(ctx *parser.InitializerContext) {}

func (v *Visitor) VisitInitializerList(ctx *parser.InitializerListContext) {}

func (v *Visitor) VisitDesignation(ctx *parser.DesignationContext) {}

func (v *Visitor) VisitDesignatorList(ctx *parser.DesignatorListContext) {}

func (v *Visitor) VisitDesignator(ctx *parser.DesignatorContext) {}

func (v *Visitor) VisitStaticAssertDeclaration(ctx *parser.StaticAssertDeclarationContext) {}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) {}

func (v *Visitor) VisitLabeledStatement(ctx *parser.LabeledStatementContext) {}

func (v *Visitor) VisitCompoundStatement(ctx *parser.CompoundStatementContext) (string, error) {
	if child := ctx.BlockItemList(); child != nil {
		return "", v.NotImplementedError(ctx.BaseParserRuleContext)
	}

	return "{" + "}", nil
}

func (v *Visitor) VisitBlockItemList(ctx *parser.BlockItemListContext) {}

func (v *Visitor) VisitBlockItem(ctx *parser.BlockItemContext) {}

func (v *Visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) {}

func (v *Visitor) VisitSelectionStatement(ctx *parser.SelectionStatementContext) {}

func (v *Visitor) VisitIterationStatement(ctx *parser.IterationStatementContext) {}

func (v *Visitor) VisitForCondition(ctx *parser.ForConditionContext) {}

func (v *Visitor) VisitForDeclaration(ctx *parser.ForDeclarationContext) {}

func (v *Visitor) VisitForExpression(ctx *parser.ForExpressionContext) {}

func (v *Visitor) VisitJumpStatement(ctx *parser.JumpStatementContext) {}

func (v *Visitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) {}

func (v *Visitor) VisitTranslationUnit(ctx *parser.TranslationUnitContext) (string, error) {
	output := ""

	for _, externalDeclaration := range ctx.AllExternalDeclaration() {
		externalDeclarationContext := externalDeclaration.(*parser.ExternalDeclarationContext)
		code, e := v.VisitExternalDeclaration(externalDeclarationContext)
		if e != nil {
			return "", e
		}

		output += code
	}

	return output, nil
}

func (v *Visitor) VisitExternalDeclaration(ctx *parser.ExternalDeclarationContext) (string, error) {
	if child := ctx.FunctionDefinition(); child != nil {
		return v.VisitFunctionDefinition(child.(*parser.FunctionDefinitionContext))
	}

	return "", v.NotImplementedError(ctx.BaseParserRuleContext)
}

func (v *Visitor) VisitFunctionDefinition(ctx *parser.FunctionDefinitionContext) (string, error) {
	var specifiers *language.Specifiers

	if child := ctx.DeclarationSpecifiers(); child != nil {
		var e error
		specifiers, e = v.VisitDeclarationSpecifiers(child.(*parser.DeclarationSpecifiersContext))
		if e != nil {
			return "", e
		}
	} else {
		specifiers = language.NewSpecifiers()
	}

	declarator, e := v.VisitDeclarator(ctx.Declarator().(*parser.DeclaratorContext))
	if e != nil {
		return "", e
	}

	if child := ctx.DeclarationList(); child != nil {
		return "", v.NotImplementedError(ctx.BaseParserRuleContext)
	}

	body := ""
	if child := ctx.CompoundStatement(); child != nil {
		body, e = v.VisitCompoundStatement(child.(*parser.CompoundStatementContext))
		if e != nil {
			return "", e
		}

		body = " " + body
	}

	returnType := ""

	if declarator.(*language.DeclaratorFunction).Identifier() != "main" {
		returnType = specifiers.String()

		if returnType != "" {
			returnType = " " + returnType
		}
	}

	return fmt.Sprintf("func %s%s%s", declarator.String(), returnType, body), nil
}

func (v *Visitor) VisitDeclarationList(ctx *parser.DeclarationListContext) (string, error) {
	output := ""

	for _, declaration := range ctx.AllDeclaration() {
		code, e := v.VisitDeclaration(declaration.(*parser.DeclarationContext))
		if e != nil {
			return "", e
		}

		output += code
	}

	return output, nil
}
