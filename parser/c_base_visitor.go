// Code generated from C.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // C

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCVisitor) VisitTranslation(ctx *TranslationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitFunctionArguments(ctx *FunctionArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}
