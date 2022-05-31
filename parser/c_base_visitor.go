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

func (v *BaseCVisitor) VisitFunctionReturn(ctx *FunctionReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitStructDeclarationBody(ctx *StructDeclarationBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitStructProperty(ctx *StructPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitEnumProperties(ctx *EnumPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitConstantExpression(ctx *ConstantExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitConstantStringExpression(ctx *ConstantStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitSizeofExpression(ctx *SizeofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitArrayIndexExpression(ctx *ArrayIndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitUnaryExpressionPost(ctx *UnaryExpressionPostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitUnaryExpressionPre(ctx *UnaryExpressionPreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitPropertyAccessExpression(ctx *PropertyAccessExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitCastExpression(ctx *CastExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAssignementOperator(ctx *AssignementOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitBinaryOperator(ctx *BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitUnaryOperatorPost(ctx *UnaryOperatorPostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitUnaryOperatorPre(ctx *UnaryOperatorPreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitFunctionCallArguments(ctx *FunctionCallArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitIncludePreprocessor(ctx *IncludePreprocessorContext) interface{} {
	return v.VisitChildren(ctx)
}
