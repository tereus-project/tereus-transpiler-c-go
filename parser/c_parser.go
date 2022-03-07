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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 91, 183,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 7, 2, 38, 10, 2, 12, 2, 14, 2, 41, 11, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 49, 10, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 58, 10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 5, 7, 71, 10, 7, 3, 7, 3, 7, 7, 7, 75, 10, 7, 12, 7, 14,
	7, 78, 11, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 86, 10, 9, 3, 9,
	3, 9, 5, 9, 90, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	5, 10, 99, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 5, 10, 112, 10, 10, 3, 10, 7, 10, 115, 10, 10,
	12, 10, 14, 10, 118, 11, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 5, 13, 127, 10, 13, 3, 14, 3, 14, 7, 14, 131, 10, 14, 12, 14, 14,
	14, 134, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 5, 15, 141, 10, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 149, 10, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 158, 10, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 5, 17, 164, 10, 17, 3, 17, 3, 17, 5, 17, 168, 10, 17, 3,
	17, 3, 17, 5, 17, 172, 10, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 2, 4, 12, 18, 19, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 4, 3, 2, 62, 72, 7, 2, 38, 44, 46,
	46, 48, 52, 55, 55, 73, 74, 2, 196, 2, 39, 3, 2, 2, 2, 4, 42, 3, 2, 2,
	2, 6, 44, 3, 2, 2, 2, 8, 53, 3, 2, 2, 2, 10, 59, 3, 2, 2, 2, 12, 70, 3,
	2, 2, 2, 14, 79, 3, 2, 2, 2, 16, 82, 3, 2, 2, 2, 18, 98, 3, 2, 2, 2, 20,
	119, 3, 2, 2, 2, 22, 121, 3, 2, 2, 2, 24, 123, 3, 2, 2, 2, 26, 128, 3,
	2, 2, 2, 28, 148, 3, 2, 2, 2, 30, 150, 3, 2, 2, 2, 32, 159, 3, 2, 2, 2,
	34, 176, 3, 2, 2, 2, 36, 38, 5, 4, 3, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3,
	2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 3, 3, 2, 2, 2, 41,
	39, 3, 2, 2, 2, 42, 43, 5, 6, 4, 2, 43, 5, 3, 2, 2, 2, 44, 45, 5, 12, 7,
	2, 45, 46, 7, 78, 2, 2, 46, 48, 7, 32, 2, 2, 47, 49, 5, 8, 5, 2, 48, 47,
	3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 7, 33, 2, 2,
	51, 52, 5, 26, 14, 2, 52, 7, 3, 2, 2, 2, 53, 54, 5, 12, 7, 2, 54, 57, 7,
	78, 2, 2, 55, 56, 7, 61, 2, 2, 56, 58, 5, 8, 5, 2, 57, 55, 3, 2, 2, 2,
	57, 58, 3, 2, 2, 2, 58, 9, 3, 2, 2, 2, 59, 60, 7, 20, 2, 2, 60, 61, 5,
	18, 10, 2, 61, 11, 3, 2, 2, 2, 62, 63, 8, 7, 1, 2, 63, 71, 7, 30, 2, 2,
	64, 71, 7, 18, 2, 2, 65, 71, 7, 21, 2, 2, 66, 71, 7, 19, 2, 2, 67, 71,
	7, 5, 2, 2, 68, 71, 7, 14, 2, 2, 69, 71, 7, 10, 2, 2, 70, 62, 3, 2, 2,
	2, 70, 64, 3, 2, 2, 2, 70, 65, 3, 2, 2, 2, 70, 66, 3, 2, 2, 2, 70, 67,
	3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 69, 3, 2, 2, 2, 71, 76, 3, 2, 2, 2,
	72, 73, 12, 3, 2, 2, 73, 75, 7, 48, 2, 2, 74, 72, 3, 2, 2, 2, 75, 78, 3,
	2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 13, 3, 2, 2, 2, 78,
	76, 3, 2, 2, 2, 79, 80, 5, 12, 7, 2, 80, 81, 5, 16, 9, 2, 81, 15, 3, 2,
	2, 2, 82, 85, 7, 78, 2, 2, 83, 84, 7, 62, 2, 2, 84, 86, 5, 18, 10, 2, 85,
	83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 88, 7, 61,
	2, 2, 88, 90, 5, 16, 9, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90,
	17, 3, 2, 2, 2, 91, 92, 8, 10, 1, 2, 92, 99, 7, 78, 2, 2, 93, 99, 7, 79,
	2, 2, 94, 95, 7, 32, 2, 2, 95, 96, 5, 18, 10, 2, 96, 97, 7, 33, 2, 2, 97,
	99, 3, 2, 2, 2, 98, 91, 3, 2, 2, 2, 98, 93, 3, 2, 2, 2, 98, 94, 3, 2, 2,
	2, 99, 116, 3, 2, 2, 2, 100, 101, 12, 4, 2, 2, 101, 102, 5, 20, 11, 2,
	102, 103, 5, 18, 10, 5, 103, 115, 3, 2, 2, 2, 104, 105, 12, 3, 2, 2, 105,
	106, 5, 22, 12, 2, 106, 107, 5, 18, 10, 4, 107, 115, 3, 2, 2, 2, 108, 109,
	12, 5, 2, 2, 109, 111, 7, 32, 2, 2, 110, 112, 5, 24, 13, 2, 111, 110, 3,
	2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 115, 7, 33, 2,
	2, 114, 100, 3, 2, 2, 2, 114, 104, 3, 2, 2, 2, 114, 108, 3, 2, 2, 2, 115,
	118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 19, 3,
	2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 120, 9, 2, 2, 2, 120, 21, 3, 2, 2,
	2, 121, 122, 9, 3, 2, 2, 122, 23, 3, 2, 2, 2, 123, 126, 5, 18, 10, 2, 124,
	125, 7, 61, 2, 2, 125, 127, 5, 24, 13, 2, 126, 124, 3, 2, 2, 2, 126, 127,
	3, 2, 2, 2, 127, 25, 3, 2, 2, 2, 128, 132, 7, 36, 2, 2, 129, 131, 5, 28,
	15, 2, 130, 129, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2,
	132, 133, 3, 2, 2, 2, 133, 135, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135,
	136, 7, 37, 2, 2, 136, 27, 3, 2, 2, 2, 137, 141, 5, 14, 8, 2, 138, 141,
	5, 18, 10, 2, 139, 141, 5, 10, 6, 2, 140, 137, 3, 2, 2, 2, 140, 138, 3,
	2, 2, 2, 140, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 7, 60, 2,
	2, 143, 149, 3, 2, 2, 2, 144, 149, 5, 30, 16, 2, 145, 149, 5, 32, 17, 2,
	146, 149, 5, 34, 18, 2, 147, 149, 5, 26, 14, 2, 148, 140, 3, 2, 2, 2, 148,
	144, 3, 2, 2, 2, 148, 145, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 147,
	3, 2, 2, 2, 149, 29, 3, 2, 2, 2, 150, 151, 7, 17, 2, 2, 151, 152, 7, 32,
	2, 2, 152, 153, 5, 18, 10, 2, 153, 154, 7, 33, 2, 2, 154, 157, 5, 28, 15,
	2, 155, 156, 7, 11, 2, 2, 156, 158, 5, 28, 15, 2, 157, 155, 3, 2, 2, 2,
	157, 158, 3, 2, 2, 2, 158, 31, 3, 2, 2, 2, 159, 160, 7, 15, 2, 2, 160,
	163, 7, 32, 2, 2, 161, 164, 5, 18, 10, 2, 162, 164, 5, 14, 8, 2, 163, 161,
	3, 2, 2, 2, 163, 162, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 3, 2,
	2, 2, 165, 167, 7, 60, 2, 2, 166, 168, 5, 18, 10, 2, 167, 166, 3, 2, 2,
	2, 167, 168, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 171, 7, 60, 2, 2, 170,
	172, 5, 18, 10, 2, 171, 170, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 173,
	3, 2, 2, 2, 173, 174, 7, 33, 2, 2, 174, 175, 5, 28, 15, 2, 175, 33, 3,
	2, 2, 2, 176, 177, 7, 31, 2, 2, 177, 178, 7, 32, 2, 2, 178, 179, 5, 18,
	10, 2, 179, 180, 7, 33, 2, 2, 180, 181, 5, 28, 15, 2, 181, 35, 3, 2, 2,
	2, 21, 39, 48, 57, 70, 76, 85, 89, 98, 111, 114, 116, 126, 132, 140, 148,
	157, 163, 167, 171,
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
	"functionReturn", "typeSpecifier", "variableDeclaration", "variableDeclarationList",
	"expression", "assignementOperator", "binaryOperator", "functionCallArguments",
	"block", "statement", "ifStatement", "forStatement", "whileStatement",
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
	CParserRULE_translation             = 0
	CParserRULE_declaration             = 1
	CParserRULE_functionDeclaration     = 2
	CParserRULE_functionArguments       = 3
	CParserRULE_functionReturn          = 4
	CParserRULE_typeSpecifier           = 5
	CParserRULE_variableDeclaration     = 6
	CParserRULE_variableDeclarationList = 7
	CParserRULE_expression              = 8
	CParserRULE_assignementOperator     = 9
	CParserRULE_binaryOperator          = 10
	CParserRULE_functionCallArguments   = 11
	CParserRULE_block                   = 12
	CParserRULE_statement               = 13
	CParserRULE_ifStatement             = 14
	CParserRULE_forStatement            = 15
	CParserRULE_whileStatement          = 16
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserVoid))) != 0 {
		{
			p.SetState(34)
			p.Declaration()
		}

		p.SetState(39)
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
		p.SetState(40)
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
		p.SetState(42)
		p.typeSpecifier(0)
	}
	{
		p.SetState(43)
		p.Match(CParserIdentifier)
	}
	{
		p.SetState(44)
		p.Match(CParserLeftParen)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserVoid))) != 0 {
		{
			p.SetState(45)
			p.FunctionArguments()
		}

	}
	{
		p.SetState(48)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(49)
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
		p.SetState(51)
		p.typeSpecifier(0)
	}
	{
		p.SetState(52)
		p.Match(CParserIdentifier)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(53)
			p.Match(CParserComma)
		}
		{
			p.SetState(54)
			p.FunctionArguments()
		}

	}

	return localctx
}

// IFunctionReturnContext is an interface to support dynamic dispatch.
type IFunctionReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionReturnContext differentiates from other interfaces.
	IsFunctionReturnContext()
}

type FunctionReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionReturnContext() *FunctionReturnContext {
	var p = new(FunctionReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_functionReturn
	return p
}

func (*FunctionReturnContext) IsFunctionReturnContext() {}

func NewFunctionReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionReturnContext {
	var p = new(FunctionReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_functionReturn

	return p
}

func (s *FunctionReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionReturnContext) Return() antlr.TerminalNode {
	return s.GetToken(CParserReturn, 0)
}

func (s *FunctionReturnContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunctionReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) FunctionReturn() (localctx IFunctionReturnContext) {
	localctx = NewFunctionReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CParserRULE_functionReturn)

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
		p.SetState(57)
		p.Match(CParserReturn)
	}
	{
		p.SetState(58)
		p.expression(0)
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

func (s *TypeSpecifierContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *TypeSpecifierContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
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
	return p.typeSpecifier(0)
}

func (p *CParser) typeSpecifier(_p int) (localctx ITypeSpecifierContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeSpecifierContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, CParserRULE_typeSpecifier, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserVoid:
		{
			p.SetState(61)
			p.Match(CParserVoid)
		}

	case CParserInt:
		{
			p.SetState(62)
			p.Match(CParserInt)
		}

	case CParserShort:
		{
			p.SetState(63)
			p.Match(CParserShort)
		}

	case CParserLong:
		{
			p.SetState(64)
			p.Match(CParserLong)
		}

	case CParserChar:
		{
			p.SetState(65)
			p.Match(CParserChar)
		}

	case CParserFloat:
		{
			p.SetState(66)
			p.Match(CParserFloat)
		}

	case CParserDouble:
		{
			p.SetState(67)
			p.Match(CParserDouble)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTypeSpecifierContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CParserRULE_typeSpecifier)
			p.SetState(70)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(71)
				p.Match(CParserStar)
			}

		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
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

func (s *VariableDeclarationContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
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
	p.EnterRule(localctx, 12, CParserRULE_variableDeclaration)

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
		p.SetState(77)
		p.typeSpecifier(0)
	}
	{
		p.SetState(78)
		p.VariableDeclarationList()
	}

	return localctx
}

// IVariableDeclarationListContext is an interface to support dynamic dispatch.
type IVariableDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationListContext differentiates from other interfaces.
	IsVariableDeclarationListContext()
}

type VariableDeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationListContext() *VariableDeclarationListContext {
	var p = new(VariableDeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_variableDeclarationList
	return p
}

func (*VariableDeclarationListContext) IsVariableDeclarationListContext() {}

func NewVariableDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationListContext {
	var p = new(VariableDeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_variableDeclarationList

	return p
}

func (s *VariableDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationListContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *VariableDeclarationListContext) Assign() antlr.TerminalNode {
	return s.GetToken(CParserAssign, 0)
}

func (s *VariableDeclarationListContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationListContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
}

func (s *VariableDeclarationListContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *VariableDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitVariableDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) VariableDeclarationList() (localctx IVariableDeclarationListContext) {
	localctx = NewVariableDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CParserRULE_variableDeclarationList)
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
		p.SetState(80)
		p.Match(CParserIdentifier)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserAssign {
		{
			p.SetState(81)
			p.Match(CParserAssign)
		}
		{
			p.SetState(82)
			p.expression(0)
		}

	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(85)
			p.Match(CParserComma)
		}
		{
			p.SetState(86)
			p.VariableDeclarationList()
		}

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

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParenthesizedExpressionContext struct {
	*ExpressionContext
}

func NewParenthesizedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesizedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *ParenthesizedExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesizedExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *ParenthesizedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitParenthesizedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentExpressionContext struct {
	*ExpressionContext
}

func NewAssignmentExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AssignmentExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentExpressionContext) AssignementOperator() IAssignementOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignementOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignementOperatorContext)
}

func (s *AssignmentExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAssignmentExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryExpressionContext struct {
	*ExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) BinaryOperator() IBinaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOperatorContext)
}

func (s *BinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstantExpressionContext struct {
	*ExpressionContext
}

func NewConstantExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantExpressionContext {
	var p = new(ConstantExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ConstantExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantExpressionContext) Constant() antlr.TerminalNode {
	return s.GetToken(CParserConstant, 0)
}

func (s *ConstantExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitConstantExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExpressionContext struct {
	*ExpressionContext
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionCallExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *FunctionCallExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *FunctionCallExpressionContext) FunctionCallArguments() IFunctionCallArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallArgumentsContext)
}

func (s *FunctionCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunctionCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	*ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, CParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserIdentifier:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(90)
			p.Match(CParserIdentifier)
		}

	case CParserConstant:
		localctx = NewConstantExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.Match(CParserConstant)
		}

	case CParserLeftParen:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.Match(CParserLeftParen)
		}
		{
			p.SetState(93)
			p.expression(0)
		}
		{
			p.SetState(94)
			p.Match(CParserRightParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAssignmentExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(99)
					p.AssignementOperator()
				}
				{
					p.SetState(100)
					p.expression(3)
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(103)
					p.BinaryOperator()
				}
				{
					p.SetState(104)
					p.expression(2)
				}

			case 3:
				localctx = NewFunctionCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(106)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(107)
					p.Match(CParserLeftParen)
				}
				p.SetState(109)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == CParserLeftParen || _la == CParserIdentifier || _la == CParserConstant {
					{
						p.SetState(108)
						p.FunctionCallArguments()
					}

				}
				{
					p.SetState(111)
					p.Match(CParserRightParen)
				}

			}

		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignementOperatorContext is an interface to support dynamic dispatch.
type IAssignementOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignementOperatorContext differentiates from other interfaces.
	IsAssignementOperatorContext()
}

type AssignementOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignementOperatorContext() *AssignementOperatorContext {
	var p = new(AssignementOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_assignementOperator
	return p
}

func (*AssignementOperatorContext) IsAssignementOperatorContext() {}

func NewAssignementOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignementOperatorContext {
	var p = new(AssignementOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_assignementOperator

	return p
}

func (s *AssignementOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignementOperatorContext) Assign() antlr.TerminalNode {
	return s.GetToken(CParserAssign, 0)
}

func (s *AssignementOperatorContext) StarAssign() antlr.TerminalNode {
	return s.GetToken(CParserStarAssign, 0)
}

func (s *AssignementOperatorContext) DivAssign() antlr.TerminalNode {
	return s.GetToken(CParserDivAssign, 0)
}

func (s *AssignementOperatorContext) ModAssign() antlr.TerminalNode {
	return s.GetToken(CParserModAssign, 0)
}

func (s *AssignementOperatorContext) PlusAssign() antlr.TerminalNode {
	return s.GetToken(CParserPlusAssign, 0)
}

func (s *AssignementOperatorContext) MinusAssign() antlr.TerminalNode {
	return s.GetToken(CParserMinusAssign, 0)
}

func (s *AssignementOperatorContext) LeftShiftAssign() antlr.TerminalNode {
	return s.GetToken(CParserLeftShiftAssign, 0)
}

func (s *AssignementOperatorContext) RightShiftAssign() antlr.TerminalNode {
	return s.GetToken(CParserRightShiftAssign, 0)
}

func (s *AssignementOperatorContext) AndAssign() antlr.TerminalNode {
	return s.GetToken(CParserAndAssign, 0)
}

func (s *AssignementOperatorContext) XorAssign() antlr.TerminalNode {
	return s.GetToken(CParserXorAssign, 0)
}

func (s *AssignementOperatorContext) OrAssign() antlr.TerminalNode {
	return s.GetToken(CParserOrAssign, 0)
}

func (s *AssignementOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignementOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignementOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAssignementOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) AssignementOperator() (localctx IAssignementOperatorContext) {
	localctx = NewAssignementOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CParserRULE_assignementOperator)
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
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-60)&-(0x1f+1)) == 0 && ((1<<uint((_la-60)))&((1<<(CParserAssign-60))|(1<<(CParserStarAssign-60))|(1<<(CParserDivAssign-60))|(1<<(CParserModAssign-60))|(1<<(CParserPlusAssign-60))|(1<<(CParserMinusAssign-60))|(1<<(CParserLeftShiftAssign-60))|(1<<(CParserRightShiftAssign-60))|(1<<(CParserAndAssign-60))|(1<<(CParserXorAssign-60))|(1<<(CParserOrAssign-60)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryOperatorContext is an interface to support dynamic dispatch.
type IBinaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOperatorContext differentiates from other interfaces.
	IsBinaryOperatorContext()
}

type BinaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperatorContext() *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_binaryOperator
	return p
}

func (*BinaryOperatorContext) IsBinaryOperatorContext() {}

func NewBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_binaryOperator

	return p
}

func (s *BinaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperatorContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

func (s *BinaryOperatorContext) Div() antlr.TerminalNode {
	return s.GetToken(CParserDiv, 0)
}

func (s *BinaryOperatorContext) Mod() antlr.TerminalNode {
	return s.GetToken(CParserMod, 0)
}

func (s *BinaryOperatorContext) Plus() antlr.TerminalNode {
	return s.GetToken(CParserPlus, 0)
}

func (s *BinaryOperatorContext) Minus() antlr.TerminalNode {
	return s.GetToken(CParserMinus, 0)
}

func (s *BinaryOperatorContext) LeftShift() antlr.TerminalNode {
	return s.GetToken(CParserLeftShift, 0)
}

func (s *BinaryOperatorContext) RightShift() antlr.TerminalNode {
	return s.GetToken(CParserRightShift, 0)
}

func (s *BinaryOperatorContext) And() antlr.TerminalNode {
	return s.GetToken(CParserAnd, 0)
}

func (s *BinaryOperatorContext) Caret() antlr.TerminalNode {
	return s.GetToken(CParserCaret, 0)
}

func (s *BinaryOperatorContext) Or() antlr.TerminalNode {
	return s.GetToken(CParserOr, 0)
}

func (s *BinaryOperatorContext) Less() antlr.TerminalNode {
	return s.GetToken(CParserLess, 0)
}

func (s *BinaryOperatorContext) Greater() antlr.TerminalNode {
	return s.GetToken(CParserGreater, 0)
}

func (s *BinaryOperatorContext) LessEqual() antlr.TerminalNode {
	return s.GetToken(CParserLessEqual, 0)
}

func (s *BinaryOperatorContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(CParserGreaterEqual, 0)
}

func (s *BinaryOperatorContext) Equal() antlr.TerminalNode {
	return s.GetToken(CParserEqual, 0)
}

func (s *BinaryOperatorContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(CParserNotEqual, 0)
}

func (s *BinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBinaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) BinaryOperator() (localctx IBinaryOperatorContext) {
	localctx = NewBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CParserRULE_binaryOperator)
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
		p.SetState(119)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(CParserLess-36))|(1<<(CParserLessEqual-36))|(1<<(CParserGreater-36))|(1<<(CParserGreaterEqual-36))|(1<<(CParserLeftShift-36))|(1<<(CParserRightShift-36))|(1<<(CParserPlus-36))|(1<<(CParserMinus-36))|(1<<(CParserStar-36))|(1<<(CParserDiv-36))|(1<<(CParserMod-36))|(1<<(CParserAnd-36))|(1<<(CParserOr-36))|(1<<(CParserCaret-36)))) != 0) || _la == CParserEqual || _la == CParserNotEqual) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionCallArgumentsContext is an interface to support dynamic dispatch.
type IFunctionCallArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallArgumentsContext differentiates from other interfaces.
	IsFunctionCallArgumentsContext()
}

type FunctionCallArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallArgumentsContext() *FunctionCallArgumentsContext {
	var p = new(FunctionCallArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_functionCallArguments
	return p
}

func (*FunctionCallArgumentsContext) IsFunctionCallArgumentsContext() {}

func NewFunctionCallArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallArgumentsContext {
	var p = new(FunctionCallArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_functionCallArguments

	return p
}

func (s *FunctionCallArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallArgumentsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionCallArgumentsContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
}

func (s *FunctionCallArgumentsContext) FunctionCallArguments() IFunctionCallArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallArgumentsContext)
}

func (s *FunctionCallArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunctionCallArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) FunctionCallArguments() (localctx IFunctionCallArgumentsContext) {
	localctx = NewFunctionCallArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CParserRULE_functionCallArguments)
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
		p.SetState(121)
		p.expression(0)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(122)
			p.Match(CParserComma)
		}
		{
			p.SetState(123)
			p.FunctionCallArguments()
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
	p.EnterRule(localctx, 24, CParserRULE_block)
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
		p.SetState(126)
		p.Match(CParserLeftBrace)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-3)&-(0x1f+1)) == 0 && ((1<<uint((_la-3)))&((1<<(CParserChar-3))|(1<<(CParserDouble-3))|(1<<(CParserFloat-3))|(1<<(CParserFor-3))|(1<<(CParserIf-3))|(1<<(CParserInt-3))|(1<<(CParserLong-3))|(1<<(CParserReturn-3))|(1<<(CParserShort-3))|(1<<(CParserVoid-3))|(1<<(CParserWhile-3))|(1<<(CParserLeftParen-3))|(1<<(CParserLeftBrace-3)))) != 0) || _la == CParserIdentifier || _la == CParserConstant {
		{
			p.SetState(127)
			p.Statement()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)
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

func (s *StatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) FunctionReturn() IFunctionReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionReturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionReturnContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
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
	p.EnterRule(localctx, 26, CParserRULE_statement)

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

	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserChar, CParserDouble, CParserFloat, CParserInt, CParserLong, CParserReturn, CParserShort, CParserVoid, CParserLeftParen, CParserIdentifier, CParserConstant:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(138)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CParserChar, CParserDouble, CParserFloat, CParserInt, CParserLong, CParserShort, CParserVoid:
			{
				p.SetState(135)
				p.VariableDeclaration()
			}

		case CParserLeftParen, CParserIdentifier, CParserConstant:
			{
				p.SetState(136)
				p.expression(0)
			}

		case CParserReturn:
			{
				p.SetState(137)
				p.FunctionReturn()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(140)
			p.Match(CParserSemi)
		}

	case CParserIf:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.IfStatement()
		}

	case CParserFor:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(143)
			p.ForStatement()
		}

	case CParserWhile:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(144)
			p.WhileStatement()
		}

	case CParserLeftBrace:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(145)
			p.Block()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(CParserIf, 0)
}

func (s *IfStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(CParserElse, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CParserRULE_ifStatement)

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
		p.SetState(148)
		p.Match(CParserIf)
	}
	{
		p.SetState(149)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(150)
		p.expression(0)
	}
	{
		p.SetState(151)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(152)
		p.Statement()
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(153)
			p.Match(CParserElse)
		}
		{
			p.SetState(154)
			p.Statement()
		}

	}

	return localctx
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IExpressionContext

	// GetCondition returns the condition rule contexts.
	GetCondition() IExpressionContext

	// GetPost returns the post rule contexts.
	GetPost() IExpressionContext

	// SetInit sets the init rule contexts.
	SetInit(IExpressionContext)

	// SetCondition sets the condition rule contexts.
	SetCondition(IExpressionContext)

	// SetPost sets the post rule contexts.
	SetPost(IExpressionContext)

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	init      IExpressionContext
	condition IExpressionContext
	post      IExpressionContext
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_forStatement
	return p
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) GetInit() IExpressionContext { return s.init }

func (s *ForStatementContext) GetCondition() IExpressionContext { return s.condition }

func (s *ForStatementContext) GetPost() IExpressionContext { return s.post }

func (s *ForStatementContext) SetInit(v IExpressionContext) { s.init = v }

func (s *ForStatementContext) SetCondition(v IExpressionContext) { s.condition = v }

func (s *ForStatementContext) SetPost(v IExpressionContext) { s.post = v }

func (s *ForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(CParserFor, 0)
}

func (s *ForStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *ForStatementContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(CParserSemi)
}

func (s *ForStatementContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(CParserSemi, i)
}

func (s *ForStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *ForStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ForStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ForStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CParserRULE_forStatement)
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
		p.SetState(157)
		p.Match(CParserFor)
	}
	{
		p.SetState(158)
		p.Match(CParserLeftParen)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserLeftParen, CParserIdentifier, CParserConstant:
		{
			p.SetState(159)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).init = _x
		}

	case CParserChar, CParserDouble, CParserFloat, CParserInt, CParserLong, CParserShort, CParserVoid:
		{
			p.SetState(160)
			p.VariableDeclaration()
		}

	case CParserSemi:

	default:
	}
	{
		p.SetState(163)
		p.Match(CParserSemi)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserLeftParen || _la == CParserIdentifier || _la == CParserConstant {
		{
			p.SetState(164)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).condition = _x
		}

	}
	{
		p.SetState(167)
		p.Match(CParserSemi)
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserLeftParen || _la == CParserIdentifier || _la == CParserConstant {
		{
			p.SetState(168)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).post = _x
		}

	}
	{
		p.SetState(171)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(172)
		p.Statement()
	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(CParserWhile, 0)
}

func (s *WhileStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *WhileStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CParserRULE_whileStatement)

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
		p.SetState(174)
		p.Match(CParserWhile)
	}
	{
		p.SetState(175)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(176)
		p.expression(0)
	}
	{
		p.SetState(177)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(178)
		p.Statement()
	}

	return localctx
}

func (p *CParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *TypeSpecifierContext = nil
		if localctx != nil {
			t = localctx.(*TypeSpecifierContext)
		}
		return p.TypeSpecifier_Sempred(t, predIndex)

	case 8:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CParser) TypeSpecifier_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
