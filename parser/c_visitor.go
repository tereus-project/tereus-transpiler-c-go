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

	// Visit a parse tree produced by CParser#functionReturn.
	VisitFunctionReturn(ctx *FunctionReturnContext) interface{}

	// Visit a parse tree produced by CParser#typeSpecifier.
	VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{}

	// Visit a parse tree produced by CParser#structDeclaration.
	VisitStructDeclaration(ctx *StructDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#structProperty.
	VisitStructProperty(ctx *StructPropertyContext) interface{}

	// Visit a parse tree produced by CParser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#enumProperties.
	VisitEnumProperties(ctx *EnumPropertiesContext) interface{}

	// Visit a parse tree produced by CParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by CParser#variableDeclarationList.
	VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{}

	// Visit a parse tree produced by CParser#ParenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by CParser#SizeofExpression.
	VisitSizeofExpression(ctx *SizeofExpressionContext) interface{}

	// Visit a parse tree produced by CParser#ArrayIndexExpression.
	VisitArrayIndexExpression(ctx *ArrayIndexExpressionContext) interface{}

	// Visit a parse tree produced by CParser#UnaryExpressionPost.
	VisitUnaryExpressionPost(ctx *UnaryExpressionPostContext) interface{}

	// Visit a parse tree produced by CParser#AssignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by CParser#BinaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by CParser#ConstantExpression.
	VisitConstantExpression(ctx *ConstantExpressionContext) interface{}

	// Visit a parse tree produced by CParser#UnaryExpressionPre.
	VisitUnaryExpressionPre(ctx *UnaryExpressionPreContext) interface{}

	// Visit a parse tree produced by CParser#CastExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by CParser#ConstantStringExpression.
	VisitConstantStringExpression(ctx *ConstantStringExpressionContext) interface{}

	// Visit a parse tree produced by CParser#FunctionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by CParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by CParser#assignementOperator.
	VisitAssignementOperator(ctx *AssignementOperatorContext) interface{}

	// Visit a parse tree produced by CParser#binaryOperator.
	VisitBinaryOperator(ctx *BinaryOperatorContext) interface{}

	// Visit a parse tree produced by CParser#unaryOperatorPost.
	VisitUnaryOperatorPost(ctx *UnaryOperatorPostContext) interface{}

	// Visit a parse tree produced by CParser#unaryOperatorPre.
	VisitUnaryOperatorPre(ctx *UnaryOperatorPreContext) interface{}

	// Visit a parse tree produced by CParser#functionCallArguments.
	VisitFunctionCallArguments(ctx *FunctionCallArgumentsContext) interface{}

	// Visit a parse tree produced by CParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by CParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by CParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by CParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by CParser#includePreprocessor.
	VisitIncludePreprocessor(ctx *IncludePreprocessorContext) interface{}
}
