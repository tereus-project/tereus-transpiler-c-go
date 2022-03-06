package remixer

import (
	"github.com/tereus-project/tereus-remixer-c-go/parser"
)

type Visitor struct{}

func NewVisitor() *Visitor {
	return &Visitor{}
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

func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) {}

func (v *Visitor) VisitDeclarationSpecifiers(ctx *parser.DeclarationSpecifiersContext) {}

func (v *Visitor) VisitDeclarationSpecifiers2(ctx *parser.DeclarationSpecifiers2Context) {}

func (v *Visitor) VisitDeclarationSpecifier(ctx *parser.DeclarationSpecifierContext) {}

func (v *Visitor) VisitInitDeclaratorList(ctx *parser.InitDeclaratorListContext) {}

func (v *Visitor) VisitInitDeclarator(ctx *parser.InitDeclaratorContext) {}

func (v *Visitor) VisitStorageClassSpecifier(ctx *parser.StorageClassSpecifierContext) {}

func (v *Visitor) VisitTypeSpecifier(ctx *parser.TypeSpecifierContext) {}

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

func (v *Visitor) VisitDeclarator(ctx *parser.DeclaratorContext) {}

func (v *Visitor) VisitDirectDeclarator(ctx *parser.DirectDeclaratorContext) {}

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

func (v *Visitor) VisitIdentifierList(ctx *parser.IdentifierListContext) {}

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

func (v *Visitor) VisitCompoundStatement(ctx *parser.CompoundStatementContext) {}

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

func (v *Visitor) VisitTranslationUnit(ctx *parser.TranslationUnitContext) {}

func (v *Visitor) VisitExternalDeclaration(ctx *parser.ExternalDeclarationContext) {}

func (v *Visitor) VisitFunctionDefinition(ctx *parser.FunctionDefinitionContext) {}

func (v *Visitor) VisitDeclarationList(ctx *parser.DeclarationListContext) {}
