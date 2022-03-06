package remixer

import (
	"github.com/tereus-project/tereus-remixer-c-go/parser"
)

type Visitor struct{}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{}

func (v *Visitor) VisitGenericSelection(ctx *parser.GenericSelectionContext) interface{}

func (v *Visitor) VisitGenericAssocList(ctx *parser.GenericAssocListContext) interface{}

func (v *Visitor) VisitGenericAssociation(ctx *parser.GenericAssociationContext) interface{}

func (v *Visitor) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{}

func (v *Visitor) VisitArgumentExpressionList(ctx *parser.ArgumentExpressionListContext) interface{}

func (v *Visitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{}

func (v *Visitor) VisitUnaryOperator(ctx *parser.UnaryOperatorContext) interface{}

func (v *Visitor) VisitCastExpression(ctx *parser.CastExpressionContext) interface{}

func (v *Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{}

func (v *Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{}

func (v *Visitor) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{}

func (v *Visitor) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{}

func (v *Visitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{}

func (v *Visitor) VisitAndExpression(ctx *parser.AndExpressionContext) interface{}

func (v *Visitor) VisitExclusiveOrExpression(ctx *parser.ExclusiveOrExpressionContext) interface{}

func (v *Visitor) VisitInclusiveOrExpression(ctx *parser.InclusiveOrExpressionContext) interface{}

func (v *Visitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{}

func (v *Visitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{}

func (v *Visitor) VisitConditionalExpression(ctx *parser.ConditionalExpressionContext) interface{}

func (v *Visitor) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) interface{}

func (v *Visitor) VisitAssignmentOperator(ctx *parser.AssignmentOperatorContext) interface{}

func (v *Visitor) VisitExpression(ctx *parser.ExpressionContext) interface{}

func (v *Visitor) VisitConstantExpression(ctx *parser.ConstantExpressionContext) interface{}

func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{}

func (v *Visitor) VisitDeclarationSpecifiers(ctx *parser.DeclarationSpecifiersContext) interface{}

func (v *Visitor) VisitDeclarationSpecifiers2(ctx *parser.DeclarationSpecifiers2Context) interface{}

func (v *Visitor) VisitDeclarationSpecifier(ctx *parser.DeclarationSpecifierContext) interface{}

func (v *Visitor) VisitInitDeclaratorList(ctx *parser.InitDeclaratorListContext) interface{}

func (v *Visitor) VisitInitDeclarator(ctx *parser.InitDeclaratorContext) interface{}

func (v *Visitor) VisitStorageClassSpecifier(ctx *parser.StorageClassSpecifierContext) interface{}

func (v *Visitor) VisitTypeSpecifier(ctx *parser.TypeSpecifierContext) interface{}

func (v *Visitor) VisitStructOrUnionSpecifier(ctx *parser.StructOrUnionSpecifierContext) interface{}

func (v *Visitor) VisitStructOrUnion(ctx *parser.StructOrUnionContext) interface{}

func (v *Visitor) VisitStructDeclarationList(ctx *parser.StructDeclarationListContext) interface{}

func (v *Visitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{}

func (v *Visitor) VisitSpecifierQualifierList(ctx *parser.SpecifierQualifierListContext) interface{}

func (v *Visitor) VisitStructDeclaratorList(ctx *parser.StructDeclaratorListContext) interface{}

func (v *Visitor) VisitStructDeclarator(ctx *parser.StructDeclaratorContext) interface{}

func (v *Visitor) VisitEnumSpecifier(ctx *parser.EnumSpecifierContext) interface{}

func (v *Visitor) VisitEnumeratorList(ctx *parser.EnumeratorListContext) interface{}

func (v *Visitor) VisitEnumerator(ctx *parser.EnumeratorContext) interface{}

func (v *Visitor) VisitEnumerationConstant(ctx *parser.EnumerationConstantContext) interface{}

func (v *Visitor) VisitAtomicTypeSpecifier(ctx *parser.AtomicTypeSpecifierContext) interface{}

func (v *Visitor) VisitTypeQualifier(ctx *parser.TypeQualifierContext) interface{}

func (v *Visitor) VisitFunctionSpecifier(ctx *parser.FunctionSpecifierContext) interface{}

func (v *Visitor) VisitAlignmentSpecifier(ctx *parser.AlignmentSpecifierContext) interface{}

func (v *Visitor) VisitDeclarator(ctx *parser.DeclaratorContext) interface{}

func (v *Visitor) VisitDirectDeclarator(ctx *parser.DirectDeclaratorContext) interface{}

func (v *Visitor) VisitGccDeclaratorExtension(ctx *parser.GccDeclaratorExtensionContext) interface{}

func (v *Visitor) VisitGccAttributeSpecifier(ctx *parser.GccAttributeSpecifierContext) interface{}

func (v *Visitor) VisitGccAttributeList(ctx *parser.GccAttributeListContext) interface{}

func (v *Visitor) VisitGccAttribute(ctx *parser.GccAttributeContext) interface{}

func (v *Visitor) VisitNestedParenthesesBlock(ctx *parser.NestedParenthesesBlockContext) interface{}

func (v *Visitor) VisitPointer(ctx *parser.PointerContext) interface{}

func (v *Visitor) VisitTypeQualifierList(ctx *parser.TypeQualifierListContext) interface{}

func (v *Visitor) VisitParameterTypeList(ctx *parser.ParameterTypeListContext) interface{}

func (v *Visitor) VisitParameterList(ctx *parser.ParameterListContext) interface{}

func (v *Visitor) VisitParameterDeclaration(ctx *parser.ParameterDeclarationContext) interface{}

func (v *Visitor) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{}

func (v *Visitor) VisitTypeName(ctx *parser.TypeNameContext) interface{}

func (v *Visitor) VisitAbstractDeclarator(ctx *parser.AbstractDeclaratorContext) interface{}

func (v *Visitor) VisitDirectAbstractDeclarator(ctx *parser.DirectAbstractDeclaratorContext) interface{}

func (v *Visitor) VisitTypedefName(ctx *parser.TypedefNameContext) interface{}

func (v *Visitor) VisitInitializer(ctx *parser.InitializerContext) interface{}

func (v *Visitor) VisitInitializerList(ctx *parser.InitializerListContext) interface{}

func (v *Visitor) VisitDesignation(ctx *parser.DesignationContext) interface{}

func (v *Visitor) VisitDesignatorList(ctx *parser.DesignatorListContext) interface{}

func (v *Visitor) VisitDesignator(ctx *parser.DesignatorContext) interface{}

func (v *Visitor) VisitStaticAssertDeclaration(ctx *parser.StaticAssertDeclarationContext) interface{}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{}

func (v *Visitor) VisitLabeledStatement(ctx *parser.LabeledStatementContext) interface{}

func (v *Visitor) VisitCompoundStatement(ctx *parser.CompoundStatementContext) interface{}

func (v *Visitor) VisitBlockItemList(ctx *parser.BlockItemListContext) interface{}

func (v *Visitor) VisitBlockItem(ctx *parser.BlockItemContext) interface{}

func (v *Visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{}

func (v *Visitor) VisitSelectionStatement(ctx *parser.SelectionStatementContext) interface{}

func (v *Visitor) VisitIterationStatement(ctx *parser.IterationStatementContext) interface{}

func (v *Visitor) VisitForCondition(ctx *parser.ForConditionContext) interface{}

func (v *Visitor) VisitForDeclaration(ctx *parser.ForDeclarationContext) interface{}

func (v *Visitor) VisitForExpression(ctx *parser.ForExpressionContext) interface{}

func (v *Visitor) VisitJumpStatement(ctx *parser.JumpStatementContext) interface{}

func (v *Visitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{}

func (v *Visitor) VisitTranslationUnit(ctx *parser.TranslationUnitContext) interface{}

func (v *Visitor) VisitExternalDeclaration(ctx *parser.ExternalDeclarationContext) interface{}

func (v *Visitor) VisitFunctionDefinition(ctx *parser.FunctionDefinitionContext) interface{}

func (v *Visitor) VisitDeclarationList(ctx *parser.DeclarationListContext) interface{}
