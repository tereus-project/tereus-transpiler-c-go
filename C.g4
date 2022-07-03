grammar C;

translation: declaration*;

declaration:
	functionDeclaration
	| (structDeclaration ';')
	| (enumDeclaration ';')
	| includePreprocessor
	| (typedefDeclaration ';')
	| BlockComment
	| LineComment;

functionDeclaration:
	typeSpecifier Identifier '(' functionArguments? ')' block;

functionArguments:
	typeSpecifier Identifier (',' functionArguments)?;

functionReturn: 'return' expression?;

typeSpecifier:
	(
		('signed' | 'unsigned')? (
			'int'
			| 'short'
			| 'long'
			| 'char'
			| Identifier
		)
	)													# TypeSpecifierWithModifier
	| ('void' | 'float' | 'double' | structDeclaration)	# TypeSpecifierGeneric
	| typeSpecifier '*'									# TypeSpecifierPointer;

structDeclaration:
	'struct' (
		Identifier structDeclarationBody?
		| structDeclarationBody
	);

structDeclarationBody: '{' structProperty* '}';

structProperty: typeSpecifier Identifier? ';';

enumDeclaration: 'enum' Identifier? '{' enumProperties ','? '}';

enumProperties:
	Identifier ('=' expression)? (',' enumProperties)?;

typedefDeclaration: 'typedef' typeSpecifier Identifier;

variableDeclaration: typeSpecifier variableDeclarationList;

variableDeclarationList:
	Identifier sizedArrayModifier? (
		'=' (expression | listInitialization)
	)? (',' variableDeclarationList)?;

sizedArrayModifier: '[' expression ']';

listInitialization:
	'{' (expression (',' expression)* ','?)? '}';

namedListInitialization:
	'{' (
		namedListInitializationItem (
			',' namedListInitializationItem
		)* ','?
	)? '}';

namedListInitializationItem: '.' Identifier '=' expression;

expression:
	Identifier													# IdentifierExpression
	| (IntegerConstant | FloatingConstant | CharacterConstant)	# ConstantExpression
	| StringLiteral												# ConstantStringExpression
	| '(' typeSpecifier ')' (
		listInitialization
		| namedListInitialization
	)																# StructInitializationExpression
	| expression '.' Identifier										# PropertyAccessExpression
	| expression '->' Identifier									# PointerPropertyAccessExpression
	| expression '[' expression ']'									# ArrayIndexExpression
	| expression '(' functionCallArguments? ')'						# FunctionCallExpression
	| '(' typeSpecifier ')' expression								# CastExpression
	| '(' expression ')'											# ParenthesizedExpression
	| expression unaryOperatorPost									# UnaryExpressionPost
	| unaryOperatorPre expression									# UnaryExpressionPre
	| expression assignementOperator expression						# AssignmentExpression
	| expression binaryOperator expression							# BinaryExpression
	| Sizeof ('(' typeSpecifier ')' | typeSpecifier | expression)	# SizeofExpression
	| expression '?' expression ':' expression						# TernaryExpression;

assignementOperator:
	'='
	| '*='
	| '/='
	| '%='
	| '+='
	| '-='
	| '<<='
	| '>>='
	| '&='
	| '^='
	| '|=';

binaryOperator:
	'*'
	| '/'
	| '%'
	| '+'
	| '-'
	| '<<'
	| '>>'
	| '&'
	| '^'
	| '|'
	| '<'
	| '>'
	| '<='
	| '>='
	| '=='
	| '!='
	| '&&'
	| '||';

unaryOperatorPost: '++' | '--';

unaryOperatorPre:
	'+'
	| '++'
	| '-'
	| '--'
	| '~'
	| '!'
	| '&'
	| '*';

functionCallArguments: expression (',' functionCallArguments)?;

block: '{' statement* '}';

statement: (
		(
			variableDeclaration
			| expression
			| functionReturn
			| 'break'
			| 'continue'
			| structDeclaration
			| enumDeclaration
			| gotoStatement
			| doWhileStatement
		) ';'
	)
	| ifStatement
	| switchStatement
	| forStatement
	| whileStatement
	| block
	| labelStatement
	| BlockComment
	| LineComment;

ifStatement:
	'if' '(' expression ')' statement ('else' statement)?;

switchStatement:
	'switch' '(' expression ')' '{' caseStatement* defaultStatement? '}';

caseStatement: 'case' expression ':' statement*;

defaultStatement: 'default' ':' statement*;

forStatement:
	'for' '(' (init = expression | variableDeclaration)? ';' (
		condition = expression
	)? ';' (post = expression)? ')' statement;

doWhileStatement: 'do' statement 'while' '(' expression ')';

whileStatement: 'while' '(' expression ')' statement;

gotoStatement: 'goto' Identifier;

labelStatement: Identifier ':';

includePreprocessor: IncludeDirective;

Break: 'break';
Case: 'case';
Char: 'char';
Const: 'const';
Continue: 'continue';
Default: 'default';
Do: 'do';
Double: 'double';
Else: 'else';
Enum: 'enum';
Extern: 'extern';
Float: 'float';
For: 'for';
Goto: 'goto';
If: 'if';
Int: 'int';
Long: 'long';
Return: 'return';
Short: 'short';
Signed: 'signed';
Sizeof: 'sizeof';
Static: 'static';
Struct: 'struct';
Switch: 'switch';
Typedef: 'typedef';
Union: 'union';
Unsigned: 'unsigned';
Void: 'void';
While: 'while';

LeftParen: '(';
RightParen: ')';
LeftBracket: '[';
RightBracket: ']';
LeftBrace: '{';
RightBrace: '}';

Less: '<';
LessEqual: '<=';
Greater: '>';
GreaterEqual: '>=';
LeftShift: '<<';
RightShift: '>>';

Plus: '+';
PlusPlus: '++';
Minus: '-';
MinusMinus: '--';
Star: '*';
Div: '/';
Mod: '%';

And: '&';
Or: '|';
AndAnd: '&&';
OrOr: '||';
Caret: '^';
Not: '!';
Tilde: '~';

Question: '?';
Colon: ':';
Semi: ';';
Comma: ',';

Assign: '=';
// '*=' | '/=' | '%=' | '+=' | '-=' | '<<=' | '>>=' | '&=' | '^=' | '|='
StarAssign: '*=';
DivAssign: '/=';
ModAssign: '%=';
PlusAssign: '+=';
MinusAssign: '-=';
LeftShiftAssign: '<<=';
RightShiftAssign: '>>=';
AndAssign: '&=';
XorAssign: '^=';
OrAssign: '|=';

Equal: '==';
NotEqual: '!=';

Arrow: '->';
Dot: '.';
Ellipsis: '...';

Identifier: IdentifierNondigit ( IdentifierNondigit | Digit)*;

fragment IdentifierNondigit:
	Nondigit
	| UniversalCharacterName; //|   // other implementation-defined characters...

fragment Nondigit: [a-zA-Z_];

fragment Digit: [0-9];

fragment UniversalCharacterName:
	'\\u' HexQuad
	| '\\U' HexQuad HexQuad;

fragment HexQuad:
	HexadecimalDigit HexadecimalDigit HexadecimalDigit HexadecimalDigit;

IntegerConstant:
	DecimalConstant IntegerSuffix?
	| OctalConstant IntegerSuffix?
	| HexadecimalConstant IntegerSuffix?
	| BinaryConstant;

fragment BinaryConstant: '0' [bB] [0-1]+;

fragment DecimalConstant: NonzeroDigit Digit*;

fragment OctalConstant: '0' OctalDigit*;

fragment HexadecimalConstant:
	HexadecimalPrefix HexadecimalDigit+;

fragment HexadecimalPrefix: '0' [xX];

fragment NonzeroDigit: [1-9];

fragment OctalDigit: [0-7];

fragment HexadecimalDigit: [0-9a-fA-F];

fragment IntegerSuffix:
	UnsignedSuffix LongSuffix?
	| UnsignedSuffix LongLongSuffix
	| LongSuffix UnsignedSuffix?
	| LongLongSuffix UnsignedSuffix?;

fragment UnsignedSuffix: [uU];

fragment LongSuffix: [lL];

fragment LongLongSuffix: 'll' | 'LL';

FloatingConstant:
	DecimalFloatingConstant
	| HexadecimalFloatingConstant;

fragment DecimalFloatingConstant:
	FractionalConstant ExponentPart? FloatingSuffix?
	| DigitSequence ExponentPart FloatingSuffix?;

fragment HexadecimalFloatingConstant:
	HexadecimalPrefix (
		HexadecimalFractionalConstant
		| HexadecimalDigitSequence
	) BinaryExponentPart FloatingSuffix?;

fragment FractionalConstant:
	DigitSequence? '.' DigitSequence
	| DigitSequence '.';

fragment ExponentPart: [eE] Sign? DigitSequence;

fragment Sign: [+-];

DigitSequence: Digit+;

fragment HexadecimalFractionalConstant:
	HexadecimalDigitSequence? '.' HexadecimalDigitSequence
	| HexadecimalDigitSequence '.';

fragment BinaryExponentPart: [pP] Sign? DigitSequence;

fragment HexadecimalDigitSequence: HexadecimalDigit+;

fragment FloatingSuffix: [flFL];

CharacterConstant:
	'\'' CCharSequence '\''
	| 'L\'' CCharSequence '\''
	| 'u\'' CCharSequence '\''
	| 'U\'' CCharSequence '\'';

fragment CCharSequence: CChar+;

fragment CChar: ~['\\\r\n] | EscapeSequence;

fragment EscapeSequence:
	SimpleEscapeSequence
	| OctalEscapeSequence
	| HexadecimalEscapeSequence
	| UniversalCharacterName;

fragment SimpleEscapeSequence: '\\' ['"?abfnrtv\\];

fragment OctalEscapeSequence:
	'\\' OctalDigit OctalDigit? OctalDigit?;

fragment HexadecimalEscapeSequence: '\\x' HexadecimalDigit+;

StringLiteral: EncodingPrefix? '"' SCharSequence? '"';

fragment EncodingPrefix: 'u8' | 'u' | 'U' | 'L';

fragment SCharSequence: SChar+;

fragment SChar:
	~["\\\r\n]
	| EscapeSequence
	| '\\\n' // Added line
	| '\\\r\n'; // Added line

ComplexDefine: '#' Whitespace? 'define' ~[#\r\n]* -> skip;

IncludeDirective:
	'#' Whitespace? 'include' Whitespace? (
		('"' ~[\r\n]* '"')
		| ('<' ~[\r\n]* '>')
	) Whitespace? Newline;

// ignore the following asm blocks:
/*
 asm { mfspr x, 286; }
 */
AsmBlock: 'asm' ~'{'* '{' ~'}'* '}' -> skip;

// ignore the lines generated by c preprocessor sample line : '#line 1 "/home/dm/files/dk1.h" 1'
LineAfterPreprocessing: '#line' Whitespace* ~[\r\n]* -> skip;

LineDirective:
	'#' Whitespace? DecimalConstant Whitespace? StringLiteral ~[\r\n]* -> skip;

PragmaDirective:
	'#' Whitespace? 'pragma' Whitespace ~[\r\n]* -> skip;

Whitespace: [ \t]+ -> skip;

Newline: ( '\r' '\n'? | '\n') -> skip;

BlockComment: '/*' .*? '*/';

LineComment: '//' ~[\r\n]*;
