// Code generated from C.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // C

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 91, 67, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 7, 2, 22, 10, 2, 12, 2, 14, 2,
	25, 11, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 33, 10, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 42, 10, 5, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 5, 7, 50, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 7,
	9, 58, 10, 9, 12, 9, 14, 9, 61, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 4, 8, 2, 5, 5, 10, 10, 14,
	14, 18, 19, 21, 21, 30, 30, 3, 2, 78, 79, 2, 62, 2, 23, 3, 2, 2, 2, 4,
	26, 3, 2, 2, 2, 6, 28, 3, 2, 2, 2, 8, 37, 3, 2, 2, 2, 10, 43, 3, 2, 2,
	2, 12, 45, 3, 2, 2, 2, 14, 53, 3, 2, 2, 2, 16, 55, 3, 2, 2, 2, 18, 64,
	3, 2, 2, 2, 20, 22, 5, 4, 3, 2, 21, 20, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2,
	23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 3, 3, 2, 2, 2, 25, 23, 3, 2,
	2, 2, 26, 27, 5, 6, 4, 2, 27, 5, 3, 2, 2, 2, 28, 29, 5, 10, 6, 2, 29, 30,
	7, 78, 2, 2, 30, 32, 7, 32, 2, 2, 31, 33, 5, 8, 5, 2, 32, 31, 3, 2, 2,
	2, 32, 33, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 35, 7, 33, 2, 2, 35, 36,
	5, 16, 9, 2, 36, 7, 3, 2, 2, 2, 37, 38, 5, 10, 6, 2, 38, 41, 7, 78, 2,
	2, 39, 40, 7, 61, 2, 2, 40, 42, 5, 8, 5, 2, 41, 39, 3, 2, 2, 2, 41, 42,
	3, 2, 2, 2, 42, 9, 3, 2, 2, 2, 43, 44, 9, 2, 2, 2, 44, 11, 3, 2, 2, 2,
	45, 46, 5, 10, 6, 2, 46, 49, 7, 78, 2, 2, 47, 48, 7, 62, 2, 2, 48, 50,
	5, 14, 8, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2,
	51, 52, 7, 60, 2, 2, 52, 13, 3, 2, 2, 2, 53, 54, 9, 3, 2, 2, 54, 15, 3,
	2, 2, 2, 55, 59, 7, 36, 2, 2, 56, 58, 5, 18, 10, 2, 57, 56, 3, 2, 2, 2,
	58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 3,
	2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63, 7, 37, 2, 2, 63, 17, 3, 2, 2, 2, 64,
	65, 5, 12, 7, 2, 65, 19, 3, 2, 2, 2, 7, 23, 32, 41, 49, 59,
}
var literalNames = []string{
	"", "'break'", "'case'", "'char'", "'const'", "'continue'", "'default'",
	"'do'", "'double'", "'else'", "'enum'", "'extern'", "'float'", "'for'",
	"'goto'", "'if'", "'int'", "'long'", "'return'", "'short'", "'signed'",
	"'sizeof'", "'static'", "'struct'", "'switch'", "'typedef'", "'union'",
	"'unsigned'", "'void'", "'while'", "'('", "')'", "'['", "']'", "'{'", "'}'",
	"'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", "'+'", "'++'", "'-'", "'--'",
	"'*'", "'/'", "'%'", "'&'", "'|'", "'&&'", "'||'", "'^'", "'!'", "'~'",
	"'?'", "':'", "';'", "','", "'='", "'*='", "'/='", "'%='", "'+='", "'-='",
	"'<<='", "'>>='", "'&='", "'^='", "'|='", "'=='", "'!='", "'->'", "'.'",
	"'...'",
}
var symbolicNames = []string{
	"", "Break", "Case", "Char", "Const", "Continue", "Default", "Do", "Double",
	"Else", "Enum", "Extern", "Float", "For", "Goto", "If", "Int", "Long",
	"Return", "Short", "Signed", "Sizeof", "Static", "Struct", "Switch", "Typedef",
	"Union", "Unsigned", "Void", "While", "LeftParen", "RightParen", "LeftBracket",
	"RightBracket", "LeftBrace", "RightBrace", "Less", "LessEqual", "Greater",
	"GreaterEqual", "LeftShift", "RightShift", "Plus", "PlusPlus", "Minus",
	"MinusMinus", "Star", "Div", "Mod", "And", "Or", "AndAnd", "OrOr", "Caret",
	"Not", "Tilde", "Question", "Colon", "Semi", "Comma", "Assign", "StarAssign",
	"DivAssign", "ModAssign", "PlusAssign", "MinusAssign", "LeftShiftAssign",
	"RightShiftAssign", "AndAssign", "XorAssign", "OrAssign", "Equal", "NotEqual",
	"Arrow", "Dot", "Ellipsis", "Identifier", "Constant", "DigitSequence",
	"StringLiteral", "ComplexDefine", "IncludeDirective", "AsmBlock", "LineAfterPreprocessing",
	"LineDirective", "PragmaDirective", "Whitespace", "Newline", "BlockComment",
	"LineComment",
}

var ruleNames = []string{
	"translation", "declaration", "functionDeclaration", "functionArguments",
	"typeSpecifier", "variableDeclaration", "expression", "block", "statement",
}

type CParser struct {
	*antlr.BaseParser
}

// NewCParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *CParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCParser(input antlr.TokenStream) *CParser {
	this := new(CParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "C.g4"

	return this
}

// CParser tokens.
const (
	CParserEOF                    = antlr.TokenEOF
	CParserBreak                  = 1
	CParserCase                   = 2
	CParserChar                   = 3
	CParserConst                  = 4
	CParserContinue               = 5
	CParserDefault                = 6
	CParserDo                     = 7
	CParserDouble                 = 8
	CParserElse                   = 9
	CParserEnum                   = 10
	CParserExtern                 = 11
	CParserFloat                  = 12
	CParserFor                    = 13
	CParserGoto                   = 14
	CParserIf                     = 15
	CParserInt                    = 16
	CParserLong                   = 17
	CParserReturn                 = 18
	CParserShort                  = 19
	CParserSigned                 = 20
	CParserSizeof                 = 21
	CParserStatic                 = 22
	CParserStruct                 = 23
	CParserSwitch                 = 24
	CParserTypedef                = 25
	CParserUnion                  = 26
	CParserUnsigned               = 27
	CParserVoid                   = 28
	CParserWhile                  = 29
	CParserLeftParen              = 30
	CParserRightParen             = 31
	CParserLeftBracket            = 32
	CParserRightBracket           = 33
	CParserLeftBrace              = 34
	CParserRightBrace             = 35
	CParserLess                   = 36
	CParserLessEqual              = 37
	CParserGreater                = 38
	CParserGreaterEqual           = 39
	CParserLeftShift              = 40
	CParserRightShift             = 41
	CParserPlus                   = 42
	CParserPlusPlus               = 43
	CParserMinus                  = 44
	CParserMinusMinus             = 45
	CParserStar                   = 46
	CParserDiv                    = 47
	CParserMod                    = 48
	CParserAnd                    = 49
	CParserOr                     = 50
	CParserAndAnd                 = 51
	CParserOrOr                   = 52
	CParserCaret                  = 53
	CParserNot                    = 54
	CParserTilde                  = 55
	CParserQuestion               = 56
	CParserColon                  = 57
	CParserSemi                   = 58
	CParserComma                  = 59
	CParserAssign                 = 60
	CParserStarAssign             = 61
	CParserDivAssign              = 62
	CParserModAssign              = 63
	CParserPlusAssign             = 64
	CParserMinusAssign            = 65
	CParserLeftShiftAssign        = 66
	CParserRightShiftAssign       = 67
	CParserAndAssign              = 68
	CParserXorAssign              = 69
	CParserOrAssign               = 70
	CParserEqual                  = 71
	CParserNotEqual               = 72
	CParserArrow                  = 73
	CParserDot                    = 74
	CParserEllipsis               = 75
	CParserIdentifier             = 76
	CParserConstant               = 77
	CParserDigitSequence          = 78
	CParserStringLiteral          = 79
	CParserComplexDefine          = 80
	CParserIncludeDirective       = 81
	CParserAsmBlock               = 82
	CParserLineAfterPreprocessing = 83
	CParserLineDirective          = 84
	CParserPragmaDirective        = 85
	CParserWhitespace             = 86
	CParserNewline                = 87
	CParserBlockComment           = 88
	CParserLineComment            = 89
)

// CParser rules.
const (
	CParserRULE_translation         = 0
	CParserRULE_declaration         = 1
	CParserRULE_functionDeclaration = 2
	CParserRULE_functionArguments   = 3
	CParserRULE_typeSpecifier       = 4
	CParserRULE_variableDeclaration = 5
	CParserRULE_expression          = 6
	CParserRULE_block               = 7
	CParserRULE_statement           = 8
)

// ITranslationContext is an interface to support dynamic dispatch.
type ITranslationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTranslationContext differentiates from other interfaces.
	IsTranslationContext()
}

type TranslationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTranslationContext() *TranslationContext {
	var p = new(TranslationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_translation
	return p
}

func (*TranslationContext) IsTranslationContext() {}

func NewTranslationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TranslationContext {
	var p = new(TranslationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_translation

	return p
}

func (s *TranslationContext) GetParser() antlr.Parser { return s.parser }

func (s *TranslationContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *TranslationContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *TranslationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TranslationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TranslationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitTranslation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Translation() (localctx ITranslationContext) {
	localctx = NewTranslationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CParserRULE_translation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserVoid))) != 0 {
		{
			p.SetState(18)
			p.Declaration()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.FunctionDeclaration()
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *FunctionDeclarationContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *FunctionDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclarationContext) FunctionArguments() IFunctionArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgumentsContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CParserRULE_functionDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.TypeSpecifier()
	}
	{
		p.SetState(27)
		p.Match(CParserIdentifier)
	}
	{
		p.SetState(28)
		p.Match(CParserLeftParen)
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserVoid))) != 0 {
		{
			p.SetState(29)
			p.FunctionArguments()
		}

	}
	{
		p.SetState(32)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(33)
		p.Block()
	}

	return localctx
}

// IFunctionArgumentsContext is an interface to support dynamic dispatch.
type IFunctionArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgumentsContext differentiates from other interfaces.
	IsFunctionArgumentsContext()
}

type FunctionArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgumentsContext() *FunctionArgumentsContext {
	var p = new(FunctionArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_functionArguments
	return p
}

func (*FunctionArgumentsContext) IsFunctionArgumentsContext() {}

func NewFunctionArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgumentsContext {
	var p = new(FunctionArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_functionArguments

	return p
}

func (s *FunctionArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgumentsContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *FunctionArgumentsContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *FunctionArgumentsContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
}

func (s *FunctionArgumentsContext) FunctionArguments() IFunctionArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgumentsContext)
}

func (s *FunctionArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunctionArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) FunctionArguments() (localctx IFunctionArgumentsContext) {
	localctx = NewFunctionArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CParserRULE_functionArguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.TypeSpecifier()
	}
	{
		p.SetState(36)
		p.Match(CParserIdentifier)
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(37)
			p.Match(CParserComma)
		}
		{
			p.SetState(38)
			p.FunctionArguments()
		}

	}

	return localctx
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_typeSpecifier
	return p
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) Void() antlr.TerminalNode {
	return s.GetToken(CParserVoid, 0)
}

func (s *TypeSpecifierContext) Int() antlr.TerminalNode {
	return s.GetToken(CParserInt, 0)
}

func (s *TypeSpecifierContext) Short() antlr.TerminalNode {
	return s.GetToken(CParserShort, 0)
}

func (s *TypeSpecifierContext) Long() antlr.TerminalNode {
	return s.GetToken(CParserLong, 0)
}

func (s *TypeSpecifierContext) Char() antlr.TerminalNode {
	return s.GetToken(CParserChar, 0)
}

func (s *TypeSpecifierContext) Float() antlr.TerminalNode {
	return s.GetToken(CParserFloat, 0)
}

func (s *TypeSpecifierContext) Double() antlr.TerminalNode {
	return s.GetToken(CParserDouble, 0)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitTypeSpecifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CParserRULE_typeSpecifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserVoid))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *VariableDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *VariableDeclarationContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *VariableDeclarationContext) Assign() antlr.TerminalNode {
	return s.GetToken(CParserAssign, 0)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CParserRULE_variableDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.TypeSpecifier()
	}
	{
		p.SetState(44)
		p.Match(CParserIdentifier)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserAssign {
		{
			p.SetState(45)
			p.Match(CParserAssign)
		}
		{
			p.SetState(46)
			p.Expression()
		}

	}
	{
		p.SetState(49)
		p.Match(CParserSemi)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *ExpressionContext) Constant() antlr.TerminalNode {
	return s.GetToken(CParserConstant, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserIdentifier || _la == CParserConstant) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(CParserLeftBrace, 0)
}

func (s *BlockContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(CParserRightBrace, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(CParserLeftBrace)
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserVoid))) != 0 {
		{
			p.SetState(54)
			p.Statement()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
		p.Match(CParserRightBrace)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.VariableDeclaration()
	}

	return localctx
}
