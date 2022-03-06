// Code generated from C.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // C

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CParser.
type CVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CParser#translation.
	VisitTranslation(ctx *TranslationContext) interface{}

	// Visit a parse tree produced by CParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by CParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#functionArguments.
	VisitFunctionArguments(ctx *FunctionArgumentsContext) interface{}

	// Visit a parse tree produced by CParser#typeSpecifier.
	VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by CParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by CParser#statement.
	VisitStatement(ctx *StatementContext) interface{}
}
